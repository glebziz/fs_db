package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/glebziz/fs_db"
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

	err = db.Set(context.Background(), "someKey", []byte("some content"))
	if err != nil {
		log.Panicln("Set:", err)
	}

	tx, err := db.Begin(context.Background(), fs_db.IsoLevelReadCommitted)
	if err != nil {
		log.Panicln("Begin:", err)
	}

	err = tx.Set(context.Background(), "someKey", []byte("some content2"))
	if err != nil {
		log.Panicln("Tx set:", err)
	}

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Get:", err)
	}

	fmt.Println("Db get:", string(b))

	b, err = tx.Get(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Get:", err)
	}

	fmt.Println("Tx get:", string(b))

	err = tx.Commit(context.Background())
	if err != nil {
		log.Panicln("Tx commit:", err)
	}

	b, err = db.Get(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Db get after tx commit:", err)
	}

	fmt.Println("Db get after tx commit:", string(b))
}
