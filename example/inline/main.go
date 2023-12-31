package main

import (
	"context"
	"fmt"

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

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Fatalln("Get:", err)
	}

	fmt.Println(string(b))
}
