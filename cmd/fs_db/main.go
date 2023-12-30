package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/glebziz/fs_db/config"
	storeService "github.com/glebziz/fs_db/internal/delivery/grpc/store"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/interceptors/server"
	"github.com/glebziz/fs_db/internal/utils/log"
	"github.com/glebziz/fs_db/pkg/inline/db"
)

var (
	confFile string
)

func init() {
	flag.StringVar(&confFile, "config", "", "config file")
	flag.Parse()
}

func main() {
	if confFile == "" {
		log.Fatalln("Config file must not be empty")
	}

	conf, err := config.ParseConfig(confFile)
	if err != nil {
		log.Fatalln("Parse config:", err)
	}

	cl, err := db.New(context.Background(), &conf.Storage)
	if err != nil {
		log.Fatalln("Inline client: ", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalln("Listen:", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			server.LoggingInterceptor,
		),
		grpc.ChainStreamInterceptor(
			server.StreamLoggingInterceptor,
		),
	)

	store.RegisterStoreV1Server(s, storeService.New(cl.GetUseCase()))
	log.Fatalln(s.Serve(lis))
}
