package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type rootUseCase interface {
	Get(ctx context.Context) (model.RootMap, error)
}

type dirRepository interface {
	Create(ctx context.Context, d model.Dir) error
	Get(ctx context.Context) ([]model.Dir, error)
}

type generator interface {
	Generate() string
}

type useCase struct {
	maxCount uint64

	root  rootUseCase
	dRepo dirRepository
	idGen generator
}

func New(maxCount uint64, root rootUseCase, dRepo dirRepository, idGen generator) *useCase {
	return &useCase{
		maxCount: maxCount,
		root:     root,
		dRepo:    dRepo,
		idGen:    idGen,
	}
}
