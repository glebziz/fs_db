package cleaner

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_DeleteOld(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(model.Transaction{
						Seq: testSeq,
					}, nil)

				td.core.EXPECT().
					DeleteOld(gomock.Any(), model.MainTxId, testSeq).
					Times(1).
					Return([]model.File{})
			},
		},
		{
			name: "tx not found",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(model.Transaction{}, fs_db.TxNotFoundErr)

				td.core.EXPECT().
					DeleteOld(gomock.Any(), model.MainTxId, gomock.Any()).
					Times(1).
					Return([]model.File{})
			},
		},
		{
			name: "get tx error",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(model.Transaction{}, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "delete files error",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(model.Transaction{
						Seq: testSeq,
					}, nil)

				td.core.EXPECT().
					DeleteOld(gomock.Any(), model.MainTxId, testSeq).
					Times(1).
					Return([]model.File{{
						ContentId: testContentId,
					}})

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Times(1).
					Return(model.ContentFile{}, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			u := td.newUseCase()

			err := u.DeleteOld(context.Background())
			require.ErrorIs(t, err, tc.err)
		})
	}
}

//func TestCleaner_deleteLines_Error(t *testing.T) {
//	for _, tc := range []struct {
//		name    string
//		prepare prepareFunc
//	}{
//		{
//			name: "tx repo oldest",
//			prepare: func(td *testDeps) error {
//				td.txRepo.EXPECT().
//					Oldest(gomock.Any()).
//					Return(nil, assert.AnError)
//
//				return assert.AnError
//			},
//		},
//		{
//			name: "file repo hard delete",
//			prepare: func(td *testDeps) error {
//				td.txRepo.EXPECT().
//					Oldest(gomock.Any()).
//					Return(nil, fs_db.TxNotFoundErr)
//
//				td.fRepo.EXPECT().
//					DeleteOld(gomock.Any(), gomock.Any(), gomock.Any()).
//					Return(nil, assert.AnError)
//
//				return assert.AnError
//			},
//		},
//		{
//			name: "delete content",
//			prepare: func(td *testDeps) error {
//				td.txRepo.EXPECT().
//					Oldest(gomock.Any()).
//					Return(nil, fs_db.TxNotFoundErr)
//
//				td.fRepo.EXPECT().
//					DeleteOld(gomock.Any(), gomock.Any(), gomock.Any()).
//					Return([]string{testContentId}, nil)
//
//				td.cfRepo.EXPECT().
//					GetIn(gomock.Any(), gomock.Any()).
//					Return(nil, assert.AnError)
//
//				return assert.AnError
//			},
//		},
//	} {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//
//			td := newTestDeps(t)
//
//			wantErr := tc.prepare(td)
//
//			cl := td.newUseCase()
//
//			err := cl.deleteLines(testCtx)
//			require.ErrorIs(t, err, wantErr)
//		})
//	}
//}
