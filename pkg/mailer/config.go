package mailer

type Config struct {
	Host         string `koanf:"host"`
	Port         int    `koanf:"port"`
	Username     string `koanf:"username"`
	Password     string `koanf:"password"`
	Sender       string `koanf:"sender"`
	ResendAPIKey string `koanf:"resend_api_key"`
}
