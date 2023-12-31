package dir

import (
	"context"
	"os"
	"testing"

	"github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/utils/log"
)

var manager *db.Manager

func TestMain(m *testing.M) {
	var (
		err    error
		dbPath = "./test.db"
	)

	manager, err = db.New(context.Background(), dbPath)
	if err != nil {
		log.Fatalln("Could not connect to database: ", err)
	}

	c := m.Run()

	manager.Close()
	err = os.Remove(dbPath)
	if err != nil {
		log.Fatalln("Could not remove database file: ", err)
	}

	os.Exit(c)
}
