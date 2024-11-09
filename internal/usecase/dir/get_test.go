package dir

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Select(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		dirs    model.Dirs
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
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
						Name: testNewName,
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

				td.dRepo.EXPECT().
					Create(gomock.Any(), model.Dir{
						Name: testNewName,
						Root: testRootPath,
					}).
					Times(1).
					Return(nil)
			},
			dirs: model.Dirs{{
				Name: testName,
				Root: testRootPath2,
			}, {
				Name: testNewName,
				Root: testRootPath,
			}},
		},
		{
			name: "get roots error",
			prepare: func(td *testDeps) {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "dir create error",
			prepare: func(td *testDeps) {
				td.dRepo.EXPECT().
					GetRoots(gomock.Any()).
					Times(1).
					Return([]model.Root{{}}, nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "dir get error",
			prepare: func(td *testDeps) {
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
			},
			err: assert.AnError,
		},
		{
			name: "dir remove err",
			prepare: func(td *testDeps) {
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
			},
			err: assert.AnError,
		},
		{
			name: "second dir create error",
			prepare: func(td *testDeps) {
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
					Return(nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			uc := td.newUseCase()

			dirs, err := uc.Get(context.Background())

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.dirs, dirs)
		})
	}
}
