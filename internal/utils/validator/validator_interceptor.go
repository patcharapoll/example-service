package validator

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor ...
func UnaryServerInterceptor(validator *CustomValidator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validator.Validate(req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
