package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Set(ctx context.Context, f model.File) error {
	data := make([]byte, fileLen(f))
	err := marshalFile(f, data)
	if err != nil {
		return fmt.Errorf("marshal file: %w, %v", err, f)
	}

	err = r.p.DB(ctx).Set(r.key(f.ContentId), data)
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}

	return nil
}
