package cookie

import (
	"net/http"
	"time"
)

type Config struct {
	Domain   string        `koanf:"domain"`
	Path     string        `koanf:"path"`
	Secure   bool          `koanf:"secure"`
	HttpOnly bool          `koanf:"http_only"`
	SameSite http.SameSite `koanf:"same_site"`
	Expires  time.Duration `koanf:"expires"`
	MaxAge   int           `koanf:"max_age"`
}
