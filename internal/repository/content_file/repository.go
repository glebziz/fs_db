package file

import (
	"github.com/glebziz/fs_db/internal/db/badger"
)

//go:generate mockgen -source ../../db/badger/manager.go -destination mocks/manager_mocks.go -typed true

type Repo struct {
	p badger.Provider
}

func New(p badger.Provider) *Repo {
	return &Repo{p}
}

func (r *Repo) key(id string) []byte {
	return []byte("fileContent/" + id)
}
