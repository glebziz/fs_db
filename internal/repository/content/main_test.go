package content

import (
	"log"
	"os"
	"testing"

	_ "github.com/glebziz/fs_db/internal/utils/log"
)

var rootPath string

func TestMain(m *testing.M) {
	var err error

	rootPath, err = os.MkdirTemp("", "content_rep")
	if err != nil {
		log.Fatalln("Could not create root dir:", err)
	}

	c := m.Run()

	err = os.RemoveAll(rootPath)
	if err != nil {
		log.Fatalln("Could not remove root dir:", err)
	}

	os.Exit(c)
}
