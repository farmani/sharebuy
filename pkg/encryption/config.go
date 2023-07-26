package encryption

type Config struct {
	Key       []byte `koanf:"key"`
	Algorithm string `koanf:"algorithm"`
}
