package root

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Get_Success(t *testing.T) {
	for _, tc := range []struct {
		name       string
		prepare    prepareFunc
		rootDirs   []string
		expRootMap model.RootMap
	}{
		{
			name: "empty rootDirs",
			prepare: func(td *testDeps) error {
				return nil
			},
			expRootMap: model.RootMap{},
		},
		{
			name: "success",
			prepare: func(td *testDeps) error {
				td.manager.EXPECT().
					Usage(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, path string) (*model.Stat, error) {
						switch path {
						case testRootPath:
							return &model.Stat{
								Path: path,
								Free: testRootFree,
							}, nil
						case testRootPath2:
							return &model.Stat{
								Path: path,
								Free: testRootFree2,
							}, nil
						default:
							return nil, assert.AnError
						}
					}).
					Times(2)

				td.repo.EXPECT().
					CountByParent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, path string) (uint64, error) {
						switch path {
						case testRootPath:
							return testRootCount, nil
						case testRootPath2:
							return testRootCount2, nil
						default:
							return 0, assert.AnError
						}
					}).
					Times(2)

				return nil
			},
			rootDirs: []string{testRootPath, testRootPath2},
			expRootMap: model.RootMap{
				testRootPath: &model.Root{
					Path:  testRootPath,
					Free:  testRootFree,
					Count: testRootCount,
				},
				testRootPath2: &model.Root{
					Path:  testRootPath2,
					Free:  testRootFree2,
					Count: testRootCount2,
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			_ = tc.prepare(td)

			uc := td.newUseCase(tc.rootDirs)

			rootMap, err := uc.Get(context.Background())

			require.NoError(t, err)
			require.Equal(t, tc.expRootMap, rootMap)
		})
	}
}

func TestUseCase_Get_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "usage",
			prepare: func(td *testDeps) error {
				td.manager.EXPECT().
					Usage(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "count by parent",
			prepare: func(td *testDeps) error {
				td.manager.EXPECT().
					Usage(gomock.Any(), gomock.Any()).
					Return(&model.Stat{}, nil)

				td.repo.EXPECT().
					CountByParent(gomock.Any(), gomock.Any()).
					Return(0, assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase([]string{testRootPath})

			rootMap, err := uc.Get(context.Background())

			require.ErrorIs(t, err, wantErr)
			require.Nil(t, rootMap)
		})
	}
}
