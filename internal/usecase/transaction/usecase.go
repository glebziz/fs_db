package transaction

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type fileRepository interface {
	UpdateTx(ctx context.Context, oldTxId string, newTxId string, filter model.FileFilter) error
	DeleteTx(ctx context.Context, txId string)
}

type txRepository interface {
	Store(ctx context.Context, tx model.Transaction) error
	Delete(ctx context.Context, id string) (*model.Transaction, error)
}

type generator interface {
	Generate() string
}

type useCase struct {
	fRepo  fileRepository
	txRepo txRepository

	idGen generator
}

func New(
	fRepo fileRepository, txRepo txRepository,
	idGen generator,
) *useCase {
	return &useCase{
		fRepo:  fRepo,
		txRepo: txRepo,
		idGen:  idGen,
	}
}
