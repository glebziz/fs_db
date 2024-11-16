package fs_db

import "errors"

var (
	SizeErr           = errors.New("no free space")
	NotFoundErr       = errors.New("not found")
	EmptyKeyErr       = errors.New("empty name")
	HeaderNotFoundErr = errors.New("header not found")

	// Tx errors.
	TxNotFoundErr      = errors.New("transaction not found")
	TxAlreadyExistsErr = errors.New("transaction already exists")
	TxSerializationErr = errors.New("serialization error")

	// Config errors.
	EmptyDbPathErr = errors.New("empty db path")
	EmptyRootDirs  = errors.New("empty root dirs")
)
