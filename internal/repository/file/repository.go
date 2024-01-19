package file

import (
	"github.com/glebziz/fs_db/internal/db"
)

type rep struct {
	p db.Provider
}

func New(p db.Provider) *rep {
	return &rep{p}
}
