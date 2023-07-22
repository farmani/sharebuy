package db

type Config struct {
	Dsn          string `koanf:"db.dsn"`
	MaxOpenConns int    `koanf:"db.max_open_conns"`
	MaxIdleConns int    `koanf:"db.max_idle_conns"`
	MaxIdleTime  string `koanf:"db.max_idle_time"`
}
