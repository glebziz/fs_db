package file

import (
	"encoding/binary"
	"unsafe"

	"github.com/google/uuid"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

const (
	uuidLen = 16
	timeLen = 8

	fileLenWithoutKey = 2*uuidLen + timeLen
)

func fileLen(f model.File) (l int) {
	return fileLenWithoutKey + len(f.Key)
}

func marshalFile(f model.File, data []byte) error {
	if len(data) != fileLen(f) {
		return model.ErrInvalidFileFormat
	}

	txId, err := uuid.Parse(f.TxId)
	if err != nil {
		return model.ErrInvalidFileFormat
	}

	contentId, err := uuid.Parse(f.ContentId)
	if err != nil {
		return model.ErrInvalidFileFormat
	}

	binary.LittleEndian.PutUint64(data[:timeLen], uint64(f.Seq))
	copy(data[timeLen:timeLen+uuidLen], txId[:])
	copy(data[timeLen+uuidLen:fileLenWithoutKey], contentId[:])
	copy(data[fileLenWithoutKey:], f.Key)

	return nil
}

func unmarshalFile(data []byte, f *model.File) error {
	if f == nil || len(data) < fileLenWithoutKey {
		return model.ErrInvalidFileFormat
	}

	f.Seq = sequence.Seq(binary.LittleEndian.Uint64(data[:timeLen]))
	f.TxId = uuid.UUID(data[timeLen : timeLen+uuidLen]).String()
	f.ContentId = uuid.UUID(data[timeLen+uuidLen : fileLenWithoutKey]).String()

	if len(data)-fileLenWithoutKey > 0 {
		f.Key = unsafe.String(&data[fileLenWithoutKey], len(data)-fileLenWithoutKey)
	}

	return nil
}
