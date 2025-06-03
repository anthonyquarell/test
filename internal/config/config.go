package config

import (
	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

var Conf = struct {
	Namespace   string `env:"NAMESPACE" envDefault:"example.com"`
	Debug       bool   `env:"DEBUG" envDefault:"false"`
	GrpcPort    string `env:"GRPC_PORT" envDefault:"5050"`
	HttpPort    string `env:"HTTP_PORT" envDefault:"80"`
	WithMetrics bool   `env:"WITH_METRICS" envDefault:"false"`
	WithTracing bool   `env:"WITH_TRACING" envDefault:"false"`
	PgDsn       string `env:"PG_DSN"`
}{}

func init() {
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}
}
