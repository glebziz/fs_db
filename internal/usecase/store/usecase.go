package store

import (
	"context"
	"math/rand/v2"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -package mock_store -destination mocks/mocks.go -typed true
//go:generate mockgen -source ../../model/io.go -package mock_store -destination mocks/mocks_io.go -typed true

type dirUsecase interface {
	Get(ctx context.Context) (model.Dirs, error)
}

type contentRepository interface {
	Get(ctx context.Context, path string) (model.ReadSeekCloser, error)
}

type contentFileRepository interface {
	Store(ctx context.Context, file model.ContentFile) error
	Get(ctx context.Context, id string) (model.ContentFile, error)
}

type fileRepository interface {
	Store(ctx context.Context, file model.File) error
	Get(ctx context.Context, txId, key string, filter model.FileFilter) (model.File, error)
	GetFiles(ctx context.Context, txId string, filter model.FileFilter) ([]model.File, error)
}

type txRepository interface {
	Get(ctx context.Context, id string) (model.Transaction, error)
}

type contentWriter interface {
	Write(ctx context.Context, path string, contents model.Contents) (err error)
}

type generator interface {
	Generate() string
}

type UseCase struct {
	dir dirUsecase

	cRepo   contentRepository
	cfRepo  contentFileRepository
	fRepo   fileRepository
	txRepo  txRepository
	cWriter contentWriter

	idGen   generator
	randGen *rand.Rand
}

func New(
	dir dirUsecase, cRepo contentRepository,
	cfRepo contentFileRepository, fRepo fileRepository,
	txRepo txRepository, cWriter contentWriter,
	idGen generator, randGen *rand.Rand,
) *UseCase {
	return &UseCase{
		dir: dir,

		cRepo:   cRepo,
		cfRepo:  cfRepo,
		fRepo:   fRepo,
		txRepo:  txRepo,
		cWriter: cWriter,

		idGen:   idGen,
		randGen: randGen,
	}
}
