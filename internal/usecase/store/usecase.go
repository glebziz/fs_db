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
}

type contentFileRepository interface {
	Store(ctx context.Context, file model.ContentFile) error
	Get(ctx context.Context, id string) (*model.ContentFile, error)
}

type fileRepository interface {
	Store(ctx context.Context, txId string, file model.File) error
	Get(ctx context.Context, txId, key string, filter *model.FileFilter) (*model.File, error)
	Delete(ctx context.Context, txId, key string) error
}

type txRepository interface {
	Get(ctx context.Context, id string) (*model.Transaction, error)
}

type generator interface {
	Generate() string
}

type useCase struct {
	dir dirUsecase

	cRepo  contentRepository
	cfRepo contentFileRepository
	fRepo  fileRepository
	txRepo txRepository

	idGen generator
}

func New(
	dir dirUsecase, cRepo contentRepository,
	cfRepo contentFileRepository, fRepo fileRepository,
	txRepo txRepository, idGen generator,
) *useCase {
	return &useCase{
		dir: dir,

		cRepo:  cRepo,
		cfRepo: cfRepo,
		fRepo:  fRepo,
		txRepo: txRepo,

		idGen: idGen,
	}
}
