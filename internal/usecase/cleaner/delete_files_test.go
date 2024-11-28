package cleaner

import (
	"context"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

func TestUseCase_DeleteFilesAsync(t *testing.T) {
	for _, tc := range []struct {
		name    string
		files   []model.File
		prepare prepareFunc
	}{
		{
			name:  "success",
			files: []model.File{{}},
			prepare: func(td *testDeps) {
				td.sender.EXPECT().
					Send(gomock.Any(), gomock.Cond(func(x any) bool {
						e, ok := x.(wpool.Event)
						if !ok {
							return false
						}

						return e.Caller == deleteFilesAsyncCaller &&
							e.Fn != nil
					})).
					Do(func(ctx context.Context, e wpool.Event) {
						e.Fn(ctx)
					}).
					Times(1)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.ContentFile{}, fs_db.ErrNotFound)

				td.db.EXPECT().
					GC().
					Times(1)
			},
		},
		{
			name:  "Get content file error",
			files: []model.File{{}},
			prepare: func(td *testDeps) {
				td.sender.EXPECT().
					Send(gomock.Any(), gomock.Any()).
					Do(func(ctx context.Context, e wpool.Event) {
						e.Fn(ctx)
					}).
					Times(1)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.ContentFile{}, assert.AnError)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			u := td.newUseCase()
			u.DeleteFilesAsync(context.Background(), tc.files)
		})
	}
}

func TestUseCase_DeleteFiles(t *testing.T) {
	for _, tc := range []struct {
		name    string
		files   []model.File
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			files: []model.File{{
				ContentId: testContentId,
			}, {
				ContentId: testContentId2,
			}, {
				ContentId: testContentId3,
			}},
			prepare: func(td *testDeps) {
				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId,
						Parent: testParent,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), path.Join(testParent, testContentId)).
					Times(1).
					Return(nil)

				td.dRepo.EXPECT().
					Add(gomock.Any(), model.Dir{
						Name: testParent,
						Root: ".",
					}).
					Times(1).
					Return(nil)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), testContentId).
					Times(1).
					Return(nil)

				td.fRepo.EXPECT().
					Delete(gomock.Any(), model.File{
						ContentId: testContentId,
					}).
					Times(1).
					Return(nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId2).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId2,
						Parent: testParent2,
					}, nil)

				td.dRepo.EXPECT().
					Add(gomock.Any(), model.Dir{
						Name: testParent2,
						Root: ".",
					}).
					Times(1).
					Return(nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), path.Join(testParent2, testContentId2)).
					Times(1).
					Return(fs_db.ErrNotFound)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), testContentId2).
					Times(1).
					Return(nil)

				td.fRepo.EXPECT().
					Delete(gomock.Any(), model.File{
						ContentId: testContentId2,
					}).
					Times(1).
					Return(nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId3).
					Times(1).
					Return(model.ContentFile{}, fs_db.ErrNotFound)

				td.db.EXPECT().
					GC().
					Times(1)
			},
		},
		{
			name:    "empty files",
			files:   nil,
			prepare: func(td *testDeps) {},
		},
		{
			name: "content file get error and content delete error",
			files: []model.File{{
				ContentId: testContentId,
			}, {
				ContentId: testContentId2,
			}},
			prepare: func(td *testDeps) {
				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Times(1).
					Return(model.ContentFile{}, assert.AnError)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId2).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId2,
						Parent: testParent2,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "dir repo add error",
			files: []model.File{{
				ContentId: testContentId,
			}},
			prepare: func(td *testDeps) {
				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId,
						Parent: testParent,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				td.dRepo.EXPECT().
					Add(gomock.Any(), gomock.Any()).
					Times(1).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "content file delete error and file delete error",
			files: []model.File{{
				ContentId: testContentId,
			}, {
				ContentId: testContentId2,
			}},
			prepare: func(td *testDeps) {
				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId2,
						Parent: testParent,
					}, nil)

				td.cfRepo.EXPECT().
					Get(gomock.Any(), testContentId2).
					Times(1).
					Return(model.ContentFile{
						Id:     testContentId2,
						Parent: testParent2,
					}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(2).
					Return(nil)

				td.dRepo.EXPECT().
					Add(gomock.Any(), gomock.Any()).
					Times(2).
					Return(nil)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), testContentId).
					Times(1).
					Return(assert.AnError)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), testContentId2).
					Times(1).
					Return(nil)

				td.fRepo.EXPECT().
					Delete(gomock.Any(), model.File{
						ContentId: testContentId2,
					}).
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

			u := td.newUseCase()
			err := u.DeleteFiles(context.Background(), tc.files)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
