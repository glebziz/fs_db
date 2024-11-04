package transaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Begin(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		txId    string
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Store(gomock.Any(), gomock.Cond(func(x any) bool {
						tx, ok := x.(model.Transaction)
						if !ok {
							return false
						}

						return tx.Id == testId && tx.IsoLevel == testIsoLevel
					})).
					Times(1).
					Return(nil)
			},
			txId: testId,
		},
		{
			name: "begin error",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
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
			txId, err := u.Begin(context.Background(), testIsoLevel)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.txId, txId)
		})
	}
}
