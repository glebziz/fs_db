package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type dirRepository interface {
	GetRoots(ctx context.Context) ([]model.Root, error)
	Get(ctx context.Context) (model.Dirs, error)
	Create(ctx context.Context, d model.Dir) error
	Remove(ctx context.Context, d model.Dir) error
}

type generator interface {
	Generate() string
}

type useCase struct {
	maxCount uint64

	dRepo   dirRepository
	nameGen generator
}

func New(maxCount uint64, dRepo dirRepository, nameGen generator) *useCase {
	return &useCase{
		maxCount: maxCount,
		dRepo:    dRepo,
		nameGen:  nameGen,
	}
}
