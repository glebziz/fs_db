package store

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true
//go:generate mockgen -source ../../../model/io.go -package mock_store -destination mocks/mocks_io.go -typed true

type storeUseCase interface {
	Set(ctx context.Context, key string, content model.Contents) error
	Get(ctx context.Context, key string) (model.ReadSeekCloser, error)
	GetKeys(ctx context.Context) ([]string, error)
	Delete(ctx context.Context, key string) error
}

type txUseCase interface {
	Begin(ctx context.Context, isoLevel model.TxIsoLevel) (string, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Service struct {
	store.UnimplementedStoreV1Server
	sUsecase  storeUseCase
	txUsecase txUseCase
}

func New(su storeUseCase, txu txUseCase) *Service {
	return &Service{
		sUsecase:  su,
		txUsecase: txu,
	}
}
