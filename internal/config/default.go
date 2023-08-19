package config

import (
	"net/http"
	"time"

	"github.com/farmani/sharebuy/internal/repository"

	"github.com/farmani/sharebuy/pkg/cookie"
	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/farmani/sharebuy/pkg/jwt"
	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/farmani/sharebuy/pkg/mailer"
	"github.com/farmani/sharebuy/pkg/rdbms"

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
		Jwt: &jwt.Config{
			PrivatePem: `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIF0V3x7RkGyiVZGXCny8vtnBajmD2TOT2TkhounyUkBR
-----END PRIVATE KEY-----
`,
			PublicPem: `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEA1JsMvBD61BAYv8+JZtvex1K7Y1CgYeNnO9WMhgxNrv8=
-----END PUBLIC KEY-----
`,
			Expiration:        30 * 24 * time.Hour,
			RefreshExpiration: 3 * 30 * 24 * time.Hour,
			CookieTokenName:   "__Secure_token",
		},
		Redis: &redis.Config{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Db:       0,
		},
		Nats: &ncfg.Config{
			Url:                 nats.DefaultURL,
			Servers:             []string{},
			Nkey:                "",
			UserJWT:             "",
			User:                "",
			Password:            "",
			Timeout:             "5s",
			DrainTimeout:        "30s",
			FlusherTimeout:      "5s",
			ReconnectWait:       "5s",
			PingInterval:        "5s",
			MaxReconnects:       5,
			Token:               "",
			TokenHandler:        "",
			Pedantic:            false,
			Secure:              false,
			MaxPingsOutstanding: 2,
			AllowReconnect:      true,
			Verbose:             false,
			NoRandomize:         false,
			NoEcho:              false,
			Name:                "",
			Compression:         true,
		},
		Db: &rdbms.Config{
			Dsn:          "postgres://sharebuy-user:sharebuy-pass@localhost:5432/sharebuy?sslmode=disable",
			Host:         "postgres",
			Port:         "5432",
			Username:     "sharebuy-user",
			Password:     "sharebuy-pass",
			Database:     "sharebuy",
			MaxOpenConns: 64,
			MaxIdleConns: 64,
			MaxIdleTime:  "15m",
		},
		Repository: &repository.Config{},
		Mailer: &mailer.Config{
			Host:         "smtp.mailtrap.io",
			Port:         2525,
			Username:     "info",
			Password:     "info",
			ResendAPIKey: "re_KV12BrjV_56WDs6qW17AqAhAeEXh4o9ZB",
			Sender:       "noreply@mail.sharebuy.bid",
		},
		Logger: &logger.Config{
			Level:    "debug",
			Path:     "logs/error.log",
			Env:      "development",
			Encoding: "console",
		},
		Cookie: &cookie.Config{
			Domain:   "localhost",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Expires:  24 * time.Hour,
			MaxAge:   86400,
		},
		Encryption: &encryption.Config{
			Key:       []byte("39kQ2y7BgQQOXAzlUY6hnSqmQdRFH3Yy"),
			Algorithm: "AES-256",
		},
	}
}
