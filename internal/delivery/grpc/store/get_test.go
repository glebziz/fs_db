package store

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestImplementation_GetFile_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.suc.EXPECT().
		Get(gomock.Any(), testKey).
		Return(&model.Content{
			Size:   testContentLen,
			Reader: testReader,
		}, nil)

	resp, err := td.client.GetFile(context.Background(), &store.GetFileRequest{
		Key: testKey,
	})

	require.NoError(t, err)

	data, err := resp.Recv()

	require.NoError(t, err)
	require.Equal(t, testKey, data.GetHeader().Key)
	require.Equal(t, testContentLen, data.GetHeader().Size)

	data, err = resp.Recv()

	require.NoError(t, err)
	require.Equal(t, testContent, data.GetChunk())
}

func TestImplementation_GetFile_Error(t *testing.T) {
	for _, tc := range []struct {
		name   string
		err    error
		reader io.ReadCloser
		code   codes.Code
	}{
		{
			name: "get",
			err:  assert.AnError,
			code: codes.Internal,
		},
		{
			name:   "read",
			reader: testErrReader,
			code:   codes.Internal,
		},
		{
			name:   "send",
			reader: io.NopCloser(strings.NewReader(strings.Repeat("A", 501))),
			code:   codes.ResourceExhausted,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			td.suc.EXPECT().
				Get(gomock.Any(), gomock.Any()).
				Return(&model.Content{
					Size:   testContentLen,
					Reader: tc.reader,
				}, tc.err)

			resp, err := td.client.GetFile(context.Background(), &store.GetFileRequest{
				Key: testKey,
			})

			require.NoError(t, err)

			data, err := resp.Recv()
			if tc.err == nil {
				require.NoError(t, err)
				require.Equal(t, testKey, data.GetHeader().Key)
				require.Equal(t, testContentLen, data.GetHeader().Size)

				data, err = resp.Recv()
			}

			require.Error(t, err)
			require.Equal(t, tc.code, status.Convert(err).Code())
			require.Nil(t, data)
		})
	}
}
