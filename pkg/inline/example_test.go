package inline_test

import (
	"context"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline"
)

func ExampleOpen() {
	db, err := inline.Open(context.Background(), &config.Storage{
		DbPath:      "test.db",
		MaxDirCount: 1,
		RootDirs:    []string{"./testStorage"},
	})
	if err != nil {
		log.Fatalln("Open:", err)
	}

	_ = db
}
