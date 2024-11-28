package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Load(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		requireU    requireUseCaseFunc
		deleteFiles []model.File
		err         error
	}{
		{
			name: "success",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					GetAll(gomock.Any()).
					Times(1).
					Return([]model.File{{
						Key:       testKey,
						TxId:      model.MainTxId,
						ContentId: testContentId,
						Seq:       1,
					}, {
						Key:       testKey,
						TxId:      model.MainTxId,
						ContentId: testContentId2,
						Seq:       2,
					}, {
						Key:       testKey2,
						TxId:      model.MainTxId,
						ContentId: testContentId3,
						Seq:       3,
					}, {
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId4,
						Seq:       4,
					}, {
						Key:       testKey2,
						TxId:      model.MainTxId,
						ContentId: testContentId5,
						Seq:       2,
					}, {
						Key:       testKey2,
						TxId:      testTxId,
						ContentId: testContentId6,
						Seq:       5,
					}}, nil)

				return td.newUseCase(), model.FileFilter{}
			},
			requireU: func(t *testing.T, u *UseCase) {
				require.Equal(t, model.File{
					Key:       testKey,
					TxId:      model.MainTxId,
					ContentId: testContentId2,
					Seq:       2,
				}, u.allStore.File(testKey).Latest())
				require.Equal(t, model.File{
					Key:       testKey2,
					TxId:      model.MainTxId,
					ContentId: testContentId3,
					Seq:       3,
				}, u.allStore.File(testKey2).Latest())
			},
			deleteFiles: []model.File{{
				Key:       testKey,
				TxId:      model.MainTxId,
				ContentId: testContentId,
				Seq:       1,
			}, {
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId4,
				Seq:       4,
			}, {
				Key:       testKey2,
				TxId:      model.MainTxId,
				ContentId: testContentId5,
				Seq:       2,
			}, {
				Key:       testKey2,
				TxId:      testTxId,
				ContentId: testContentId6,
				Seq:       5,
			}},
		},
		{
			name: "db get all error",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					GetAll(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)

				return td.newUseCase(), model.FileFilter{}
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, _ := tc.initUseCase(td)

			deleteFiles, err := u.Load(context.Background())

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.deleteFiles, deleteFiles)

			if tc.requireU != nil {
				tc.requireU(t, u)
			}
		})
	}
}
