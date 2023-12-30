package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Get_Success(t *testing.T) {
	t.Parallel()

	var (
		dir = model.Dir{
			Id:         testDirId,
			ParentPath: testRootPath,
		}
		file = model.File{
			Id:         testId,
			Key:        testKey,
			ParentPath: dir.GetPath(),
		}
		content = model.Content{
			Size:   testSize,
			Reader: testReader,
		}
	)
	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Get(gomock.Any(), testKey).
		Return(&file, nil)

	td.cRepo.EXPECT().
		Get(gomock.Any(), file.GetPath()).
		Return(&content, nil)

	uc := td.newUseCase()

	actContent, err := uc.Get(context.Background(), testKey)

	require.NoError(t, err)
	require.Equal(t, &content, actContent)
}

func TestUseCase_Get_Error(t *testing.T) {
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
			name: "content get",
			prepare: func(td *testDeps) error {
				td.fRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(&model.File{
						Id:         testId,
						Key:        testKey,
						ParentPath: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
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

			uc := td.newUseCase()

			actContent, err := uc.Get(context.Background(), testKey)

			require.ErrorIs(t, err, wantErr)
			require.Nil(t, actContent)
		})
	}
}
