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
}

type useCase struct {
	txStore  core.Transactions
	allStore core.Transaction

	fileRepo fileRepository
}

func New(fileRepo fileRepository) *useCase {
	return &useCase{
		allStore: core.Transaction{
			WithoutSearch: true,
		},

		fileRepo: fileRepo,
	}
}
