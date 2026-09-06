package fs_db

import (
	"io"
)

type File interface {
	io.WriteSeeker
	io.WriterAt
	io.Closer
}

type ReadFile interface {
	io.ReadSeekCloser
	io.ReaderAt
}
