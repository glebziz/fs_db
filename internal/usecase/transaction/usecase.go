package transaction

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type cleaner interface {
	DeleteFilesAsync(ctx context.Context, files []model.File)
}

type fileRepository interface {
	UpdateTx(ctx context.Context, oldTxId string, newTxId string, filter model.FileFilter) ([]model.File, error)
	DeleteTx(ctx context.Context, txId string) []model.File
}

type txRepository interface {
	Store(ctx context.Context, tx model.Transaction) error
	Delete(ctx context.Context, id string) (model.Transaction, error)
}

type generator interface {
	Generate() string
}

type useCase struct {
	cleaner cleaner
	fRepo   fileRepository
	txRepo  txRepository

	idGen generator
}

func New(
	cleaner cleaner, fRepo fileRepository,
	txRepo txRepository, idGen generator,
) *useCase {
	return &useCase{
		cleaner: cleaner,
		fRepo:   fRepo,
		txRepo:  txRepo,
		idGen:   idGen,
	}
}
