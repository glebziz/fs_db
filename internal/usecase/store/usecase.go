package store

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type dirUsecase interface {
	Select(ctx context.Context, size uint64) (*model.Dir, error)
}

type contentRepository interface {
	Store(ctx context.Context, path string, content *model.Content) error
	Get(ctx context.Context, path string) (*model.Content, error)
	Delete(ctx context.Context, path string) error
}

type fileRepository interface {
	Store(ctx context.Context, file model.File) error
	Get(ctx context.Context, key string) (*model.File, error)
	Delete(ctx context.Context, key string) error
}

type generator interface {
	Generate() string
}

type useCase struct {
	dir dirUsecase

	cRepo contentRepository
	fRepo fileRepository

	idGen generator
}

func New(dir dirUsecase, cRepo contentRepository, fRepo fileRepository, idGen generator) *useCase {
	return &useCase{
		dir: dir,

		cRepo: cRepo,
		fRepo: fRepo,

		idGen: idGen,
	}
}
