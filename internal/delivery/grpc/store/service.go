package store

import (
	"context"
	"io"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true

type storeUseCase interface {
	Set(ctx context.Context, key string, content io.Reader) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

type txUseCase interface {
	Begin(ctx context.Context, isoLevel model.TxIsoLevel) (string, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type implementation struct {
	store.UnimplementedStoreV1Server
	sUsecase  storeUseCase
	txUsecase txUseCase
}

func New(su storeUseCase, txu txUseCase) *implementation {
	return &implementation{
		UnimplementedStoreV1Server: store.UnimplementedStoreV1Server{},

		sUsecase:  su,
		txUsecase: txu,
	}
}
