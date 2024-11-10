package file

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_GetAll(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		files   []model.File
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, data)
				require.NoError(t, err)

				data2 := make([]byte, fileLenWithoutKey+len(testKey2))
				err = marshalFile(model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       testSeq2,
				}, data2)
				require.NoError(t, err)

				td.qm.EXPECT().
					GetAll([]byte("file/")).
					Times(1).
					Return([]badger.Item{{
						Key:   []byte(testContentId),
						Value: data,
					}, {
						Key:   []byte(testContentId2),
						Value: data2,
					}}, nil)
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId2,
				Seq:       testSeq2,
			}},
		},
		{
			name: "files not found",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					GetAll(gomock.Any()).
					Times(1).
					Return(nil, nil)
			},
			files: []model.File{},
		},
		{
			name: "db get all error",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					GetAll(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "unmarshal file error",
			prepare: func(td *testDeps) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq2,
				}, data)
				require.NoError(t, err)

				td.qm.EXPECT().
					GetAll(gomock.Any()).
					Times(1).
					Return([]badger.Item{{
						Key:   []byte(testContentId),
						Value: data,
					}, {
						Key:   []byte(testContentId2),
						Value: []byte{1, 2, 3},
					}}, nil)
			},
			err: model.ErrInvalidFileFormat,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRep()
			files, err := r.GetAll(context.Background())

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.files, files)
		})
	}
}

func TestRep_GetAll_Int(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareIntFunc
		wrapTx  bool
		files   []model.File
	}{
		{
			name: "success",
			prepare: func(t *testing.T, p badger.Provider) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, data)
				require.NoError(t, err)

				data2 := make([]byte, fileLenWithoutKey+len(testKey2))
				err = marshalFile(model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       testSeq2,
				}, data2)
				require.NoError(t, err)

				err = p.DB(context.Background()).Set(append([]byte("file/"), testContentId...), data)
				require.NoError(t, err)

				err = p.DB(context.Background()).Set(append([]byte("file/"), testContentId2...), data2)
				require.NoError(t, err)
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId2,
				Seq:       testSeq2,
			}},
		},
		{
			name:    "files not found",
			prepare: func(t *testing.T, p badger.Provider) {},
			files:   []model.File{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := newTestRep(t)
			tc.prepare(t, r.p)

			files, err := r.GetAll(context.Background())

			require.NoError(t, err)
			require.Equal(t, tc.files, files)
		})
	}
}
