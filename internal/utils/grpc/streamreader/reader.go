package streamreader

import "bytes"

//go:generate mockgen -source reader.go -destination mocks/mocks.go -typed true

type Request interface {
	GetChunk() []byte
}

type Stream interface {
	Recv() (Request, error)
}

type reader struct {
	stream Stream
	buf    bytes.Buffer
}

func New(stream Stream) *reader {
	return &reader{stream: stream}
}

func (r *reader) Read(p []byte) (int, error) {
	for len(p) > r.buf.Len() {
		req, err := r.stream.Recv()
		if err != nil {
			break
		}

		r.buf.Write(req.GetChunk())
	}

	return r.buf.Read(p)
}
