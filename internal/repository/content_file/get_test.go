package file

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		cf      model.ContentFile
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Get(append([]byte("fileContent/"), testId...)).
					Times(1).
					Return([]byte(testParent), nil)
			},
			cf: model.ContentFile{
				Id:     testId,
				Parent: testParent,
			},
		},
		{
			name: "db get error",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Get(gomock.Any()).
					Times(1).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRep()
			cf, err := r.Get(context.Background(), testId)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.cf, cf)
		})
	}
}

func TestRep_Get_Int(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareIntFunc
		cf      model.ContentFile
		err     error
	}{
		{
			name: "success",
			prepare: func(t *testing.T, p badger.Provider) {
				err := p.DB(context.Background()).Set(append([]byte("fileContent/"), testId...), []byte(testParent))
				require.NoError(t, err)
			},
			cf: model.ContentFile{
				Id:     testId,
				Parent: testParent,
			},
		},
		{
			name:    "content not found",
			prepare: func(t *testing.T, p badger.Provider) {},
			err:     fs_db.ErrNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := newTestRep(t)
			tc.prepare(t, r.p)

			cf, err := r.Get(context.Background(), testId)
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.cf, cf)
		})
	}
}
