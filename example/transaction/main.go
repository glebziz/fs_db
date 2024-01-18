package main

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline"
)

func main() {
	db, err := inline.Open(context.Background(), &config.Storage{
		DbPath:      "test.db",
		MaxDirCount: 1,
		RootDirs:    []string{"./testStorage"},
	})
	if err != nil {
		log.Fatalln("Open db inline:", err)
	}

	err = db.Set(context.Background(), "someKey", []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	tx, err := db.Begin(context.Background(), fs_db.IsoLevelReadCommitted)
	if err != nil {
		log.Fatalln("Begin:", err)
	}

	err = tx.Set(context.Background(), "someKey", []byte("some content2"))
	if err != nil {
		log.Fatalln("Tx set:", err)
	}

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Fatalln("Get:", err)
	}

	fmt.Println("Db get:", string(b))

	b, err = tx.Get(context.Background(), "someKey")
	if err != nil {
		log.Fatalln("Get:", err)
	}

	fmt.Println("Tx get:", string(b))

	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalln("Tx commit:", err)
	}

	b, err = db.Get(context.Background(), "someKey")
	if err != nil {
		log.Fatalln("Db get after tx commit:", err)
	}

	fmt.Println("Db get after tx commit:", string(b))
}
