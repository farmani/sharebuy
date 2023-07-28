package config

import (
	"github.com/farmani/sharebuy/internal/repository"
	"github.com/farmani/sharebuy/pkg/cookie"
	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/farmani/sharebuy/pkg/jwt"
	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/farmani/sharebuy/pkg/mailer"
	"github.com/farmani/sharebuy/pkg/nats"
	"github.com/farmani/sharebuy/pkg/rdbms"
	"github.com/farmani/sharebuy/pkg/redis"
)

type Config struct {
	App        *App               `koanf:"app"`
	Cors       *Cors              `koanf:"cors"`
	Limiter    *Limiter           `koanf:"limiter"`
	Logger     *logger.Config     `koanf:"logger"`
	Redis      *redis.Config      `koanf:"redis"`
	Nats       *nats.Config       `koanf:"nats"`
	Db         *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Mailer     *mailer.Config     `koanf:"mailer"`
	Jwt        *jwt.Config        `koanf:"jwt"`
	Cookie     *cookie.Config     `koanf:"cookie"`
	Encryption *encryption.Config `koanf:"encryption"`
}
