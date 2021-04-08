package container

import (
	"context"
	"example-service/internal/config"
	"example-service/internal/controller"

	"example-service/internal/infrastructure/database"
	"example-service/internal/infrastructure/grpc"

	"go.uber.org/fx"
)

// Container ...
type Container struct{}

// NewContainer ...
func NewContainer() *Container {
	return new(Container)
}

func (c *Container) configure() []fx.Option {
	return []fx.Option{
		controller.Module,
		config.Module,
		grpc.Module,
		database.Module,
	}
}

func runApplication(
	lifecycle fx.Lifecycle,
	server *grpc.HTTPGRPCServer,
	db *database.DB,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.Start(ctx)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				server.Stop(ctx)
				// db.Close()
				return nil
			},
		},
	)
}

// Run ...
func (c *Container) Run() {
	options := append(c.configure(), fx.Invoke(runApplication))
	fx.New(
		options...,
	).Run()
}
