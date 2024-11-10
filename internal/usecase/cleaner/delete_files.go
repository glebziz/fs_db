package cleaner

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

const (
	deleteFilesAsyncCaller = "cleaner.DeleteFileAsync"

	chunkSize = 1_000
)

func (u *useCase) DeleteFilesAsync(ctx context.Context, files []model.File) {
	for cFiles := range slices.Chunk(files, chunkSize) {
		u.sender.Send(ctx, wpool.Event{
			Caller: deleteFilesAsyncCaller,
			Fn: func(ctx context.Context) error {
				err := u.DeleteFiles(ctx, cFiles)
				if err != nil {
					return fmt.Errorf("delete files: %w", err)
				}

				return nil
			},
		})
	}
}

func (u *useCase) DeleteFiles(ctx context.Context, files []model.File) error {
	if len(files) == 0 {
		return nil
	}

	var err error
	for _, file := range files {
		dErr := u.deleteFile(ctx, file)
		if dErr != nil {
			err = errors.Join(err, dErr)
		}
	}
	if err != nil {
		return fmt.Errorf("delete file: %w", err)
	}

	return nil
}

func (u *useCase) deleteFile(ctx context.Context, file model.File) error {
	cf, err := u.cfRepo.Get(ctx, file.ContentId)
	if errors.Is(err, fs_db.NotFoundErr) {
		return nil
	} else if err != nil {
		return fmt.Errorf("content file repo get: %w", err)
	}

	err = u.cRepo.Delete(ctx, cf.Path())
	if err != nil && !errors.Is(err, fs_db.NotFoundErr) {
		return fmt.Errorf("content repo delete: %w", err)
	}

	err = u.cfRepo.Delete(ctx, file.ContentId)
	if err != nil {
		return fmt.Errorf("content file repo delete: %w", err)
	}

	err = u.fRepo.Delete(ctx, file)
	if err != nil {
		return fmt.Errorf("file repo delete: %w", err)
	}

	return nil
}
