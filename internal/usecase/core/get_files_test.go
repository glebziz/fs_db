package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_GetFiles(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		files       []model.File
		err         error
	}{
		{
			name: "success without filter",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
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

				return u, model.FileFilter{}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId4,
			}},
		},
		{
			name: "success without filter and without tx",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

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

				return u, model.FileFilter{}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId2,
			}},
		},
		{
			name: "empty without filter",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				return td.newUseCase(), model.FileFilter{}
			},
			files: nil,
		},
		{
			name: "success with txId filter",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey3,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId3,
					ContentId: testContentId6,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId5,
			}, {
				Key:       testKey3,
				TxId:      testTxId2,
				ContentId: testContentId4,
			}},
		},
		{
			name: "success with txId filter and without tx",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId3,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			}, {
				Key:       testKey2,
				TxId:      testTxId2,
				ContentId: testContentId3,
			}},
		},
		{
			name: "not found with txId filter and without txs",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
		},
		{
			name: "success with filter",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId6,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId:      ptr.Ptr(testTxId2),
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			}, {
				Key:       testKey2,
				TxId:      testTxId,
				ContentId: testContentId5,
			}},
		},
		{
			name: "success with filter and without tx",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
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

				return u, model.FileFilter{
					TxId:      ptr.Ptr(testTxId2),
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			files: []model.File{{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, filter := tc.initUseCase(td)

			files, err := u.GetFiles(context.Background(), testTxId, filter)

			require.ErrorIs(t, err, tc.err)
			requireEqualFiles(t, tc.files, files)
		})
	}
}
