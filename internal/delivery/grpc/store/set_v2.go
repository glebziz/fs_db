package store

import (
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	errorsAdapter "github.com/glebziz/fs_db/internal/adapter/errors"
	"github.com/glebziz/fs_db/internal/adapter/seek_direction"
	"github.com/glebziz/fs_db/internal/model/content"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) SetFileV2(stream store.StoreV1_SetFileV2Server) error {
	req, err := stream.Recv()
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("stream recv: %w", err))
	}

	header := req.GetHeader()
	if header == nil {
		return errorsAdapter.Error(fs_db.ErrHeaderNotFound)
	}

	rw, contents := content.New()
	rw.Add(1)
	defer rw.Wait()
	go func() {
		defer rw.Done()

		fmt.Println("Hello World")
		err = i.sUsecase.Set(stream.Context(), header.GetKey(), contents)
		if err != nil {
			rw.SetError(fmt.Errorf("store usecase set: %w", err))
		}
	}()

	err = getChunk(stream, rw)
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("get chunk: %w", err))
	}

	err = rw.Close()
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("store close: %w", err))
	}

	err = stream.SendAndClose(&store.SetFileV2Response{})
	if err != nil {
		return errorsAdapter.Error(fmt.Errorf("stream send and close: %w", err))
	}

	return nil
}

func getChunk(stream store.StoreV1_SetFileV2Server, w io.WriteSeeker) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("stream recv: %w", err)
		}

		chunk := req.GetChunk()
		_, err = w.Seek(chunk.GetPos().GetOffset(), seek_direction.ConvertToWhence(chunk.GetPos().GetDir()))
		if err != nil {
			return fmt.Errorf("writer seek: %w", err)
		}

		_, err = w.Write(chunk.GetChunk())
		if err != nil {
			return fmt.Errorf("writer write: %w", err)
		}
	}

	return nil
}
