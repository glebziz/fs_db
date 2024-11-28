package fs_db

import "errors"

var (
	ErrUnknown = errors.New("unknown error")

	ErrNoFreeSpace    = errors.New("no free space")
	ErrNotFound       = errors.New("not found")
	ErrEmptyKey       = errors.New("empty key")
	ErrHeaderNotFound = errors.New("header not found")

	// Tx errors.
	ErrTxNotFound      = errors.New("transaction not found")
	ErrTxAlreadyExists = errors.New("transaction already exists")
	ErrTxSerialization = errors.New("serialization error")

	// Config errors.
	ErrEmptyDbPath   = errors.New("empty db path")
	ErrEmptyRootDirs = errors.New("empty root dirs")
)

// Old errors to maintain backward compatibility.
//
//nolint:errname
var (
	SizeErr           = ErrNoFreeSpace
	NotFoundErr       = ErrNotFound
	EmptyKeyErr       = ErrEmptyKey
	HeaderNotFoundErr = ErrHeaderNotFound

	// Tx errors.
	TxNotFoundErr      = ErrTxNotFound
	TxAlreadyExistsErr = ErrTxAlreadyExists
	TxSerializationErr = ErrTxSerialization

	// Config errors.
	EmptyDbPathErr = ErrEmptyDbPath
	EmptyRootDirs  = ErrEmptyRootDirs
)
