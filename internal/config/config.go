package config

import (
	"github.com/farmani/sharebuy/pkg/cookie"
	"github.com/farmani/sharebuy/pkg/db"
	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/farmani/sharebuy/pkg/jwt"
	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/farmani/sharebuy/pkg/mailer"
	"github.com/farmani/sharebuy/pkg/nats"
	"github.com/farmani/sharebuy/pkg/redis"
)

type Config struct {
	App        *App               `koanf:"app"`
	Cors       *Cors              `koanf:"cors"`
	Limiter    *Limiter           `koanf:"limiter"`
	Logger     *logger.Config     `koanf:"logger"`
	Redis      *redis.Config      `koanf:"redis"`
	Nats       *nats.Config       `koanf:"nats"`
	Db         *db.Config         `koanf:"db"`
	Mailer     *mailer.Config     `koanf:"mailer"`
	Jwt        *jwt.Config        `koanf:"jwt"`
	Cookie     *cookie.Config     `koanf:"cookie"`
	Encryption *encryption.Config `koanf:"encryption"`
}
