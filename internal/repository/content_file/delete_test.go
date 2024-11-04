package file

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/db/badger"
)

func TestRep_Delete(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Delete(append([]byte("fileContent/"), testId...)).
					Times(1).
					Return(nil)
			},
		},
		{
			name: "db delete error",
			prepare: func(td *testDeps) {
				td.qm.EXPECT().
					Delete(gomock.Any()).
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

			r := td.newRep()
			err := r.Delete(context.Background(), testId)

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRep_Delete_Int(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareIntFunc
	}{
		{
			name: "success",
			prepare: func(t *testing.T, p badger.Provider) {
				err := p.DB(context.Background()).Set([]byte(testId), []byte(testParent))
				require.NoError(t, err)
			},
		},
		{
			name:    "success without content",
			prepare: func(t *testing.T, p badger.Provider) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := newTestRep(t)
			tc.prepare(t, r.p)

			err := r.Delete(context.Background(), testId)
			require.NoError(t, err)
		})
	}
}
