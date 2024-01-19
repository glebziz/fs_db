package cleaner

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestCleaner_deleteContent_Success(t *testing.T) {
	for _, tc := range []struct {
		name       string
		contentIds []string
		prepare    prepareFunc
	}{
		{
			name: "empty content ids",
			prepare: func(td *testDeps) error {
				return nil
			},
		},
		{
			name:       "content files not found",
			contentIds: []string{testContentId},
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), []string{testContentId}).
					Return(nil, nil)

				return nil
			},
		},
		{
			name:       "success",
			contentIds: []string{testContentId},
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), []string{testContentId}).
					Return([]model.ContentFile{{
						Id:         testContentId,
						ParentPath: testParent,
					}}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), path.Join(testParent, testContentId)).
					Return(nil)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), []string{testContentId}).
					Return(nil)

				return nil
			},
		},
		{
			name:       "content not found",
			contentIds: []string{testContentId, testContentId2},
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), []string{testContentId, testContentId2}).
					Return([]model.ContentFile{{
						Id:         testContentId,
						ParentPath: testParent,
					}, {
						Id:         testContentId2,
						ParentPath: testParent,
					}}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), path.Join(testParent, testContentId)).
					Return(fs_db.NotFoundErr)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), path.Join(testParent, testContentId2)).
					Return(nil)

				td.cfRepo.EXPECT().
					Delete(gomock.Any(), []string{testContentId, testContentId2}).
					Return(nil)

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

			err := cl.deleteContent(testCtx, tc.contentIds)
			require.NoError(t, err)
		})
	}
}

func TestCleaner_deleteContent_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "content file get in",
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content delete",
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), gomock.Any()).
					Return([]model.ContentFile{{}}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content file delete",
			prepare: func(td *testDeps) error {
				td.cfRepo.EXPECT().
					GetIn(gomock.Any(), gomock.Any()).
					Return([]model.ContentFile{{}}, nil)

				td.cRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
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

			cl := td.newCleaner()

			err := cl.deleteContent(testCtx, []string{testContentId})
			require.ErrorIs(t, err, wantErr)
		})
	}
}
