package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	mock_core "github.com/glebziz/fs_db/internal/usecase/core/mocks"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_Get(t *testing.T) {
	for _, tc := range []struct {
		name       string
		newUseCase func(t *testing.T) (*useCase, model.FileFilter)
		file       model.File
		err        error
	}{
		{
			name: "success without filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey2,
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

				return u, model.FileFilter{}
			},
			err: fs_db.NotFoundErr,
		},
		{
			name: "success with txId filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
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
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

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
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId3,
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			err: fs_db.NotFoundErr,
		},
		{
			name: "not found with txId filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})

				return u, model.FileFilter{
					TxId: ptr.Ptr(testTxId2),
				}
			},
			err: fs_db.NotFoundErr,
		},
		{
			name: "success with filter",
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
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
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

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
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId5,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId3,
					ContentId: testContentId2,
					Seq:       sequence.Next(),
				})
				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId3,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey2,
					TxId:      testTxId,
					ContentId: testContentId4,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
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
			newUseCase: func(t *testing.T) (*useCase, model.FileFilter) {
				u := New(mock_core.NewMockfileRepository(gomock.NewController(t)))

				beforeTs := sequence.Next()
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId2,
					ContentId: testContentId,
					Seq:       sequence.Next(),
				})
				u.testStore(t, model.File{
					Key:       testKey,
					TxId:      testTxId3,
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
					TxId:      ptr.Ptr(testTxId2),
					BeforeSeq: ptr.Ptr(beforeTs),
				}
			},
			err: fs_db.NotFoundErr,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u, filter := tc.newUseCase(t)

			file, err := u.Get(context.Background(), testTxId, testKey, filter)

			require.ErrorIs(t, err, tc.err)
			requireEqualFiles(t, tc.file, file)
		})
	}
}
