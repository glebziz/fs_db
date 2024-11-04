package core

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	mock_core "github.com/glebziz/fs_db/internal/usecase/core/mocks"
)

func TestUseCase_DeleteTx(t *testing.T) {
	for _, tc := range []struct {
		name        string
		newUseCase  func(t *testing.T) *useCase
		deleteFiles []model.File
	}{
		{
			name: "success",
			newUseCase: func(t *testing.T) *useCase {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
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

				return u
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
			newUseCase: func(t *testing.T) *useCase {
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

				return u
			},
			deleteFiles: []model.File{},
		},
		{
			name: "success without tx",
			newUseCase: func(t *testing.T) *useCase {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

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

				return u
			},
			deleteFiles: []model.File{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := tc.newUseCase(t)

			u.DeleteTx(context.Background(), testTxId)

			_, ok := u.txStore.Get(testTxId)
			require.False(t, ok)

			require.Len(t, u.deleteFiles, len(tc.deleteFiles))

			slices.SortFunc(u.deleteFiles, func(a, b model.File) int {
				return strings.Compare(a.ContentId, b.ContentId)
			})
			slices.SortFunc(tc.deleteFiles, func(a, b model.File) int {
				return strings.Compare(a.ContentId, b.ContentId)
			})

			for i := range u.deleteFiles {
				requireEqualFiles(t, tc.deleteFiles[i], u.deleteFiles[i])
			}
		})
	}
}
