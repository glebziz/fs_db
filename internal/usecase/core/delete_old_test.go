package core

import (
	"context"
	"testing"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_DeleteOld(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		deleteFiles []model.File
	}{
		{
			name: "success",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
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
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			deleteFiles: []model.File{{
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId,
			}, {
				Key:       testKey,
				TxId:      testTxId,
				ContentId: testContentId2,
			}, {
				Key:       testKey2,
				TxId:      testTxId,
				ContentId: testContentId3,
			}},
		},
		{
			name: "success with only latest",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(td, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			deleteFiles: []model.File{},
		},
		{
			name: "success with all after seq",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()

				beforeSeq := sequence.Next()
				u.testStore(td, model.File{
					Key:       testKey,
					TxId:      testTxId,
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
					BeforeSeq: ptr.Ptr(beforeSeq),
				}
			},
			deleteFiles: []model.File{},
		},
		{
			name: "success with empty tx",
			initUseCase: func(td *testDeps) (*useCase, model.FileFilter) {
				u := td.newUseCase()
				u.testAddEmptyTx(td, testTxId, testKey, testKey2)

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			deleteFiles: []model.File{},
		},
		{
			name: "success with nil tx",
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

				return u, model.FileFilter{
					BeforeSeq: ptr.Ptr(sequence.Next()),
				}
			},
			deleteFiles: []model.File{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, filter := tc.initUseCase(td)
			deleteFiles := u.DeleteOld(context.Background(), testTxId, ptr.Val(filter.BeforeSeq))
			requireEqualFiles(t, tc.deleteFiles, deleteFiles)
		})
	}
}
