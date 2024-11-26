package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/external"
)

func main() {
	db, err := external.Open(context.Background(), "localhost:8888")
	if err != nil {
		log.Panicln("Open db:", err)
	}
	defer db.Close()

	err = db.Set(context.Background(), "someKey", []byte("some content"))
	if err != nil {
		log.Panicln("Set:", err)
	}

	b, err := db.Get(context.Background(), "someKey")
	if err != nil {
		log.Panicln("Get:", err)
	}

	fmt.Println(string(b))
}
