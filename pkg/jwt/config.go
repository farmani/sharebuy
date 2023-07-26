package jwt

import "time"

type Config struct {
	PrivatePem        string        `koanf:"private_pem"`
	PublicPem         string        `koanf:"public_pem"`
	Expiration        time.Duration `koanf:"expiration"`
	RefreshExpiration time.Duration `koanf:"refresh_expiration"`
	CookieTokenName   string        `koanf:"cookie_token_name"`
}
