package fs_db_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
	_ "github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline"
)

func ExampleStore_Set() {
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

	err = db.Set(context.Background(), "someKey", []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}
}

func ExampleStore_SetReader() {
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

	content := []byte("some content")
	err = db.SetReader(context.Background(), "someKey", bytes.NewReader(content))
	if err != nil {
		log.Fatalln("SetReader:", err)
	}
}

func ExampleStore_Get() {
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

	key := "someKey"
	err = db.Set(context.Background(), key, []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	b, err := db.Get(context.Background(), key)
	if err != nil {
		log.Fatalln("Get:", err)
	}

	fmt.Println(string(b))

	// Output:
	// some content
}

func ExampleStore_GetReader() {
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

	key := "someKey"
	err = db.Set(context.Background(), key, []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	r, err := db.GetReader(context.Background(), key)
	if err != nil {
		log.Fatalln("Get:", err)
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatalln("ReadAll:", err)
	}

	fmt.Println(string(b))

	// Output:
	// some content
}

func ExampleStore_GetKeys() {
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

	key := "someKey"
	err = db.Set(context.Background(), key, []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	keys, err := db.GetKeys(context.Background())
	if err != nil {
		log.Fatalln("Get keys:", err)
	}

	fmt.Println(keys)

	// Output:
	// [someKey]
}

func ExampleStore_Delete() {
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

	key := "someKey"
	err = db.Set(context.Background(), key, []byte("some content"))
	if err != nil {
		log.Fatalln("Set:", err)
	}

	err = db.Delete(context.Background(), key)
	if err != nil {
		log.Fatalln("Get:", err)
	}

	_, err = db.Get(context.Background(), key)
	if errors.Is(err, fs_db.NotFoundErr) {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatalln("Get:", err)
	}

	// Output:
	// key not found
}

func ExampleDB_Begin() {
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

	tx, err := db.Begin(context.Background(), fs_db.IsoLevelReadCommitted)
	if err != nil {
		log.Fatalln("Begin:", err)
	}
	defer tx.Rollback(context.Background())

	err = tx.Commit(context.Background())
}
