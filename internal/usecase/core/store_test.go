package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
)

func TestUseCase_Store(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		requireU    requireUseCaseFunc
		err         error
	}{
		{
			name: "new tx",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(_ context.Context, file model.File) error {
						requireEqualFile(t, model.File{
							Key:       testKey,
							TxId:      testTxId,
							ContentId: testContentId,
						}, file)

						return nil
					})

				return td.newUseCase(), model.FileFilter{}
			},
			requireU: func(t *testing.T, u *useCase) {
				tx, ok := u.txStore.Get(testTxId)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, u.allStore.File(testKey).Latest())
			},
		},
		{
			name: "tx already exists",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(_ context.Context, file model.File) error {
						requireEqualFile(t, model.File{
							Key:       testKey,
							TxId:      testTxId,
							ContentId: testContentId,
						}, file)

						return nil
					})

				u := td.newUseCase()
				tx := &core.Transaction{}

				tx.PushBack((&core.Node[model.File]{}).SetV(model.File{
					Key:       testKey2,
					ContentId: testContentId2,
				}))
				tx.PushBack((&core.Node[model.File]{}).SetV(model.File{
					Key:       testKey,
					ContentId: testContentId2,
				}))
				u.txStore.Put(testTxId, tx)

				return u, model.FileFilter{}
			},
			requireU: func(t *testing.T, u *useCase) {
				tx, ok := u.txStore.Get(testTxId)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, u.allStore.File(testKey).Latest())
			},
		},
		{
			name: "file repo error",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)

				return td.newUseCase(), model.FileFilter{}
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, _ := tc.initUseCase(td)
			err := u.Store(context.Background(), model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
			})

			require.ErrorIs(t, err, tc.err)

			if tc.requireU != nil {
				tc.requireU(t, u)
			}
		})
	}
}
