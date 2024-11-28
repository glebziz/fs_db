package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/model/transactor"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_UpdateTx(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		requireU    requireUseCaseFunc
		deleteFiles []model.File
		err         error
	}{
		{
			name: "success without filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.AnyOf(gomock.Cond(func(x any) bool {
						file, ok := x.(model.File)
						if !ok {
							return false
						}

						file.Seq = 0

						return gomock.Eq(model.File{
							Key:       testKey,
							TxId:      testTxId2,
							ContentId: testContentId5,
						}).Matches(file)
					}), gomock.Cond(func(x any) bool {
						file, ok := x.(model.File)
						if !ok {
							return false
						}

						file.Seq = 0

						return gomock.Eq(model.File{
							Key:       testKey2,
							TxId:      testTxId2,
							ContentId: testContentId3,
						}).Matches(file)
					}))).
					Times(2).
					Return(nil)

				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId5,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
				}, tx.File(testKey2).Latest())
			},
			deleteFiles: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
			}},
		},
		{
			name: "success with filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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

				td.fileRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(len(files)).
					DoAndReturn(func(_ context.Context, file model.File) error {
						f, ok := files[file.ContentId]
						require.True(t, ok)
						requireEqualFile(t, f, file)
						delete(files, file.ContentId)
						return nil
					})

				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId3,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success without newTx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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

				td.fileRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})
				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(len(files)).
					DoAndReturn(func(_ context.Context, file model.File) error {
						f, ok := files[file.ContentId]
						require.True(t, ok)
						requireEqualFile(t, f, file)
						delete(files, file.ContentId)
						return nil
					})

				u := td.newUseCase()

				beforeTs := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success without oldTx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				beforeTs := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "success with empty oldTx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testAddEmptyTx(td, testTxId, testKey, testKey2)
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
		},
		{
			name: "serialization error",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				beforeTs := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey2).Latest())
			},
			deleteFiles: []model.File{{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId3,
			}},
			err: fs_db.ErrTxSerialization,
		},
		{
			name: "file repo error",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				td.fileRepo.EXPECT().
					RunTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
						return fn(ctx)
					})

				call := td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				td.fileRepo.EXPECT().
					Set(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError).
					After(call)

				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			requireU: func(t *testing.T, u *UseCase) {
				_, ok := u.txStore.Get(testTxId)
				require.False(t, ok)

				tx, ok := u.txStore.Get(testTxId2)
				require.True(t, ok)

				requireEqualFile(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
				}, tx.File(testKey).Latest())
				requireEqualFile(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId2,
				}, tx.File(testKey2).Latest())
			},
			deleteFiles: []model.File{{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId3,
			}},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, filter := tc.initUseCase(td)
			deleteFiles, err := u.UpdateTx(context.Background(), testTxId, testTxId2, filter)

			require.ErrorIs(t, err, tc.err)
			requireEqualFiles(t, tc.deleteFiles, deleteFiles)
			tc.requireU(t, u)
		})
	}
}
