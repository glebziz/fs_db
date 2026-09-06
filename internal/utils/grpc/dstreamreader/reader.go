package dstreamreader

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source reader.go -destination mocks/mocks.go -typed true

type Request[T proto.Message] func(seek model.Seek) T

type Response interface {
	GetChunk() []byte
}

type Stream[Req proto.Message, Resp Response] interface {
	Send(req Req) error
	Recv() (Resp, error)
	CloseSend() error
}

type reader[Req proto.Message, Resp Response] struct {
	currentSeek model.Seek
	pos         model.Position
	stream      Stream[Req, Resp]
	req         Request[Req]
	buf         bytes.Buffer
}

func New[Req proto.Message, Resp Response](stream Stream[Req, Resp], req Request[Req], size int64) *reader[Req, Resp] {
	r := reader[Req, Resp]{
		stream: stream,
		req:    req,
	}

	r.pos.SetEnd(size)

	return &r
}

func (r *reader[Req, Resp]) Read(p []byte) (int, error) {
	for len(p) > r.buf.Len() {
		err := r.readChunk()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return 0, fmt.Errorf("read chunk: %w", err)
		}
	}

	return r.buf.Read(p)
}

func (r *reader[Req, Resp]) Close() error {
	r.buf.Reset()

	err := r.stream.CloseSend()
	if err != nil {
		return fmt.Errorf("close send: %w", err)
	}

	return nil
}

func (r *reader[Req, Resp]) Seek(offset int64, whence int) (int64, error) {
	r.buf.Reset()

	r.currentSeek = model.NewSeek(offset, whence)
	err := r.pos.Seek(r.currentSeek)
	if err != nil {
		return 0, err
	}

	return r.pos.Pos(), nil
}

func (r *reader[Req, Resp]) ReadAt(p []byte, off int64) (int, error) {
	_, err := r.Seek(off, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return r.Read(p)
}

func (r *reader[Req, Resp]) readChunk() error {
	defer func() {
		r.currentSeek = model.Seek{
			Dir: model.SeekCurrent,
		}
	}()

	err := r.stream.Send(r.req(r.currentSeek))
	if err != nil {
		return fmt.Errorf("stream send: %w", err)
	}

	resp, err := r.stream.Recv()
	if err != nil {
		return err
	}

	r.buf.Write(resp.GetChunk())

	_ = r.pos.Seek(model.Seek{
		Pos: int64(len(resp.GetChunk())),
		Dir: model.SeekCurrent,
	})

	return nil
}
