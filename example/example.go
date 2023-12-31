package main

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/fs_db"
)

func main() {
	db, err := fs_db.Open(context.Background(), "localhost:8888")
	if err != nil {
		log.Fatalln("Open db:", err)
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
