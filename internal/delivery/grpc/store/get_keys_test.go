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

func TestImplementation_GetKeys(t *testing.T) {
	t.Parallel()

	const (
		key = "key"
	)

	for _, tc := range []struct {
		name    string
		prepare func(td *testDeps)
		resp    *store.GetKeysResponse
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					GetKeys(gomock.Any()).
					Return([]string{key}, nil)
			},
			resp: &store.GetKeysResponse{
				Keys: []string{key},
			},
		},
		{
			name: "get keys error",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					GetKeys(gomock.Any()).
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			resp, err := s.GetKeys(context.Background(), &store.GetKeysRequest{})

			require.Equal(t, tc.errCode, status.Code(err))
			require.Equal(t, tc.resp, resp)
		})
	}
}
