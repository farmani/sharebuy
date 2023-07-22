package config

import (
	"github.com/farmani/sharebuy/pkg/db"
	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/farmani/sharebuy/pkg/mailer"

	ncfg "github.com/farmani/sharebuy/pkg/nats"
	"github.com/farmani/sharebuy/pkg/redis"
	"github.com/nats-io/nats.go"
)

func Default() *Config {
	return &Config{
		App: &App{
			Port:    4000,
			Env:     "development",
			Version: "1.0.0",
		},
		Cors: &Cors{
			TrustedOrigins: []string{"http://localhost:3000"},
		},
		Limiter: &Limiter{
			Rps:     2,
			Burst:   4,
			Enabled: true,
		},
		Redis: &redis.Config{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Db:       0,
		},
		Nats: &ncfg.Config{
			Url:            nats.DefaultURL,
			Servers:        []string{},
			Nkey:           "",
			UserJWT:        "",
			User:           "",
			Password:       "",
			Timeout:        "5s",
			DrainTimeout:   "30s",
			FlusherTimeout: "5s",
			MaxReconnects:  5,
			ReconnectWait:  5,
			PingInterval:   "5s",
			Token:          "",
			TokenHandler:   "",
			Pedantic:       false,
			Secure:         false,
			MaxPingsOut:    2,
			AllowReconnect: true,
			Verbose:        false,
			NoRandomize:    false,
			NoEcho:         false,
			Name:           "",
			Compression:    0,
		},
		Db: &db.Config{
			Dsn: "postgres://postgres:ringsport@localhost/sharebuy?sslmode=disable",
		},
		Mailer: &mailer.Config{
			Host:     "smtp.mailtrap.io",
			Port:     2525,
			Username: "info",
			Password: "info",
		},
		Logger: &logger.Config{
			Level:    "debug",
			Path:     "logs/error.log",
			Env:      "development",
			Encoding: "console",
		},
	}
}
