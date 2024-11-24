package streamwriter_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/utils/grpc/streamwriter"
	mock_streamwriter "github.com/glebziz/fs_db/internal/utils/grpc/streamwriter/mocks"
)

func TestWriter(t *testing.T) {
	for bufLen := 1; bufLen < 50; bufLen += 5 {
		t.Run(fmt.Sprintf("bufLen %d", bufLen), func(t *testing.T) {
			t.Parallel()

			var (
				err error
				n   int

				data   []byte
				chunks = [5][]byte{
					[]byte("hello"),
					[]byte("world"),
					[]byte("i am"),
					[]byte("gleb zhizhchenko"),
					[]byte("how are you?"),
				}

				allData     = bytes.Join(chunks[:], []byte(""))
				sendCounter = len(allData) / bufLen
			)

			ctrl := gomock.NewController(t)
			stream := mock_streamwriter.NewMockStreamTest(ctrl)

			stream.EXPECT().
				Send(gomock.Any()).
				DoAndReturn(func(r *streamwriter.TestRequest) error {
					data = append(data, r.P...)
					if sendCounter > 0 {
						require.Len(t, r.P, bufLen)
					} else {
						require.Len(t, r.P, len(allData)%bufLen)
					}
					sendCounter--
					return nil
				}).
				AnyTimes()

			stream.EXPECT().
				CloseAndRecv().
				Return(nil, nil)

			w := streamwriter.New(bufLen, stream, func(p []byte) *streamwriter.TestRequest {
				return &streamwriter.TestRequest{
					P: p,
				}
			})

			for _, chunk := range chunks {
				n, err = w.Write(chunk)

				require.NoError(t, err)
				require.Len(t, chunk, n)
			}

			err = w.Close()

			require.NoError(t, err)
			require.Equal(t, allData, data)
		})
	}
}

func TestWriter_Write(t *testing.T) {
	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			Send(gomock.Any()).
			Return(assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P: p,
			}
		})

		n, err := w.Write(bytes.Repeat([]byte("12"), chunkSize))

		require.ErrorIs(t, err, assert.AnError)
		require.Zero(t, n)
	})
}

func TestWriter_Close(t *testing.T) {
	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			Send(gomock.Any()).
			Return(assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P: p,
			}
		})

		n, err := w.Write([]byte("12"))
		require.NoError(t, err)
		require.Equal(t, 2, n)

		err = w.Close()

		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			CloseAndRecv().
			Return(nil, assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P: p,
			}
		})

		err := w.Close()
		require.ErrorIs(t, err, assert.AnError)
	})
}
