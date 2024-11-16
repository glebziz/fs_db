package transaction

import (
	"github.com/glebziz/containers/omap"

	"github.com/glebziz/fs_db/internal/model"
)

type rep struct {
	storage *omap.OMap[string, model.Transaction]
}

func New() *rep {
	return &rep{
		storage: omap.New[string, model.Transaction](),
	}
}
