package writer

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -package mocks -destination mocks/mocks.go -typed true
//go:generate mockgen -source ../../model/io.go -package mocks -destination mocks/mocks_io.go -typed true

type contentRepository interface {
	Create(ctx context.Context, path string) (model.ReadWriteSeekCloser, error)
}

type UseCase struct {
	cRepo contentRepository
}

func New(cRepo contentRepository) *UseCase {
	return &UseCase{
		cRepo: cRepo,
	}
}
