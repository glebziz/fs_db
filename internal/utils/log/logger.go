package log

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dusted-go/logging/prettylog"
)

var (
	logger *slog.Logger
)

func init() {
	logger = slog.New(prettylog.NewHandler(nil))

	slog.SetDefault(logger)
}

func Fatalln(v ...any) {
	logger.Error(fmt.Sprintln(v...))
	os.Exit(1)
}
