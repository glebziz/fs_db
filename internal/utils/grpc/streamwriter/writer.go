package streamwriter

import (
	"bytes"
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/glebziz/fs_db/internal/model"
)

type size interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Request[T proto.Message] func(p []byte, seek model.Seek) T

type Stream[Req proto.Message, Resp proto.Message] interface {
	Send(req Req) error
	CloseAndRecv() (Resp, error)
}

type writer[SizeT size, Req proto.Message, Resp proto.Message] struct {
	currentSeek model.Seek
	pos         model.Position
	stream      Stream[Req, Resp]
	req         Request[Req]
	buf         bytes.Buffer
	chunkSize   SizeT
}

func New[SizeT size, Req proto.Message, Resp proto.Message](
	chunkSize SizeT, stream Stream[Req, Resp], req Request[Req],
) *writer[SizeT, Req, Resp] {
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
		err := w.send(buf)
		if err != nil {
			return 0, err
		}
	}

	_ = w.pos.Seek(model.Seek{
		Pos: int64(len(p)),
		Dir: model.SeekCurrent,
	})

	return len(p), nil
}

func (w *writer[SizeT, Req, Resp]) Close() error {
	err := w.flush()
	if err != nil {
		return err
	}

	_, err = w.stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func (w *writer[SizeT, Req, Resp]) Seek(offset int64, whence int) (int64, error) {
	err := w.flush()
	if err != nil {
		return 0, err
	}

	w.currentSeek = model.NewSeek(offset, whence)
	err = w.pos.Seek(w.currentSeek)
	if err != nil {
		return 0, err
	}

	return w.pos.Pos(), nil
}

func (w *writer[SizeT, Req, Resp]) WriteAt(p []byte, off int64) (int, error) {
	_, err := w.Seek(off, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return w.Write(p)
}

func (w *writer[SizeT, Req, Resp]) send(p []byte) error {
	defer func() {
		w.currentSeek = model.Seek{
			Dir: model.SeekCurrent,
		}
	}()

	return w.stream.Send(w.req(p, w.currentSeek))
}

func (w *writer[SizeT, Req, Resp]) flush() error {
	if w.buf.Len() == 0 {
		return nil
	}

	err := w.send(w.buf.Bytes())
	if err != nil {
		return err
	}

	w.buf.Reset()

	return nil
}
