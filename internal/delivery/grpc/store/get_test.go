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

func TestImplementation_GetFile(t *testing.T) {
	t.Parallel()

	const (
		key     = "key"
		content = "content"
	)

	for _, tc := range []struct {
		name      string
		key       string
		prepare   prepareFunc
		checkResp func(t *testing.T, stream store.StoreV1_GetFileClient)
	}{
		{
			name: "success",
			key:  key,
			prepare: func(td *testDeps) {
				r := strings.NewReader(content)
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
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileClient) {
				resp, err := stream.Recv()
				require.NoError(t, err)
				require.Equal(t, key, resp.GetHeader().GetKey())

				var buf bytes.Buffer
				for {
					resp, err = stream.Recv()
					if errors.Is(err, io.EOF) {
						break
					}
					require.NoError(t, err)

					buf.Write(resp.GetChunk())
				}

				require.Equal(t, content, buf.String())
			},
		},
		{
			name: "Get error",
			key:  key,
			prepare: func(td *testDeps) {
				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileClient) {
				resp, err := stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "Send header error",
			key:  strings.Repeat("A", 501),
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Close().
					Return(nil).
					AnyTimes()

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileClient) {
				resp, err := stream.Recv()
				require.Equal(t, codes.ResourceExhausted, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "Read error",
			key:  key,
			prepare: func(td *testDeps) {
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
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileClient) {
				resp, err := stream.Recv()
				require.NoError(t, err)
				require.Equal(t, key, resp.GetHeader().GetKey())

				resp, err = stream.Recv()
				require.Equal(t, codes.Internal, status.Code(err))
				require.Nil(t, resp)
			},
		},
		{
			name: "Send chunk error",
			key:  key,
			prepare: func(td *testDeps) {
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
					Get(gomock.Any(), key).
					Return(td.r, nil)
			},
			checkResp: func(t *testing.T, stream store.StoreV1_GetFileClient) {
				resp, err := stream.Recv()
				require.NoError(t, err)
				require.Equal(t, key, resp.GetHeader().GetKey())

				resp, err = stream.Recv()
				require.Equal(t, codes.ResourceExhausted, status.Code(err))
				require.Nil(t, resp)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			stream, err := td.client.GetFile(context.Background(), &store.GetFileRequest{
				Key: tc.key,
			})

			require.NoError(t, err)
			tc.checkResp(t, stream)
		})
	}
}
