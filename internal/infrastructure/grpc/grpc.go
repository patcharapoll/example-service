package grpc

import (
	"context"
	"log"
	"net"

	"example-service/internal/config"
	"example-service/internal/controller"

	grpcErrors "example-service/internal/utils/grpc_errors"
	validatorUtils "example-service/internal/utils/validator"
	port_v1 "example-service/pkg/api/v1"
	grpc_health_v1 "example-service/pkg/grpc/health/v1"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Module ...
var Module = fx.Provide(
	NewHTTPGRPCServer,
	validatorUtils.NewCustomValidator,
)

// HTTPGRPCServer ...
type HTTPGRPCServer struct {
	server  *grpc.Server
	gateway controllerGateway
	config  *config.Configuration
}

type controllerGateway struct {
	fx.In
	HealthCtrl *controller.HealthZController
	PingCtrl   *controller.PingPongController
}

// NewHTTPGRPCServer ...
func NewHTTPGRPCServer(
	config *config.Configuration,
	gateway controllerGateway,
	validator *validatorUtils.CustomValidator,
) *HTTPGRPCServer {
	s := &HTTPGRPCServer{
		server: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				grpcErrors.UnaryServerInterceptor(),
				validatorUtils.UnaryServerInterceptor(validator),
			),
		),
		config:  config,
		gateway: gateway,
	}
	s.configure()
	return s
}

func (s *HTTPGRPCServer) configure() {
	grpc_health_v1.RegisterHealthServer(s.server, s.gateway.HealthCtrl)
	port_v1.RegisterPingPongServiceServer(s.server, s.gateway.PingCtrl)
}

// Start ...
func (s *HTTPGRPCServer) Start(_ context.Context) {
	go func() {
		listen, err := net.Listen("tcp", ":"+s.config.Port)
		if err != nil {
			log.Fatalln(err)
		}

		if err := s.server.Serve(listen); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("Listening and serving HTTP on", s.config.Port)
}

// Stop ...
func (s *HTTPGRPCServer) Stop(_ context.Context) {
	s.server.GracefulStop()
	log.Println("Server gracefully stopped")
}
