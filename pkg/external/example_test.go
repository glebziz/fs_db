package external_test

import (
	"context"
	"log"

	_ "github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/external"
)

func ExampleOpen() {
	db, err := external.Open(context.Background(), "localhost:8888")
	if err != nil {
		log.Fatalln("Open:", err)
	}

	_ = db
}
