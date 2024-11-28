package streamwriter

import (
	"bytes"
)

type size interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Request[T any] func(p []byte) *T

type Stream[Req any, Resp any] interface {
	Send(req *Req) error
	CloseAndRecv() (*Resp, error)
}

type writer[SizeT size, Req any, Resp any] struct {
	stream    Stream[Req, Resp]
	req       Request[Req]
	buf       bytes.Buffer
	chunkSize SizeT
}

func New[SizeT size, Req any, Resp any](chunkSize SizeT, stream Stream[Req, Resp], req Request[Req]) *writer[SizeT, Req, Resp] {
	return &writer[SizeT, Req, Resp]{
		stream:    stream,
		req:       req,
		chunkSize: chunkSize,
	}
}

func (w *writer[SizeT, Req, Resp]) Write(p []byte) (int, error) {
	_, _ = w.buf.Write(p)
	buf := make([]byte, w.chunkSize)
	for w.buf.Len() >= int(w.chunkSize) {
		_, _ = w.buf.Read(buf)
		err := w.stream.Send(w.req(buf))
		if err != nil {
			return 0, err
		}
	}

	return len(p), nil
}

func (w *writer[SizeT, Req, Resp]) Close() error {
	data := w.buf.Bytes()
	if len(data) > 0 {
		err := w.stream.Send(w.req(w.buf.Bytes()))
		if err != nil {
			return err
		}
	}

	_, err := w.stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}
