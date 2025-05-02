package store

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_GetFileV2(t *testing.T) {
	t.Parallel()

	const (
		zero int64 = iota
		offset

		key     = "key"
		content = "content"
	)

	for _, tc := range []struct {
		name      string
		prepare   prepareFunc
		checkResp func(t *testing.T, stream store.StoreV1_GetFileV2Client)
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				r := strings.NewReader(content)

				td.r.EXPECT().
					Seek(zero, io.SeekEnd).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						return r.Read(p)
					}).
					AnyTimes()

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), key).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.NoError(t, err)
				require.EqualValues(t, len(content), resp.GetSize())

				var data bytes.Buffer
				for {
					err = stream.Send(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Dir: store.Seek_SeekCurrent,
							},
						},
					})
					require.NoError(t, err)

					resp, err = stream.Recv()
					if errors.Is(err, io.EOF) {
						break
					}
					require.NoError(t, err)

					data.Write(resp.GetChunk())
				}

				require.Equal(t, content, data.String())
			},
		},
		{
			name: "success with seek",
			prepare: func(td *testDeps) {
				r := strings.NewReader(content)

				td.r.EXPECT().
					Seek(zero, io.SeekEnd).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)

				td.r.EXPECT().
					Seek(offset, io.SeekStart).
					Return(offset, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						return r.Read(p)
					}).
					AnyTimes()

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), key).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.NoError(t, err)
				require.EqualValues(t, len(content), resp.GetSize())

				var (
					counter int
					data    bytes.Buffer
				)
				for {
					seek := store.Seek{
						Dir: store.Seek_SeekCurrent,
					}

					if counter == 0 {
						seek.Offset = offset
						seek.Dir = store.Seek_SeekStart
					}

					err = stream.Send(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &seek,
						},
					})
					require.NoError(t, err)

					resp, err = stream.Recv()
					if errors.Is(err, io.EOF) {
						break
					}
					require.NoError(t, err)

					data.Write(resp.GetChunk())
					counter++
				}

				require.Equal(t, content, data.String())
			},
		},
		{
			name: "Get error",
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "seek end error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "seek start error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "seek error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.NoError(t, err)
				require.EqualValues(t, len(content), resp.GetSize())

				err = stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Pos{
						Pos: &store.Seek{
							Offset: offset,
							Dir:    store.Seek_SeekStart,
						},
					},
				})
				require.NoError(t, err)

				resp, err = stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "read error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, assert.AnError)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.NoError(t, err)
				require.EqualValues(t, len(content), resp.GetSize())

				err = stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Pos{
						Pos: &store.Seek{
							Dir: store.Seek_SeekCurrent,
						},
					},
				})
				require.NoError(t, err)

				resp, err = stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "send chunk error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, nil)

				r := strings.NewReader(strings.Repeat("A", 501))
				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						return r.Read(p)
					}).
					AnyTimes()

				td.r.EXPECT().
					Close().
					Return(nil).
					AnyTimes()

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileV2Client) {
				err := stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Header{
						Header: &store.FileHeader{
							Key: key,
						},
					},
				})
				require.NoError(t, err)

				resp, err := stream.Recv()
				require.NoError(t, err)
				require.EqualValues(t, len(content), resp.GetSize())

				err = stream.Send(&store.GetFileV2Request{
					Data: &store.GetFileV2Request_Pos{
						Pos: &store.Seek{
							Dir: store.Seek_SeekCurrent,
						},
					},
				})
				require.NoError(t, err)

				for {
					err = stream.Send(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Dir: store.Seek_SeekCurrent,
							},
						},
					})
					require.NoError(t, err)

					resp, err = stream.Recv()
					if err != nil {
						require.Equal(t, codes.ResourceExhausted, status.Code(err))
						break
					}
				}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			stream, err := td.client.GetFileV2(context.Background())

			require.NoError(t, err)
			tc.checkResp(t, stream)
		})
	}
}
