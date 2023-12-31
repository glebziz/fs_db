package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/usecase"
)

type rep struct {
	p db.Provider
}

func New(p db.Provider) *rep {
	return &rep{p}
}

func (r *rep) RunTransaction(ctx context.Context, fn usecase.TransactionFn) error {
	return r.p.RunTransaction(ctx, fn)
}
