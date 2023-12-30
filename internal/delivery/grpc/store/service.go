package store

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true

type useCase interface {
	Set(ctx context.Context, key string, content *model.Content) error
	Get(ctx context.Context, key string) (*model.Content, error)
	Delete(ctx context.Context, key string) error
}

type implementation struct {
	store.UnimplementedStoreV1Server
	usecase useCase
}

func New(u useCase) *implementation {
	return &implementation{
		store.UnimplementedStoreV1Server{},
		u,
	}
}
