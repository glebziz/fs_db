package fs_db

import (
	"io"
)

type File interface {
	io.WriteCloser
}
