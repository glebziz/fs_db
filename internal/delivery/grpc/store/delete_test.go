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

func TestImplementation_DeleteFile_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.uc.EXPECT().
		Delete(gomock.Any(), testKey).
		Return(nil)

	_, err := td.client.DeleteFile(context.Background(), &store.DeleteFileRequest{
		Key: testKey,
	})

	require.NoError(t, err)
}

func TestImplementation_DeleteFile_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.uc.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	_, err := td.client.DeleteFile(context.Background(), &store.DeleteFileRequest{
		Key: testKey,
	})

	st := status.Convert(err)

	require.Error(t, err)
	require.Equal(t, codes.Internal, st.Code())
}
