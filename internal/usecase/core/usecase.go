package core

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/transactor"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type fileRepository interface {
	transactor.Transactor
	Set(ctx context.Context, file model.File) error
	GetAll(ctx context.Context) ([]model.File, error)
}

type UseCase struct {
	txStore  core.Transactions
	allStore core.Transaction
	txPool   *core.Pool[core.Transaction]
	nodePool core.Pool[core.Node[model.File]]

	fileRepo fileRepository
}

func New(fileRepo fileRepository) *UseCase {
	return &UseCase{
		allStore: core.Transaction{
			WithoutSearch: true,
		},
		txPool: core.NewPool(func(tx *core.Transaction) {
			tx.Clear()
		}),

		fileRepo: fileRepo,
	}
}
