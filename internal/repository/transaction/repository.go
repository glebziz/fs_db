package transaction

import (
	"github.com/puzpuzpuz/xsync/v3"

	"github.com/glebziz/fs_db/internal/model"
)

type rep struct {
	storage *xsync.MapOf[string, *model.Transaction]
}

func New() *rep {
	return &rep{
		storage: xsync.NewMapOf[string, *model.Transaction](),
	}
}
