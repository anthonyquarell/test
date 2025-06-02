package config

import (
	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/yaml.v3"
)

var Conf = struct {
	Namespace               string                       `env:"NAMESPACE" envDefault:"example.com"`
	Debug                   bool                         `env:"DEBUG" envDefault:"false"`
	GrpcPort                string                       `env:"GRPC_PORT" envDefault:"5050"`
	HttpPort                string                       `env:"HTTP_PORT" envDefault:"80"`
	HttpCors                bool                         `env:"HTTP_CORS" envDefault:"false"`
	WithMetrics             bool                         `env:"WITH_METRICS" envDefault:"false"`
	WithTracing             bool                         `env:"WITH_TRACING" envDefault:"false"`
	JaegerAddress           string                       `env:"JAEGER_ADDRESS"`
	RedisUrl                string                       `env:"REDIS_URL"`
	RedisPsw                string                       `env:"REDIS_PSW"`
	RedisDb                 int                          `env:"REDIS_DB"`
	PgDsn                   string                       `env:"PG_DSN"`
	ShopGrpcUrl             string                       `env:"SHOP_GRPC_URL"`
	ShopGrpcUsername        string                       `env:"SHOP_GRPC_USERNAME"`
	ShopGrpcPassword        string                       `env:"SHOP_GRPC_PASSWORD"`
	ShopUsername            string                       `env:"SHOP_USERNAME"` // for auth in 1c
	ShopPassword            string                       `env:"SHOP_PASSWORD"` // for auth in 1c
	ShopHttpRoutesRaw       string                       `env:"SHOP_HTTP_ROUTES"`
	ShopHttpRoutes          map[string]map[string]string `env:"-"`
	NsiGrpcUrl              string                       `env:"NSI_GRPC_URL"`
	NsiGrpcUsername         string                       `env:"NSI_GRPC_USERNAME"`
	NsiGrpcPassword         string                       `env:"NSI_GRPC_PASSWORD"`
	JwtsGrpcUrl             string                       `env:"JWTS_GRPC_URL"`
	ShLinkUrl               string                       `env:"SH_LINK_URL"`
	ShLinkApiKey            string                       `env:"SH_LINK_API_KEY"`
	ShLinkTag               string                       `env:"SH_LINK_TAG"`
	QrLinkTemplate          string                       `env:"QR_LINK_TEMPLATE"` // https://example.com/{token}
	KaspiUrl                string                       `env:"KASPI_URL"`
	KaspiToken              string                       `env:"KASPI_TOKEN"`
	MbBrokerGrpcUrl         string                       `env:"MB_BROKER_GRPC_URL"`
	AbBrokerUrl             string                       `env:"AB_BROKER_URL"`
	AbBrokerToken           string                       `env:"AB_BROKER_TOKEN"`
	SmsGrpcUrl              string                       `env:"SMS_GRPC_URL"`
	SmsGrpcUsername         string                       `env:"SMS_GRPC_USERNAME"`
	SmsGrpcPassword         string                       `env:"SMS_GRPC_PASSWORD"`
	EventWriterGrpcUrl      string                       `env:"EVENT_WRITER_GRPC_URL"`
	EventWriterGrpcUsername string                       `env:"EVENT_WRITER_GRPC_USERNAME"`
	EventWriterGrpcPassword string                       `env:"EVENT_WRITER_GRPC_PASSWORD"`
	DefaultTerminalId       string                       `env:"DEFAULT_TERMINAL_ID"`
	TestShop                bool                         `env:"TEST_SHOP" envDefault:"false"`
}{}

func init() {
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}

	Conf.ShopHttpRoutes = make(map[string]map[string]string)
	if len(Conf.ShopHttpRoutesRaw) > 0 {
		err := yaml.Unmarshal([]byte(Conf.ShopHttpRoutesRaw), &Conf.ShopHttpRoutes)
		if err != nil {
			panic(err)
		}
	}
}
