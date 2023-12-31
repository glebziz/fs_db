package server

import (
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func StreamLoggingInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	service, method := splitFullMethodName(info.FullMethod)
	logger := slog.With(
		slog.Group(
			"grpc",
			slog.String("service", service),
			slog.String("method", method),
		),
	)

	logger.InfoContext(
		ss.Context(), "request",
	)

	err := handler(srv, ss)
	if err != nil {
		st := status.Convert(err)

		logger.ErrorContext(
			ss.Context(), "response error",
			slog.Group("error",
				slog.String("code", st.Code().String()),
				slog.String("message", st.Message()),
			),
		)

		return err
	}

	return nil
}
