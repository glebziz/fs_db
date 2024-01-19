package store

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestImplementation_SetFile_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.suc.EXPECT().
		Set(gomock.Any(), testKey, gomock.Any()).
		Do(func(ctx context.Context, s string, content *model.Content) error {
			data, err := io.ReadAll(content.Reader)

			require.NoError(t, err)
			require.Equal(t, testContent, data)
			require.Equal(t, testContentLen, content.Size)

			return nil
		}).
		Return(nil)

	stream, err := td.client.SetFile(context.Background())

	require.NoError(t, err)

	err = stream.Send(&store.SetFileRequest{
		Data: &store.SetFileRequest_Header{
			Header: &store.SetFileHeader{
				Key:  testKey,
				Size: testContentLen,
			},
		},
	})

	require.NoError(t, err)

	err = stream.Send(&store.SetFileRequest{
		Data: &store.SetFileRequest_Chunk{
			Chunk: testContent,
		},
	})

	require.NoError(t, err)

	_, err = stream.CloseAndRecv()

	require.NoError(t, err)
}

func TestImplementation_SetFile_Error(t *testing.T) {
	t.Run("empty header", func(t *testing.T) {
		t.Parallel()

		td := newTestDeps(t)

		stream, err := td.client.SetFile(context.Background())

		require.NoError(t, err)

		err = stream.Send(&store.SetFileRequest{
			Data: &store.SetFileRequest_Chunk{
				Chunk: testContent,
			},
		})

		require.NoError(t, err)

		_, err = stream.CloseAndRecv()

		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Convert(err).Code())
	})

	t.Run("upload", func(t *testing.T) {
		t.Parallel()

		td := newTestDeps(t)

		td.suc.EXPECT().
			Set(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(assert.AnError)

		stream, err := td.client.SetFile(context.Background())

		require.NoError(t, err)

		err = stream.Send(&store.SetFileRequest{
			Data: &store.SetFileRequest_Header{
				Header: &store.SetFileHeader{
					Key:  testKey,
					Size: testContentLen,
				},
			},
		})

		require.NoError(t, err)

		err = stream.Send(&store.SetFileRequest{
			Data: &store.SetFileRequest_Chunk{
				Chunk: testContent,
			},
		})

		require.NoError(t, err)

		_, err = stream.CloseAndRecv()

		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Convert(err).Code())
	})
}
