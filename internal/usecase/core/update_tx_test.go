package core

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/model/transactor"
	mock_core "github.com/glebziz/fs_db/internal/usecase/core/mocks"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_UpdateTx(t *testing.T) {
	for _, tc := range []struct {
		name       string
		newUseCase func(t *testing.T) (*useCase, model.FileFilter)
		requireU   func(t *testing.T, u *useCase)
		err        error
	}{
		{
			name: "success without filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				var (
					files = map[string]model.File{
						testContentId5: {
							Key:       testKey,
							TxId:      testTxId2,
							ContentId: testContentId5,
						},
						testContentId3: {
							Key:       testKey2,
							TxId:      testTxId2,
							ContentId: testContentId3,
						},
					}
				)

				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(len(files)).
					DoAndReturn(func(_ context.Context, file model.File) error {
						f, ok := files[file.ContentId]
						require.True(t, ok)
						requireEqualFiles(t, f, file)
						delete(files, file.ContentId)
						return nil
					})

				u := New(fRepo)

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId5,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
				}, tx.File(testKey2).Latest())

				require.Len(t, u.deleteFiles, 1)
				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				}, u.deleteFiles[0])
			},
		},
		{
			name: "success with filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				var (
					files = map[string]model.File{
						testContentId3: {
							Key:       testKey,
							TxId:      testTxId2,
							ContentId: testContentId3,
						},
						testContentId4: {
							Key:       testKey2,
							TxId:      testTxId2,
							ContentId: testContentId4,
						},
					}
				)

				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(len(files)).
					DoAndReturn(func(_ context.Context, file model.File) error {
						f, ok := files[file.ContentId]
						require.True(t, ok)
						requireEqualFiles(t, f, file)
						delete(files, file.ContentId)
						return nil
					})

				u := New(fRepo)

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId3,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success without newTx",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				var (
					files = map[string]model.File{
						testContentId: {
							Key:       testKey,
							TxId:      testTxId2,
							ContentId: testContentId,
						},
						testContentId2: {
							Key:       testKey2,
							TxId:      testTxId2,
							ContentId: testContentId2,
						},
					}
				)

				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(len(files)).
					DoAndReturn(func(_ context.Context, file model.File) error {
						f, ok := files[file.ContentId]
						require.True(t, ok)
						requireEqualFiles(t, f, file)
						delete(files, file.ContentId)
						return nil
					})

				u := New(fRepo)

				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success without oldTx",
			newUseCase: func(*testing.T) (*useCase, model.FileFilter) {
				u := New(nil)

				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success with empty oldTx",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testAddEmptyTx(t, testTxId, testKey, testKey2)
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "serialization error",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey2).Latest())

				slices.SortFunc(u.deleteFiles, func(a, b model.File) int {
					return strings.Compare(a.ContentId, b.ContentId)
				})

				require.Len(t, u.deleteFiles, 2)
				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, u.deleteFiles[0])
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
				}, u.deleteFiles[1])
			},
			err: fs_db.TxSerializationErr,
		},
		{
			name: "file repo error",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				fRepo := mock_core.NewMockfileRepository(gomock.NewController(t))
				fRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})

				call := fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				fRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError).
					After(call)

				u := New(fRepo)

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			requireU: func(t *testing.T, u *useCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFiles(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey).Latest())
				requireEqualFiles(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u, filter := tc.newUseCase(t)
			err := u.UpdateTx(context.Background(), testTxId, testTxId2, filter)

			require.ErrorIs(t, err, tc.err)
			tc.requireU(t, u)
		})
	}
}
