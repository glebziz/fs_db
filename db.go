package fs_db

import (
	"context"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	// IsoLevelReadUncommitted transaction iso level read uncommitted.
	IsoLevelReadUncommitted = model.TxIsoLevel(iota)

	// IsoLevelReadCommitted transaction iso level read committed.
	IsoLevelReadCommitted

	// IsoLevelRepeatableRead transaction iso level repeatable read.
	IsoLevelRepeatableRead

	// IsoLevelSerializable transaction iso level serializable.
	//
	// Since the set operation contains insert and update operations,
	// the serializable level is equal to the repeatable read.
	IsoLevelSerializable
)

const (
	// IsoLevelDefault the default transaction iso level is ReadCommitted.
	IsoLevelDefault = IsoLevelReadCommitted
)

// Store provides KV operations with files.
type Store interface {
	// Set sets the contents of b using the key.
	Set(ctx context.Context, key string, b []byte) error

	// SetReader sets the reader content using the key.
	SetReader(ctx context.Context, key string, reader io.Reader) error

	// Get returns content by key.
	Get(ctx context.Context, key string) ([]byte, error)

	// GetReader returns content as io.ReadCloser by key.
	GetReader(ctx context.Context, key string) (io.ReadCloser, error)

	// GetKeys returns all keys from the database.
	GetKeys(ctx context.Context) ([]string, error)

	// Delete delete content by key.
	Delete(ctx context.Context, key string) error

	// Create returns the File for to write to.
	Create(ctx context.Context, key string) (File, error)
}

// DB provides fs db interface.
type DB interface {
	Store

	// Begin starts a transaction with isoLevel.
	Begin(ctx context.Context, isoLevel ...model.TxIsoLevel) (Tx, error)

	// Close closed the connection to fs_db.
	Close() error
}
