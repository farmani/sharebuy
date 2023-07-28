package redis

type Config struct {
	Host                  string `koanf:"redis.host"`
	Port                  int    `koanf:"redis.port"`
	Password              string `koanf:"redis.password"`
	Db                    int    `koanf:"redis.rdbms"`
	MaxRetries            int    `koanf:"redis.max_retries"`
	MinRetryBackoff       string `koanf:"redis.min_retry_backoff"`
	MaxRetryBackoff       string `koanf:"redis.max_retry_backoff"`
	DialTimeout           string `koanf:"redis.dial_timeout"`
	ReadTimeout           string `koanf:"redis.read_timeout"`
	WriteTimeout          string `koanf:"redis.write_timeout"`
	ContextTimeoutEnabled bool   `koanf:"redis.context_timeout_enabled"`
	PoolFIFO              bool   `koanf:"redis.pool_fifo"`
	PoolSize              int    `koanf:"redis.pool_size"`
	PoolTimeout           string `koanf:"redis.pool_timeout"`
	MinIdleConns          int    `koanf:"redis.min_idle_conns"`
	MaxIdleConns          int    `koanf:"redis.max_idle_conns"`
	ConnMaxIdleTime       string `koanf:"redis.conn_max_idle_time"`
	ConnMaxLifetime       string `koanf:"redis.conn_max_lifetime"`
}
