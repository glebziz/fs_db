package dir

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Select_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.dRepo.EXPECT().
		GetRoots(gomock.Any()).
		Times(1).
		Return([]model.Root{{
			Path:  testRootPath,
			Count: 1,
		}, {
			Path: testRootPath2,
		}}, nil)

	td.dRepo.EXPECT().
		Create(gomock.Any(), model.Dir{
			Name: testName,
			Root: testRootPath2,
		}).
		Times(1).
		Return(nil)

	td.dRepo.EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return(model.Dirs{{
			Name: testName,
			Root: testRootPath2,
		}, {
			Name:  testName2,
			Root:  testRootPath,
			Count: testMaxCount,
		}}, nil)

	td.dRepo.EXPECT().
		Remove(gomock.Any(), model.Dir{
			Name:  testName2,
			Root:  testRootPath,
			Count: testMaxCount,
		}).
		Times(1).
		Return(nil)

	uc := td.newUseCase()

	dirs, err := uc.Get(context.Background())

	require.NoError(t, err)
	require.Equal(t, model.Dirs{{
		Name: testName,
		Root: testRootPath2,
	}}, dirs)
}

func TestUseCase_Select_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "get roots",
			prepare: func(td *testDeps) error {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "dir create",
			prepare: func(td *testDeps) error {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return([]model.Root{{}}, nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "dir get",
			prepare: func(td *testDeps) error {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return([]model.Root{{
						Count: 1,
					}}, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "dir remove",
			prepare: func(td *testDeps) error {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return([]model.Root{{
						Count: 1,
					}}, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Times(1).
					Return(model.Dirs{{
						Count: testMaxCount,
					}}, nil)

				td.dRepo.EXPECT().
					Remove(gomock.Any(), gomock.Any()).
					Times(1).
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

			dirs, err := uc.Get(context.Background())

			require.ErrorIs(t, err, wantErr)
			require.Nil(t, dirs)
		})
	}
}
