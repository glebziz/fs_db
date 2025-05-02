package seek_direction

import (
	"io"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func ConvertToWhence(dir store.Seek_Direction) int {
	switch dir {
	case store.Seek_SeekStart:
		return io.SeekStart
	case store.Seek_SeekCurrent:
		return io.SeekCurrent
	case store.Seek_SeekEnd:
		return io.SeekEnd
	default:
		return io.SeekStart
	}
}

func ConvertToGrpc(dir model.SeekDirection) store.Seek_Direction {
	switch dir {
	case model.SeekStart:
		return store.Seek_SeekStart
	case model.SeekCurrent:
		return store.Seek_SeekCurrent
	case model.SeekEnd:
		return store.Seek_SeekEnd
	default:
		return store.Seek_SeekStart
	}
}
