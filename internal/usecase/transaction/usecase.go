package transaction

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type cleaner interface {
	Send(contentIds []string)
}

type fileRepository interface {
	HardDelete(ctx context.Context, txId string, filter *model.FileFilter) (contentIds []string, err error)
	UpdateTx(ctx context.Context, oldTxId string, newTxId string, filter *model.FileFilter) error
}

type txRepository interface {
	Store(ctx context.Context, tx model.Transaction) error
	Delete(ctx context.Context, id string) (*model.Transaction, error)
}

type generator interface {
	Generate() string
}

type useCase struct {
	cleaner cleaner

	fRepo  fileRepository
	txRepo txRepository

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
