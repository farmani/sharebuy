package rdbms

type Config struct {
	Dsn          string `koanf:"rdbms.dsn"`
	Host         string `koanf:"rdbms.host"`
	Port         string `koanf:"rdbms.port"`
	Username     string `koanf:"rdbms.username"`
	Password     string `koanf:"rdbms.password"`
	Database     string `koanf:"rdbms.database"`
	MaxOpenConns int    `koanf:"rdbms.max_open_conns"`
	MaxIdleConns int    `koanf:"rdbms.max_idle_conns"`
	MaxIdleTime  string `koanf:"rdbms.max_idle_time"`
}
