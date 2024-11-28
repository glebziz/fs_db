package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_Get(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		keys    []string
		err     error
	}{
		{
			name: "read uncommitted",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), testTxId).
					Times(1).
					Return(model.Transaction{
						Id:       testTxId,
						IsoLevel: fs_db.IsoLevelReadUncommitted,
						Seq:      testTxSeq,
					}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), testTxId, model.FileFilter{}).
					Times(1).
					Return([]model.File{{
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}, {
						Key:       testKey2,
						TxId:      testTxId,
						ContentId: testContentId2,
					}}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Return(model.ContentFile{}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId2).
					Return(model.ContentFile{}, fs_db.ErrNotFound)

				return nil
			},
			keys: []string{testKey},
		},
		{
			name: "read committed",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), testTxId).
					Times(1).
					Return(model.Transaction{
						Id:       testTxId,
						IsoLevel: fs_db.IsoLevelReadCommitted,
						Seq:      testTxSeq,
					}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), testTxId, model.FileFilter{
						TxId: ptr.Ptr(model.MainTxId),
					}).
					Times(1).
					Return([]model.File{{
						Key:       testKey2,
						TxId:      testTxId,
						ContentId: testContentId2,
					}, {
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Return(model.ContentFile{}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId2).
					Return(model.ContentFile{}, nil)

				return nil
			},
			keys: []string{testKey, testKey2},
		},
		{
			name: "repeatable read",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), testTxId).
					Times(1).
					Return(model.Transaction{
						Id:       testTxId,
						IsoLevel: fs_db.IsoLevelRepeatableRead,
						Seq:      testTxSeq,
					}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), testTxId, model.FileFilter{
						TxId:      ptr.Ptr(model.MainTxId),
						BeforeSeq: ptr.Ptr(testTxSeq),
					}).
					Times(1).
					Return([]model.File{{
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Return(model.ContentFile{}, nil)

				return nil
			},
			keys: []string{testKey},
		},
		{
			name: "serializable",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), testTxId).
					Times(1).
					Return(model.Transaction{
						Id:       testTxId,
						IsoLevel: fs_db.IsoLevelRepeatableRead,
						Seq:      testTxSeq,
					}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), testTxId, model.FileFilter{
						TxId:      ptr.Ptr(model.MainTxId),
						BeforeSeq: ptr.Ptr(testTxSeq),
					}).
					Times(1).
					Return([]model.File{{
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Return(model.ContentFile{}, nil)

				return nil
			},
			keys: []string{testKey},
		},
		{
			name: "tx get error",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
		{
			name: "files get error",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
		{
			name: "content file get error",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					GetFiles(gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]model.File{{}}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.ContentFile{}, assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			uc := td.newUseCase()

			keys, err := uc.GetKeys(testCtx)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.keys, keys)
		})
	}
}
