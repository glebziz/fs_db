package inline_test

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/glebziz/fs_db/config"
	_ "github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline"
)

func ExampleOpen() {
	db, err := inline.Open(context.Background(), config.Config{
		Storage: config.Storage{
			DbPath:      "test_db",
			MaxDirCount: 1,
			RootDirs:    []string{"./testStorage"},
			GCPeriod:    1 * time.Minute,
		},
		WPool: config.WPool{
			NumWorkers:   runtime.GOMAXPROCS(0),
			SendDuration: 1 * time.Millisecond,
		},
	})
	if err != nil {
		log.Fatalln("Open:", err)
	}

	_ = db
}
