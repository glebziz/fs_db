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

func TestService_SetFile(t *testing.T) {
	t.Parallel()

	const (
		key     = "key"
		content = "content"
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.sStream.EXPECT().
					Context().
					Return(context.Background())

				td.sStream.EXPECT().
					Recv().
					Return(&store.SetFileRequest{
						Data: &store.SetFileRequest_Header{
							Header: &store.FileHeader{
								Key: key,
							},
						},
					}, nil)

				td.suc.EXPECT().
					Set(gomock.Any(), key, gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						var (
							counter int
							data    bytes.Buffer
						)
						for c := range contents {
							counter++
							_, err := io.Copy(&data, c.Reader)
							require.NoError(t, err)
						}

						require.Equal(t, content, data.String())
						require.Equal(t, 1, counter)

						return nil
					})

				td.sStream.EXPECT().
					Recv().
					Return(&store.SetFileRequest{
						Data: &store.SetFileRequest_Chunk{
							Chunk: []byte(content),
						},
					}, nil)

				td.sStream.EXPECT().
					Recv().
					Return(nil, io.EOF)

				td.sStream.EXPECT().
					Recv().
					Return(nil, io.EOF)

				td.sStream.EXPECT().
					SendAndClose(&store.SetFileResponse{}).
					Return(nil)
			},
		},
		{
			name: "Recv header error",
			prepare: func(td *testDeps) {
				td.sStream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "ErrHeaderNotFound error",
			prepare: func(td *testDeps) {
				td.sStream.EXPECT().
					Recv().
					Return(&store.SetFileRequest{}, nil)
			},
			errCode: codes.Internal,
		},
		{
			name: "Set error",
			prepare: func(td *testDeps) {
				td.sStream.EXPECT().
					Context().
					Return(context.Background())

				td.sStream.EXPECT().
					Recv().
					Return(&store.SetFileRequest{
						Data: &store.SetFileRequest_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "SendAndClose error",
			prepare: func(td *testDeps) {
				td.sStream.EXPECT().
					Context().
					Return(context.Background())

				td.sStream.EXPECT().
					Recv().
					Return(&store.SetFileRequest{
						Data: &store.SetFileRequest_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.sStream.EXPECT().
					SendAndClose(gomock.Any()).
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
			err := s.SetFile(td.sStream)
			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}

// func TestImplementation_SetFile_Success(t *testing.T) {
// 	t.Parallel()
//
// 	td := newTestDeps(t)
//
// 	td.suc.EXPECT().
// 		Set(gomock.Any(), testKey, gomock.Any()).
// 		Do(func(ctx context.Context, s string, contents model.Contents) error {
// 			var data bytes.Buffer
// 			for c := range contents {
// 				_, err := io.Copy(&data, c.Reader)
// 				require.NoError(t, err)
// 			}
//
// 			require.Equal(t, testContent, data.Bytes())
//
// 			return nil
// 		}).
// 		Return(nil)
//
// 	stream, err := td.client.SetFile(context.Background())
//
// 	require.NoError(t, err)
//
// 	err = stream.Send(&store.SetFileRequest{
// 		Data: &store.SetFileRequest_Header{
// 			Header: &store.FileHeader{
// 				Key: testKey,
// 			},
// 		},
// 	})
//
// 	require.NoError(t, err)
//
// 	err = stream.Send(&store.SetFileRequest{
// 		Data: &store.SetFileRequest_Chunk{
// 			Chunk: testContent,
// 		},
// 	})
//
// 	require.NoError(t, err)
//
// 	_, err = stream.CloseAndRecv()
//
// 	require.NoError(t, err)
// }
//
// func TestImplementation_SetFile_Error(t *testing.T) {
// 	t.Run("empty header", func(t *testing.T) {
// 		t.Parallel()
//
// 		td := newTestDeps(t)
//
// 		stream, err := td.client.SetFile(context.Background())
//
// 		require.NoError(t, err)
//
// 		err = stream.Send(&store.SetFileRequest{
// 			Data: &store.SetFileRequest_Chunk{
// 				Chunk: testContent,
// 			},
// 		})
//
// 		require.NoError(t, err)
//
// 		_, err = stream.CloseAndRecv()
//
// 		require.Error(t, err)
// 		require.Equal(t, codes.Internal, status.Convert(err).Code())
// 	})
//
// 	t.Run("upload", func(t *testing.T) {
// 		t.Parallel()
//
// 		td := newTestDeps(t)
//
// 		td.suc.EXPECT().
// 			Set(gomock.Any(), gomock.Any(), gomock.Any()).
// 			Return(assert.AnError)
//
// 		stream, err := td.client.SetFile(context.Background())
//
// 		require.NoError(t, err)
//
// 		err = stream.Send(&store.SetFileRequest{
// 			Data: &store.SetFileRequest_Header{
// 				Header: &store.FileHeader{
// 					Key: testKey,
// 				},
// 			},
// 		})
//
// 		require.NoError(t, err)
//
// 		err = stream.Send(&store.SetFileRequest{
// 			Data: &store.SetFileRequest_Chunk{
// 				Chunk: testContent,
// 			},
// 		})
//
// 		require.NoError(t, err)
//
// 		_, err = stream.CloseAndRecv()
//
// 		require.Error(t, err)
// 		require.Equal(t, codes.Internal, status.Convert(err).Code())
// 	})
// }
