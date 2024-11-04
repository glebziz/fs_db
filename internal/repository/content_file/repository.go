package file

import (
	"github.com/glebziz/fs_db/internal/db/badger"
)

//go:generate mockgen -source ../../db/badger/manager.go -destination mocks/manager_mocks.go -typed true

type rep struct {
	p badger.Provider
}

func New(p badger.Provider) *rep {
	return &rep{p}
}

func (r *rep) key(id string) []byte {
	return []byte("fileContent/" + id)
}
