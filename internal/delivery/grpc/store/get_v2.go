package store

import (
	"errors"
	"fmt"
	"io"

	errorsAdapter "github.com/glebziz/fs_db/internal/adapter/errors"
	"github.com/glebziz/fs_db/internal/adapter/seek_direction"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) GetFileV2(stream store.StoreV1_GetFileV2Server) error {
	req, err := stream.Recv()
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("stream recv: %w", err))
	}

	content, err := i.sUsecase.Get(stream.Context(), req.GetHeader().GetKey())
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("store usecase get: %w", err))
	}
	defer content.Close()

	size, err := getSize(content)
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("get size: %w", err))
	}

	err = stream.Send(&store.GetFileV2Response{
		Data: &store.GetFileV2Response_Size{
			Size: size,
		},
	})
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("send size: %w", err))
	}

	err = sendContent(stream, size, content)
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("send content: %w", err))
	}

	return nil
}

func sendContent(stream store.StoreV1_GetFileV2Server, size int64, content model.ReadSeekCloser) error {
	var pos model.Position

	pos.SetEnd(size)
	chunk := make([]byte, store.ChunkSize_MAX)
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("stream recv: %w", err)
		}

		seek := model.NewSeek(req.GetPos().GetOffset(), seek_direction.ConvertToWhence(req.GetPos().GetDir()))
		if pos.NeedSeek(seek) {
			_ = pos.Seek(seek)
			_, err = seek.Apply(content)
			if err != nil {
				return fmt.Errorf("seek apply: %w", err)
			}
		}

		n, err := content.Read(chunk)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}

		err = stream.Send(&store.GetFileV2Response{
			Data: &store.GetFileV2Response_Chunk{
				Chunk: chunk[:n],
			},
		})
		if err != nil {
			return fmt.Errorf("stream chunk send: %w", err)
		}

		_ = pos.Seek(model.Seek{
			Pos: int64(n),
			Dir: model.SeekCurrent,
		})
	}

	return nil
}

func getSize(s io.Seeker) (int64, error) {
	size, err := s.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, fmt.Errorf("seek end: %w", err)
	}

	_, err = s.Seek(0, io.SeekStart)
	if err != nil {
		return 0, fmt.Errorf("seek start: %w", err)
	}

	return size, nil
}
