package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
)

var (
	ErrInvalidLength = errors.New("invalid length")
)

type Encryption struct {
	key []byte `koanf:"key"`
}

func New(cfg *Config) Encryption {
	return Encryption{
		key: cfg.Key,
	}
}

func (enc *Encryption) Encrypt(plain []byte) ([]byte, error) {
	c, err := aes.NewCipher(enc.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func (enc *Encryption) Decrypt(ciphered []byte) ([]byte, error) {
	c, err := aes.NewCipher(enc.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphered) < nonceSize {
		return nil, errors.New("ciphered too short")
	}

	nonce, ciphered := ciphered[:nonceSize], ciphered[nonceSize:]
	return gcm.Open(nil, nonce, ciphered, nil)
}

func (enc *Encryption) Hash(s string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func (enc *Encryption) HashMatch(s string, hashed []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(s))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func GenerateRandomString(l int) (string, error) {
	if l <= 0 || l > 64 {
		return "", ErrInvalidLength
	}

	randomBytes := make([]byte, 64)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:l], nil
}
