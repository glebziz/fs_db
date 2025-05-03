package os

import (
	"context"
	"os"

	"github.com/shirou/gopsutil/disk"
)

type Adapter struct{}

func (a Adapter) Create(_ context.Context, path string) (*os.File, error) {
	return os.Create(path)
}

func (a Adapter) Open(_ context.Context, path string) (*os.File, error) {
	return os.Open(path)
}

func (a Adapter) Remove(_ context.Context, path string) error {
	return os.Remove(path)
}

func (a Adapter) MkdirAll(_ context.Context, path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (a Adapter) ReadDir(_ context.Context, name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func (a Adapter) Usage(ctx context.Context, path string) (*disk.UsageStat, error) {
	return disk.UsageWithContext(ctx, path)
}
