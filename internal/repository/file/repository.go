package file

import (
	"context"

	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model/transactor"
)

//go:generate mockgen -source ../../db/badger/manager.go -destination mocks/manager_mocks.go -typed true

type Repo struct {
	p badger.Provider
}

func New(p badger.Provider) *Repo {
	return &Repo{
		p: p,
	}
}

func (r *Repo) RunTransaction(ctx context.Context, fn transactor.TransactionFn) error {
	return r.p.RunTransaction(ctx, fn)
}

func (r *Repo) key(contentId string) []byte {
	return []byte("file/" + contentId)
}
