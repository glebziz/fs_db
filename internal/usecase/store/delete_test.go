package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Delete_Success(t *testing.T) {
	t.Parallel()

	var (
		file = model.File{
			Id:         testId,
			Key:        testKey,
			ParentPath: testRootPath,
		}
	)

	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Get(gomock.Any(), testKey).
		Return(&file, nil)

	td.cRepo.EXPECT().
		Delete(gomock.Any(), file.GetPath()).
		Return(nil)

	td.fRepo.EXPECT().
		Delete(gomock.Any(), testKey).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Delete(context.Background(), testKey)

	require.NoError(t, err)
}

func TestUseCase_Delete_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "file get",
			prepare: func(td *testDeps) error {
				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content delete",
			prepare: func(td *testDeps) error {
				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(&model.File{
						Id:         testId,
						ParentPath: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file delete",
			prepare: func(td *testDeps) error {
				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(&model.File{
						Id:         testId,
						ParentPath: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(nil)

				td.fRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Delete(context.Background(), testKey)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
