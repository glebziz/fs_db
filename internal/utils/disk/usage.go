package disk

import (
	"context"
	"errors"
	"fmt"
	"os"

	diskUtil "github.com/shirou/gopsutil/disk"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	dirPerm = 0700
)

func Usage(ctx context.Context, path string) (*model.Stat, error) {
	st, err := diskUtil.UsageWithContext(ctx, path)
	if errors.Is(err, model.ErrPathNotFound) {
		err = os.MkdirAll(path, dirPerm)
		if err != nil {
			return nil, fmt.Errorf("mkdir all: %w", err)
		}

		st, err = diskUtil.UsageWithContext(ctx, path)
	}
	if err != nil {
		return nil, fmt.Errorf("usage with context: %w", err)
	}

	return &model.Stat{
		Path:  st.Path,
		Total: st.Total,
		Free:  st.Free,
		Used:  st.Used,
	}, nil
}
