package os

import "os"

var (
	ErrNotExist = os.ErrNotExist
)

func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func Create(name string) (File, error) {
	f, err := os.Create(name)
	return File{f}, err
}

func Open(name string) (File, error) {
	f, err := os.Open(name)
	return File{f}, err
}

func Remove(name string) error {
	return os.Remove(name)
}
