package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_RollbackTx(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Rollback(gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "Rollback error",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Rollback(gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			_, err := s.RollbackTx(context.Background(), &store.RollbackTxRequest{})
			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}
