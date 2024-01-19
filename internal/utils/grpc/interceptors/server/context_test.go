package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestContextInterceptor(t *testing.T) {
	t.Parallel()

	resp, err := ContextInterceptor(testCtx, testRequest, nil, func(c context.Context, req any) (any, error) {
		require.Equal(t, testRequest, req)
		require.Equal(t, testTxId, model.GetTxId(c))

		return testResponse, nil
	})

	require.NoError(t, err)
	require.Equal(t, testResponse, resp)
}
