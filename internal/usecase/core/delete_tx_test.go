package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func TestUseCase_DeleteTx(t *testing.T) {
	for _, tc := range []struct {
		name        string
		initUseCase initUseCaseFunc
		deleteFiles []model.File
	}{
		{
			name: "success",
			initUseCase: func(td *testDeps) (*UseCase, model.FileFilter) {
				var (
					files = []model.File{{
						Key:       testKey,
						TxId:      testTxId2,
						ContentId: testContentId,
						Seq:       sequence.Next(),
					}, {
						Key:       testKey,
						TxId:      testTxId,
						ContentId: testContentId2,
						Seq:       sequence.Next(),
					}, {
						Key:       testKey2,
						TxId:      testTxId,
						ContentId: testContentId3,
						Seq:       sequence.Next(),
					}, {
						Key:       testKey2,
						TxId:      testTxId2,
						ContentId: testContentId4,
						Seq:       sequence.Next(),
					}}
				)

				u := td.newUseCase()

				for _, file := range files {
					u.testStore(td, file)
				}

				return u, model.FileFilter{}
			},
			deleteFiles: []model.File{{
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
			name: "success with empty tx",
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

				return u, model.FileFilter{}
			},
			deleteFiles: []model.File{},
		},
		{
			name: "success without tx",
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
					TxId:      testTxId2,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{}
			},
			deleteFiles: []model.File{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			u, _ := tc.initUseCase(td)

			deleteFiles := u.DeleteTx(context.Background(), testTxId)

			for i := range deleteFiles {
				deleteFiles[i].Seq = 0
			}
			require.True(t, gomock.InAnyOrder(tc.deleteFiles).Matches(deleteFiles))

			_, ok := u.txStore.Get(testTxId)
			require.False(t, ok)
		})
	}
}
