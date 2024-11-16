package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	TxIdKey = "mdTxIdKey"
)

func ContextInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		txId := md.Get(TxIdKey)
		if len(txId) > 0 {
			ctx = model.StoreTxId(ctx, txId[0])
		}
	}

	return handler(ctx, req)
}
