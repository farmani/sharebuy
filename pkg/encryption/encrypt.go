package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
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
