package cleaner

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestCleaner_deleteLines_Success(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "success",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(&model.Transaction{
						CreateTs: testTxCreateTs,
					}, nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), model.MainTxId, &model.FileFilter{
						BeforeTs: ptr.Ptr(testTxCreateTs),
					}).
					Return(nil, nil)

				return nil
			},
		},
		{
			name: "tx not found",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(nil, fs_db.TxNotFoundErr)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), model.MainTxId, gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, f *model.FileFilter) ([]string, error) {
						require.NotNil(t, f)
						require.True(t, f.BeforeTs.Second()-time.Now().Second() < 5)

						return nil, nil
					})

				return nil
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			_ = tc.prepare(td)

			cl := td.newCleaner()

			err := cl.deleteLines(testCtx)
			require.NoError(t, err)
		})
	}
}

func TestCleaner_deleteLines_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "tx repo oldest",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file repo hard delete",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(nil, fs_db.TxNotFoundErr)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "delete content",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Oldest(gomock.Any()).
					Return(nil, fs_db.TxNotFoundErr)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]string{testContentId}, nil)

				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			cl := td.newCleaner()

			err := cl.deleteLines(testCtx)
			require.ErrorIs(t, err, wantErr)
		})
	}
}
