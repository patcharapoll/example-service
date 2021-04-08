package controller

import (
	"context"

	grpc_health_v1 "example-service/pkg/grpc/health/v1"
)

// HealthZController ...
type HealthZController struct{}

// NewHealthZController ...
func NewHealthZController() *HealthZController {
	return new(HealthZController)
}

// Check ...
func (*HealthZController) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
