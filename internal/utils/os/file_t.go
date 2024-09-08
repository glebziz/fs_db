//go:build test

package os

var (
	spaceLimit int64
)

func init() {
	SetSpaceLimit(1 << 40)
}

func SetSpaceLimit(limit int64) {
	spaceLimit = limit
}

func (f File) Write(p []byte) (n int, err error) {
	spaceLimit -= int64(len(p))

	if spaceLimit <= 0 {
		return 0, ErrNotEnoughSpace
	}

	return f.File.Write(p)
}
