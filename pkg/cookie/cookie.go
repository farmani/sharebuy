package cookie

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/labstack/echo/v4"
)

type Cookie struct {
	Domain   string
	Path     string
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
	Expires  time.Duration
	MaxAge   int
}

func New(cfg *Config) *Cookie {
	cookie := &Cookie{}
	cookie.Domain = cfg.Domain
	cookie.Path = cfg.Path
	cookie.Secure = cfg.Secure
	cookie.HttpOnly = cfg.HttpOnly
	cookie.SameSite = cfg.SameSite
	cookie.Expires = cfg.Expires
	cookie.MaxAge = cfg.MaxAge
	return cookie
}

func (co *Cookie) SetCookies(c echo.Context, n string, v string) {
	cookie := new(http.Cookie)
	cookie.Path = co.Path
	cookie.Secure = co.Secure
	cookie.HttpOnly = co.HttpOnly
	cookie.SameSite = co.SameSite
	cookie.Expires = time.Now().Add(co.Expires)
	cookie.MaxAge = co.MaxAge
	cookie.Domain = co.Domain
	cookie.Name = n
	cookie.Value = v
	c.SetCookie(cookie)
}

func (co *Cookie) SetEncryptedCookies(c echo.Context, e encryption.Encryption, n string, v string) error {
	encV, err := e.Encrypt([]byte(v))
	if err != nil {
		return err
	}

	base64 := base64.StdEncoding.EncodeToString(encV)

	co.SetCookies(c, n, base64)

	return nil
}

func (co *Cookie) ReadEncryptedCookies(c echo.Context, e encryption.Encryption, n string) (string, error) {
	cookie, err := c.Cookie(n)
	if err != nil {
		return "", err
	}

	base64, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", err
	}

	v, err := e.Decrypt([]byte(base64))
	if err != nil {
		return "", err
	}

	return string(v), nil
}
