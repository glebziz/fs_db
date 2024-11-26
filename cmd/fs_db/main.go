package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/app"
	_ "github.com/glebziz/fs_db/internal/utils/log"
)

var (
	confFile string
)

func init() {
	flag.StringVar(&confFile, "config", "", "config file")
	flag.Parse()
}

func main() {
	log.Println("Start fs_db")

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	err := run(ctx)
	if err != nil {
		log.Fatalln("Run:", err)
	}

	log.Println("Stop fs_db")
}

func run(ctx context.Context) error {
	conf, err := config.ParseConfig(confFile)
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	a, err := app.New(ctx, conf)
	if err != nil {
		return fmt.Errorf("new app: %w", err)
	}

	err = a.Run(ctx)
	if err != nil {
		return fmt.Errorf("run app: %w", err)
	}

	err = a.Stop()
	if err != nil {
		return fmt.Errorf("stop app: %w", err)
	}

	return nil
}
