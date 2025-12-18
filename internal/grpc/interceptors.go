package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerLoggingInterceptor(log *Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start)
		statusCode := codes.OK
		if err != nil {
			statusCode = status.Code(err)
		}

		log.Info("gRPC unary call",
			"method", info.FullMethod,
			"status", statusCode.String(),
			"duration_ms", duration.Milliseconds(),
		)

		if err != nil {
			log.Error("gRPC unary call error",
				"method", info.FullMethod,
				"error", err.Error(),
			)
		}

		return resp, err
	}
}

func StreamServerLoggingInterceptor(log *Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		err := handler(srv, ss)

		duration := time.Since(start)
		statusCode := codes.OK
		if err != nil {
			statusCode = status.Code(err)
		}

		log.Info("gRPC stream call",
			"method", info.FullMethod,
			"status", statusCode.String(),
			"duration_ms", duration.Milliseconds(),
		)

		if err != nil {
			log.Error("gRPC stream call error",
				"method", info.FullMethod,
				"error", err.Error(),
			)
		}

		return err
	}
}
