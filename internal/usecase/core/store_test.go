package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	mock_core "github.com/glebziz/fs_db/internal/usecase/core/mocks"
)

func TestUseCase_Store(t *testing.T) {
	for _, tc := range []struct {
		name       string
		newUseCase func(t *testing.T) *useCase
		requireU   func(t *testing.T, u *useCase)
		err        error
	}{
		{
			name: "new tx",
			newUseCase: func(t *testing.T) *useCase {
				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(_ context.Context, file model.File) error {
						requireEqualFiles(t, model.File{
							Key:       testKey,
							TxId:      testTxId,
							ContentId: testContentId,
						}, file)

						return nil
					})

				return New(fRepo)
			},
			requireU: func(t *testing.T, u *useCase) {
				tx, ok := u.txStore.Get(testTxId)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, u.allStore.File(testKey).Latest())
			},
		},
		{
			name: "tx already exists",
			newUseCase: func(t *testing.T) *useCase {
				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(_ context.Context, file model.File) error {
						requireEqualFiles(t, model.File{
							Key:       testKey,
							TxId:      testTxId,
							ContentId: testContentId,
						}, file)

						return nil
					})

				u := New(fRepo)
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

				return u
			},
			requireU: func(t *testing.T, u *useCase) {
				tx, ok := u.txStore.Get(testTxId)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, u.allStore.File(testKey).Latest())
			},
		},
		{
			name: "file repo error",
			newUseCase: func(t *testing.T) *useCase {
				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)

				return New(fRepo)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file := model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
			}

			u := tc.newUseCase(t)
			err := u.Store(context.Background(), file)

			require.ErrorIs(t, err, tc.err)

			if tc.requireU != nil {
				tc.requireU(t, u)
			}
		})
	}
}
