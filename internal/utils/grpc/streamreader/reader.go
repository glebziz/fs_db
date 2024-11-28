package streamreader

import "bytes"

//go:generate mockgen -source reader.go -destination mocks/mocks.go -typed true

type Request interface {
	GetChunk() []byte
}

type Stream[T Request] interface {
	Recv() (T, error)
}

type reader[T Request] struct {
	stream Stream[T]
	buf    bytes.Buffer
}

func New[T Request](stream Stream[T]) *reader[T] {
	return &reader[T]{
		stream: stream,
	}
}

func (r *reader[T]) Read(p []byte) (int, error) {
	for len(p) > r.buf.Len() {
		resp, err := r.stream.Recv()
		if err != nil {
			break
		}

		r.buf.Write(resp.GetChunk())
	}

	return r.buf.Read(p)
}
