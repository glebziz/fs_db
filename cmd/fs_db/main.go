package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/glebziz/fs_db/config"
	storeService "github.com/glebziz/fs_db/internal/delivery/grpc/store"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/interceptors/server"
	_ "github.com/glebziz/fs_db/internal/utils/log"
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
			server.ContextInterceptor,
		),
		grpc.ChainStreamInterceptor(
			server.StreamLoggingInterceptor,
			server.ContextStreamInterceptor,
		),
	)

	store.RegisterStoreV1Server(s, storeService.New(cl.GetStoreUseCase(), cl.GetTxUseCase()))
	log.Fatalln(s.Serve(lis))
}
