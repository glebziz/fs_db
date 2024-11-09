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

func TestUseCase_Get_Success(t *testing.T) {
	for _, tc := range []struct {
		name   string
		lvl    model.TxIsoLevel
		filter model.FileFilter
	}{
		{
			name: "read uncommitted",
			lvl:  fs_db.IsoLevelReadUncommitted,
		},
		{
			name: "read committed",
			lvl:  fs_db.IsoLevelReadCommitted,
			filter: model.FileFilter{
				TxId: ptr.Ptr(model.MainTxId),
			},
		},
		{
			name: "repeatable read",
			lvl:  fs_db.IsoLevelRepeatableRead,
			filter: model.FileFilter{
				TxId:      ptr.Ptr(model.MainTxId),
				BeforeSeq: ptr.Ptr(testTxSeq),
			},
		},
		{
			name: "serializable",
			lvl:  fs_db.IsoLevelSerializable,
			filter: model.FileFilter{
				TxId:      ptr.Ptr(model.MainTxId),
				BeforeSeq: ptr.Ptr(testTxSeq),
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				tx = model.Transaction{
					Id:       testTxId,
					IsoLevel: tc.lvl,
					Seq:      testTxSeq,
				}
				dir = model.Dir{
					Name: testDirName,
					Root: testRootPath,
				}
				cFile = model.ContentFile{
					Id:     testContentId,
					Parent: dir.Path(),
				}
				file = model.File{
					Key:       testKey,
					ContentId: testContentId,
				}
				content = testReader
			)
			td := newTestDeps(t)

			td.txRepo.EXPECT().
				Get(gomock.Any(), testTxId).
				Return(tx, nil)

			td.fRepo.EXPECT().
				Get(gomock.Any(), testTxId, testKey, tc.filter).
				Return(file, nil)

			td.cfRepo.EXPECT().
				Get(gomock.Any(), testContentId).
				Return(cFile, nil)

			td.cRepo.EXPECT().
				Get(gomock.Any(), cFile.Path()).
				Return(content, nil)

			uc := td.newUseCase()

			actContent, err := uc.Get(testCtx, testKey)

			require.NoError(t, err)
			require.Equal(t, content, actContent)
		})
	}
}

func TestUseCase_Get_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "tx get",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file get",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.File{}, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content file get",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.File{
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.ContentFile{}, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content get",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.File{
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId,
					}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(model.ContentFile{
						Id:     testContentId,
						Parent: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			actContent, err := uc.Get(testCtx, testKey)

			require.ErrorIs(t, err, wantErr)
			require.Nil(t, actContent)
		})
	}
}
