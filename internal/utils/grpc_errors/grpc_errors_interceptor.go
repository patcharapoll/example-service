package grpcerrors

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor ...
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, Error(err)
		}
		return resp, nil
	}
}
