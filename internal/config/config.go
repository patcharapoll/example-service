package config

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(
	NewConfiguration,
)

// Configuration ...
type Configuration struct {
	Port        string `env:"PORT" envDefault:"3001"`
	Environment string `env:"ENV" envDefault:"localhost"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
}

// NewConfiguration ...
func NewConfiguration() (*Configuration, error) {
	cfg := new(Configuration)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
