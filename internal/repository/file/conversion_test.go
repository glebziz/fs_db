package file

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func Test_fileLen(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    model.File
		len  int
	}{
		{
			name: "non empty file",
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			len: fileLenWithoutKey + len(testKey),
		},
		{
			name: "empty file",
			f:    model.File{},
			len:  fileLenWithoutKey,
		},
		{
			name: "file without key",
			f: model.File{
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			len: fileLenWithoutKey,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := fileLen(tc.f)
			require.Equal(t, tc.len, l)
		})
	}
}

func Test_marshalFile(t *testing.T) {
	for _, tc := range []struct {
		name   string
		f      model.File
		data   []byte
		result []byte
		err    error
	}{
		{
			name: "success",
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			data: make([]byte, fileLenWithoutKey+len(testKey)),
			result: append([]byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
				0x0, 0x0, 0x1, 0x1, 0x2, 0x2, 0x3, 0x3,
				0x4, 0x4, 0x5, 0x5, 0x6, 0x6, 0x7, 0x7,
			}, testKey...),
		},
		{
			name: "file with empty key",
			f: model.File{
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			data: make([]byte, fileLenWithoutKey),
			result: []byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
				0x0, 0x0, 0x1, 0x1, 0x2, 0x2, 0x3, 0x3,
				0x4, 0x4, 0x5, 0x5, 0x6, 0x6, 0x7, 0x7,
			},
		},
		{
			name: "invalid buffer len",
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			data:   make([]byte, fileLenWithoutKey),
			result: make([]byte, fileLenWithoutKey),
			err:    model.ErrInvalidFileFormat,
		},
		{
			name: "invalid txId format",
			f: model.File{
				Key:       testKey,
				TxId:      "testTxId",
				ContentId: testContentId,
				Seq:       testSeq,
			},
			data:   make([]byte, fileLenWithoutKey+len(testKey)),
			result: make([]byte, fileLenWithoutKey+len(testKey)),
			err:    model.ErrInvalidFileFormat,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := marshalFile(tc.f, tc.data)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.result, tc.data)
		})
	}
}

func Test_unmarshalFile(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []byte
		inF  *model.File
		f    *model.File
		err  error
	}{
		{
			name: "success",
			data: append([]byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
				0x0, 0x0, 0x1, 0x1, 0x2, 0x2, 0x3, 0x3,
				0x4, 0x4, 0x5, 0x5, 0x6, 0x6, 0x7, 0x7,
			}, testKey...),
			inF: &model.File{},
			f: &model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
		{
			name: "file with empty key",
			data: []byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
				0x0, 0x0, 0x1, 0x1, 0x2, 0x2, 0x3, 0x3,
				0x4, 0x4, 0x5, 0x5, 0x6, 0x6, 0x7, 0x7,
			},
			inF: &model.File{},
			f: &model.File{
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
		{
			name: "nil in file",
			data: []byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
				0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
				0x0, 0x0, 0x1, 0x1, 0x2, 0x2, 0x3, 0x3,
				0x4, 0x4, 0x5, 0x5, 0x6, 0x6, 0x7, 0x7,
			},
			err: model.ErrInvalidFileFormat,
		},
		{
			name: "invalid buffer len",
			data: make([]byte, fileLenWithoutKey-1),
			inF:  &model.File{},
			f:    &model.File{},
			err:  model.ErrInvalidFileFormat,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := unmarshalFile(tc.data, tc.inF)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.f, tc.inF)
		})
	}
}

func Test_fileConversion(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    model.File
	}{
		{
			name: "success",
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
		{
			name: "file with empty key",
			f: model.File{
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data := make([]byte, fileLen(tc.f))
			err := marshalFile(tc.f, data)
			require.NoError(t, err)

			var f model.File
			err = unmarshalFile(data, &f)
			require.NoError(t, err)

			require.Equal(t, tc.f, f)
		})
	}
}
