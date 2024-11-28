package main

import (
	"context"
	"fmt"
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

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Get:", err)
	}

	fmt.Println(string(b))
}
