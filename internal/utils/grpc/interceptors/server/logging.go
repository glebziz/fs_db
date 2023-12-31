package server

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const unknownValue = "unknown"

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	service, method := splitFullMethodName(info.FullMethod)
	logger := slog.With(
		slog.Group(
			"grpc",
			slog.String("service", service),
			slog.String("method", method),
		),
	)

	logger.InfoContext(
		ctx, "request",
		"body", req,
	)

	resp, err = handler(ctx, req)
	if err != nil {
		st := status.Convert(err)

		logger.ErrorContext(
			ctx, "response error",
			slog.Group("error",
				slog.String("code", st.Code().String()),
				slog.String("message", st.Message()),
			),
		)

		return nil, err
	}

	logger.InfoContext(
		ctx, "response",
		"body", resp,
	)

	return resp, nil
}

func splitFullMethodName(fullMethod string) (string, string) {
	fullMethod = strings.TrimPrefix(fullMethod, "/") // remove leading slash
	if i := strings.Index(fullMethod, "/"); i >= 0 {
		return fullMethod[:i], fullMethod[i+1:]
	}
	return unknownValue, unknownValue
}
