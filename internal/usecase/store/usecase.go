package store

import (
	"context"
	"io"
	"math/rand/v2"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type dirUsecase interface {
	Get(ctx context.Context) (model.Dirs, error)
}

type contentRepository interface {
	Store(ctx context.Context, path string, content io.Reader) error
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}

type contentFileRepository interface {
	Store(ctx context.Context, file model.ContentFile) error
	Get(ctx context.Context, id string) (model.ContentFile, error)
}

type fileRepository interface {
	Store(ctx context.Context, file model.File) error
	Get(ctx context.Context, txId, key string, filter model.FileFilter) (model.File, error)
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

	idGen   generator
	randGen *rand.Rand
}

func New(
	dir dirUsecase, cRepo contentRepository,
	cfRepo contentFileRepository, fRepo fileRepository,
	txRepo txRepository, idGen generator,
	randGen *rand.Rand,
) *useCase {
	return &useCase{
		dir: dir,

		cRepo:  cRepo,
		cfRepo: cfRepo,
		fRepo:  fRepo,
		txRepo: txRepo,

		idGen:   idGen,
		randGen: randGen,
	}
}
