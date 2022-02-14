package config

import (
	"github.com/caarlos0/env/v6"
)

var Config = FromEnv()

type config struct {
	App struct {
		Environment string `env:"APP_ENV"              envDefault:"dev"`
		LogLevel    string `env:"LOG_LEVEL"            envDefault:"INFO"`
	}
	HTTP struct {
		Port int `env:"HTTP_PORT" envDefault:"3000"`
	}
	POSTGRES struct {
		Host     string `env:"POSTGRES_HOST"     envDefault:"0.0.0.0"`
		Port     int    `env:"POSTGRES_PORT"     envDefault:"5432"`
		User     string `env:"POSTGRES_USER"     envDefault:"postgres"`
		Pass     string `env:"POSTGRES_PASS"     envDefault:"123123123"`
		Database string `env:"POSTGRES_DATABASE" envDefault:"gshop"`
		Sslmode  string `env:"POSTGRES_SSLMODE"  envDefault:"disable"`
	}
	REDIS struct {
		Host string `env:"REDIS_HOST" envDefault:"localhost:6379"`
		DB   int    `env:"REDIS_DB"   envDefault:"0"`
	}
	JWT struct {
		SigningKey string `env:"SIGNING_KEY"            envDefault:"123def"`
		TokenTTL   int    `env:"TOKEN_TTL"              envDefault:"3600"`
	}
}

func FromEnv() *config {
	var c config

	if err := env.Parse(&c.App); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.HTTP); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.POSTGRES); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.REDIS); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.JWT); err != nil {
		panic(err)
	}

	return &c
}
