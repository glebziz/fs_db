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

func TestRep_Store(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		f       model.File
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

				td.qm.EXPECT().
					Set(append([]byte("file/"), testContentId...), data).
					Times(1).
					Return(nil)
			},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
		{
			name:    "marshal file error",
			prepare: func(td *testDeps) {},
			f: model.File{
				Key:       testKey,
				TxId:      "testTxId",
				ContentId: testContentId,
				Seq:       testSeq,
			},
			err: model.ErrInvalidFileFormat,
		},
		{
			name: "db set error",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRep()
			err := r.Set(context.Background(), tc.f)

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRep_Store_Int(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareIntFunc
		f       model.File
		wrapTx  bool
		require prepareIntFunc
	}{
		{
			name:    "success",
			prepare: func(t *testing.T, p badger.Provider) {},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			require: func(t *testing.T, p badger.Provider) {
				data, err := p.DB(context.Background()).Get(append([]byte("file/"), testContentId...))
				require.NoError(t, err)

				var f model.File
				err = unmarshalFile(data, &f)
				require.NoError(t, err)

				require.Equal(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, f)
			},
		},
		{
			name: "success with file already exists",
			prepare: func(t *testing.T, p badger.Provider) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, data)
				require.NoError(t, err)

				err = p.DB(context.Background()).Set(append([]byte("file/"), testContentId...), data)
				require.NoError(t, err)
			},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
				Seq:       testSeq2,
			},
			require: func(t *testing.T, p badger.Provider) {
				data, err := p.DB(context.Background()).Get(append([]byte("file/"), testContentId...))
				require.NoError(t, err)

				var f model.File
				err = unmarshalFile(data, &f)
				require.NoError(t, err)

				require.Equal(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       testSeq2,
				}, f)
			},
		},
		{
			name:    "success with tx",
			prepare: func(t *testing.T, p badger.Provider) {},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			},
			wrapTx: true,
			require: func(t *testing.T, p badger.Provider) {
				data, err := p.DB(context.Background()).Get(append([]byte("file/"), testContentId...))
				require.NoError(t, err)

				var f model.File
				err = unmarshalFile(data, &f)
				require.NoError(t, err)

				require.Equal(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, f)
			},
		},
		{
			name: "success with file already exists with tx",
			prepare: func(t *testing.T, p badger.Provider) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, data)
				require.NoError(t, err)

				err = p.DB(context.Background()).Set(append([]byte("file/"), testContentId...), data)
				require.NoError(t, err)
			},
			f: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
				Seq:       testSeq2,
			},
			wrapTx: true,
			require: func(t *testing.T, p badger.Provider) {
				data, err := p.DB(context.Background()).Get(append([]byte("file/"), testContentId...))
				require.NoError(t, err)

				var f model.File
				err = unmarshalFile(data, &f)
				require.NoError(t, err)

				require.Equal(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       testSeq2,
				}, f)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := newTestRep(t)
			tc.prepare(t, r.p)

			var err error
			if tc.wrapTx {
				err = r.RunTransaction(context.Background(), func(ctx context.Context) error {
					return r.Set(ctx, tc.f)
				})
			} else {
				err = r.Set(context.Background(), tc.f)
			}
			require.NoError(t, err)

			tc.require(t, r.p)
		})
	}
}
