package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_Get(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		file        model.File
		err         error
	}{
		{
			name: "success without filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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
			file: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			},
		},
		{
			name: "success without filter from other tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

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

				return u, model.FileFilter{}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId2,
			},
		},
		{
			name: "success without filter and without file in tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			},
		},
		{
			name: "success without filter and without tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			},
		},
		{
			name: "not found without filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey2,
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

				return u, model.FileFilter{}
			},
			err: fs_db.ErrNotFound,
		},
		{
			name: "success with txId filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			},
		},
		{
			name: "success with txId filter from other tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

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
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId2,
			},
		},
		{
			name: "success with txId filter and without file in tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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
					TxId: ptr.Ptr(testTxId2),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			},
		},
		{
			name: "success with txId filter and without tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			},
		},
		{
			name: "success with txId filter and without filter tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
			},
		},
		{
			name: "not found with txId filter and without txs",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			err: fs_db.ErrNotFound,
		},
		{
			name: "not found with txId filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			err: fs_db.ErrNotFound,
		},
		{
			name: "success with filter",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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
			file: model.File{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			},
		},
		{
			name: "success with filter from other tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				u := td.newUseCase()

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
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId2,
			},
		},
		{
			name: "success with filter and without file in tx",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
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
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId:      ptr.Ptr(testTxId2),
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			file: model.File{
				Key:       testKey,
				TxId:      testTxId2,
				ContentId: testContentId,
			},
		},
		{
			name: "not found with txId filter",
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
					Key:       testKey,
					TxId:      testTxId3,
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
					TxId:      ptr.Ptr(testTxId2),
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			err: fs_db.ErrNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, filter := tc.initUseCase(td)

			file, err := u.Get(context.Background(), testTxId, testKey, filter)

			require.ErrorIs(t, err, tc.err)
			requireEqualFile(t, tc.file, file)
		})
	}
}
