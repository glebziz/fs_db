package model

import "errors"

var (
	SizeErr           = errors.New("no free space")
	NotFoundErr       = errors.New("not found")
	EmptyKeyErr       = errors.New("empty name")
	HeaderNotFoundErr = errors.New("header not found")

	// Config errors
	EmptyDbPathErr = errors.New("empty db path")
	EmptyRootDirs  = errors.New("empty root dirs")
)
