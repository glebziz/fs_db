package file

import (
	"context"

	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model/transactor"
)

//go:generate mockgen -source ../../db/badger/manager.go -destination mocks/manager_mocks.go -typed true

type rep struct {
	p badger.Provider
}

func New(p badger.Provider) *rep {
	return &rep{
		p: p,
	}
}

func (r *rep) RunTransaction(ctx context.Context, fn transactor.TransactionFn) error {
	return r.p.RunTransaction(ctx, fn)
}

func (r *rep) key(contentId string) []byte {
	return []byte("file/" + contentId)
}
