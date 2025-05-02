package model

import (
	"io"
)

type ReadWriteSeekCloser interface {
	io.ReadWriteSeeker
	io.ReaderAt
	io.Closer
}

type ReadSeekCloser interface {
	io.ReadSeekCloser
	io.ReaderAt
}
