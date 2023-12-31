package root

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type diskManager interface {
	Usage(ctx context.Context, path string) (*model.Stat, error)
}

type dirRepository interface {
	CountByParent(ctx context.Context, parent string) (uint64, error)
}

type useCase struct {
	rootDirs []string

	manager diskManager
	repo    dirRepository
}

func New(rootDirs []string, manager diskManager, repo dirRepository) *useCase {
	return &useCase{
		rootDirs: rootDirs,
		manager:  manager,
		repo:     repo,
	}
}
