package store

import (
	"bytes"
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

func TestService_SetFileV2(t *testing.T) {
	t.Parallel()

	const (
		key     = "key"
		content = "content"
	)

	for _, tc := range []struct {
		name      string
		prepare   prepareFunc
		checkResp func(t *testing.T, stream store.StoreV1_SetFileV2Client)
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					Set(gomock.Any(), key, gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						var (
							counter int
							data    bytes.Buffer
						)
						for c := range contents {
							counter++

							require.Equal(td.t, model.Seek{
								Dir: model.SeekStart,
							}, c.Seek)

							_, err := io.Copy(&data, c.Reader)
							require.NoError(td.t, err)
						}

						require.Equal(td.t, content, data.String())
						require.Equal(td.t, 1, counter)

						return nil
					})
			},
			checkResp: func(t *testing.T, stream store.StoreV1_SetFileV2Client) {
				err := stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				err = stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Chunk{
						Chunk: &store.DataChunk{
							Chunk: []byte(content),
							Pos:   &store.Seek{},
						},
					},
				})
				require.NoError(t, err)

				_, err = stream.CloseAndRecv()
				require.NoError(t, err)
			},
		},
		{
			name: "ErrHeaderNotFound error",
			prepare: func(td *testDeps) {
			},
			checkResp: func(t *testing.T, stream store.StoreV1_SetFileV2Client) {
				err := stream.Send(&store.SetFileV2Request{})
				require.NoError(t, err)

				err = stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Chunk{
						Chunk: &store.DataChunk{
							Chunk: []byte(content),
							Pos:   &store.Seek{},
						},
					},
				})
				require.NoError(t, err)

				_, err = stream.CloseAndRecv()
				require.Equal(t, codes.Internal, status.Code(err))
			},
		},
		{
			name: "Set error",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					Set(gomock.Any(), key, gomock.Any()).
					Return(assert.AnError)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_SetFileV2Client) {
				err := stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				err = stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Chunk{
						Chunk: &store.DataChunk{
							Chunk: []byte(content),
							Pos:   &store.Seek{},
						},
					},
				})
				require.NoError(t, err)

				_, err = stream.CloseAndRecv()
				require.Equal(t, codes.Internal, status.Code(err))
			},
		},
		{
			name: "Seek error",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					Set(gomock.Any(), key, gomock.Any()).
					Return(nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_SetFileV2Client) {
				err := stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				err = stream.Send(&store.SetFileV2Request{
					Data: &store.SetFileV2Request_Chunk{
						Chunk: &store.DataChunk{
							Chunk: []byte(content),
							Pos: &store.Seek{
								Offset: -1,
							},
						},
					},
				})
				require.NoError(t, err)

				_, err = stream.CloseAndRecv()
				require.Equal(t, codes.Internal, status.Code(err))
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			stream, err := td.client.SetFileV2(context.Background())

			require.NoError(t, err)
			tc.checkResp(t, stream)
		})
	}
}
