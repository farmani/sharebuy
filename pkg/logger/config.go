package logger

type Config struct {
	Level    string `koanf:"level"`
	Path     string `koanf:"path"`
	Env      string `koanf:"env"`
	Encoding string `koanf:"encoding"`
}
