package fs_db

import (
	"context"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	IsoLevelReadUncommitted = model.TxIsoLevel(iota)
	IsoLevelReadCommitted
	IsoLevelRepeatableRead
	IsoLevelSerializable
)

const (
	IsoLevelDefault = IsoLevelReadCommitted
)

type store interface {
	Set(ctx context.Context, key string, b []byte) error
	SetReader(ctx context.Context, key string, reader io.Reader, size uint64) error
	Get(ctx context.Context, key string) ([]byte, error)
	GetReader(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

type DB interface {
	store
	Begin(ctx context.Context, isoLevel ...model.TxIsoLevel) (Tx, error)
}
