//go:build !windows

package disk

import (
	"context"
	"errors"
	"fmt"
	"os"
	"syscall"

	diskUtil "github.com/shirou/gopsutil/disk"

	"github.com/glebziz/fs_db/internal/model"
)

func (d *disk) Usage(ctx context.Context, path string) (*model.Stat, error) {
	st, err := diskUtil.UsageWithContext(ctx, path)
	if errors.Is(err, syscall.ENOENT) {
		err = os.MkdirAll(path, 0700)
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
