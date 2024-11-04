package log

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
}

func Fatalln(v ...any) {
	logger.Error(fmt.Sprintln(v...))
	os.Exit(1)
}
