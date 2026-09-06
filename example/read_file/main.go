package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/glebziz/fs_db/config"
	_ "github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline"
)

func main() {
	db, err := inline.Open(context.Background(), config.Config{
		Storage: config.Storage{
			DbPath:      "test_db",
			MaxDirCount: 1,
			RootDirs:    []string{"./testStorage"},
			GCPeriod:    time.Minute,
		},
		WPool: config.WPool{
			NumWorkers:   runtime.GOMAXPROCS(0),
			SendDuration: time.Millisecond,
		},
	})
	if err != nil {
		log.Panicln("Open db inline:", err)
	}
	defer db.Close()

	f, err := db.Create(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Create:", err)
	}

	_, err = f.Write([]byte("some content"))
	if err != nil {
		log.Panicln("Write:", err)
	}

	err = f.Close()
	if err != nil {
		log.Panicln("Close:", err)
	}

	rf, err := db.Open(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Open:", err)
	}

	_, err = rf.Seek(int64(len("some ")), io.SeekStart)
	if err != nil {
		log.Panicln("Seek:", err)
	}

	b, err := io.ReadAll(rf)
	if err != nil {
		log.Panicln("Get:", err)
	}

	fmt.Println(string(b))
}
