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

func TestRep_Delete(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Delete(append([]byte("file/"), testContentId...)).
					Times(1).
					Return(nil)
			},
		},
		{
			name: "db delete error",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Delete(gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRep()
			err := r.Delete(context.Background(), model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
				Seq:       testSeq,
			})

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRep_Delete_Int(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareIntFunc
		wrapTx  bool
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

				err = p.DB(context.Background()).Set([]byte(testContentId), data)
				require.NoError(t, err)
			},
		},
		{
			name: "success with tx",
			prepare: func(t *testing.T, p badger.Provider) {
				data := make([]byte, fileLenWithoutKey+len(testKey))
				err := marshalFile(model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}, data)
				require.NoError(t, err)

				err = p.DB(context.Background()).Set([]byte(testContentId), data)
				require.NoError(t, err)
			},
			wrapTx: true,
		},
		{
			name:    "file not found",
			prepare: func(_ *testing.T, _ badger.Provider) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := newTestRep(t)
			tc.prepare(t, r.p)

			var (
				err error

				f = model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       testSeq,
				}
			)
			if tc.wrapTx {
				err = r.RunTransaction(context.Background(), func(ctx context.Context) error {
					return r.Delete(ctx, f)
				})
			} else {
				err = r.Delete(context.Background(), f)
			}

			require.NoError(t, err)
		})
	}
}
