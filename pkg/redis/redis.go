package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(cfg *Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb.Ping(ctx)

	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		return nil, err
	}

	_, err = rdb.Get(ctx, "foo").Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
