package transaction

import (
	"github.com/glebziz/containers/omap"

	"github.com/glebziz/fs_db/internal/model"
)

type Repo struct {
	storage *omap.OMap[string, model.Transaction]
}

func New() *Repo {
	return &Repo{
		storage: omap.New[string, model.Transaction](),
	}
}
