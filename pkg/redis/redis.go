package redis

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func New(cfg *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})

	return rdb
}
