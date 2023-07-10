package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App
	Db
	Redis
	Nats
	Limiter
	Mailer
	Cors
}

func NewConfig() *Config {
	c := &Config{}
	c.ParseFlags()
	return c
}

func (c *Config) ParseFlags() {
	var file string
	flag.StringVar(&file, "config", "", "API config file path")

	flag.IntVar(&c.App.Port, "port", 4000, "API server port")
	flag.StringVar(&c.App.Env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()
	if file != "" {
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic(fmt.Errorf("config file not found, %w", err))
			} else {
				panic(fmt.Errorf("fatal error config file: %w", err))
			}
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/sharebuy/")
	viper.AddConfigPath("$HOME/.sharebuy")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found, %w", err))
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	c.App.LogPath = viper.GetString("log.path")

	// SHAREBUY_DB_DSN='postgres://postgres:ringsport@localhost/sharebuy?sslmode=disable'
	c.Db.Dsn = viper.GetString("db.dsn")
	c.Db.MaxOpenConns = viper.GetInt("db.maxOpenConns")
	c.Db.MaxIdleConns = viper.GetInt("db.maxIdleConns")
	c.Db.MaxIdleTime = viper.GetString("db.maxIdleTime")

	// SHAREBUY_REDIS_DSN='redis://username:password@localhost:6379?sslmode=disable'
	c.Redis.Addr = viper.GetString("redis.dsn")
	c.Redis.Db = viper.GetInt("redis.db")
	c.Redis.MaxRetries = viper.GetInt("redis.maxRetries")
	c.Redis.MinRetryBackoff = viper.GetDuration("redis.minRetryBackoff")
	c.Redis.MaxRetryBackoff = viper.GetDuration("redis.maxRetryBackoff")
	c.Redis.DialTimeout = viper.GetDuration("redis.dialTimeout")
	c.Redis.ReadTimeout = viper.GetDuration("redis.readTimeout")
	c.Redis.WriteTimeout = viper.GetDuration("redis.writeTimeout")
	c.Redis.ContextTimeoutEnabled = viper.GetBool("redis.contextTimeoutEnabled")
	c.Redis.PoolFIFO = viper.GetBool("redis.poolFIFO")
	c.Redis.PoolSize = viper.GetInt("redis.poolSize")
	c.Redis.PoolTimeout = viper.GetDuration("redis.poolTimeout")
	c.Redis.MinIdleConns = viper.GetInt("redis.minIdleConns")
	c.Redis.MaxIdleConns = viper.GetInt("redis.maxIdleConns")
	c.Redis.ConnMaxIdleTime = viper.GetDuration("redis.connMaxIdleTime")
	c.Redis.ConnMaxLifetime = viper.GetDuration("redis.connMaxLifetime")

	c.Nats.Url = viper.GetString("nats.url")
	c.Nats.Name = viper.GetString("nats.name")
	c.Nats.User = viper.GetString("nats.user")
	c.Nats.Password = viper.GetString("nats.password")
	c.Nats.Token = viper.GetString("nats.token")
	c.Nats.UseOldRequestStyle = viper.GetBool("nats.useOldRequestStyle")
	c.Nats.NoCallbacksAfterClientClose = viper.GetBool("nats.noCallbacksAfterClientClose")
	c.Nats.RetryOnFailedConnect = viper.GetBool("nats.retryOnFailedConnect")
	c.Nats.Compression = viper.GetBool("nats.compression")
	c.Nats.ProxyPath = viper.GetString("nats.proxyPath")
	c.Nats.InboxPrefix = viper.GetString("nats.inboxPrefix")
	c.Nats.IgnoreAuthErrorAbort = viper.GetBool("nats.ignoreAuthErrorAbort")
	c.Nats.SkipHostLookup = viper.GetBool("nats.skipHostLookup")

}
