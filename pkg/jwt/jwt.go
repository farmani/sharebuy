package jwt

import (
	"context"
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

const (
	invalidToken                 = "invalid token"
	errorMappingPayload          = "unable to map payload"
	errorUnmarshalData           = "unable to unmarshal data"
	redisUserTokenKeyPrefix      = "sharebuy:user:token:%s"
	redisUserTokenBlockKeyPrefix = "sharebuy:user:block:token:%s"
)

type Token interface {
	GenerateToken(data any) (string, error)
	ExtractData(token string, data any) error
	GenerateRefreshToken(data []byte) (string, error)
	BlockToken(token string) error
}

type token struct {
	privateEd25519Key  crypto.PrivateKey
	publicEd25519Key   crypto.PublicKey
	expiration         time.Duration
	refresh_expiration time.Duration
	redis              *redis.Client
}

type Claims struct {
	Data []byte `json:"data" xml:"data" form:"data"`
	jwt.RegisteredClaims
}

func New(cfg *Config, r *redis.Client) (Token, error) {
	token := &token{}

	token.redis = r

	var err error

	privatePemKey := []byte(cfg.PrivatePem)
	token.privateEd25519Key, err = jwt.ParseEdPrivateKeyFromPEM(privatePemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 private key: %w", err)
	}

	publicPemKey := []byte(cfg.PublicPem)
	token.publicEd25519Key, err = jwt.ParseEdPublicKeyFromPEM(publicPemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 public key: %w", err)
	}

	token.expiration = cfg.Expiration
	token.refresh_expiration = cfg.RefreshExpiration

	return token, nil
}

func GenerateTokenAndSetCookies(data any, c echo.Context) error {

	return nil
}

func (t *token) GenerateToken(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}

	expiredAt := jwt.NewNumericDate(time.Now().Add(t.expiration))
	registeredClaim := jwt.RegisteredClaims{
		Issuer:    "sharebuy",
		ExpiresAt: expiredAt,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Audience:  jwt.ClaimStrings{"sharebuy"},
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	claims := &Claims{dataBytes, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	token, err := jwtToken.SignedString(t.privateEd25519Key)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hash := md5.Sum([]byte(token))
	key := fmt.Sprintf(redisUserTokenKeyPrefix, hex.EncodeToString(hash[:]))
	err = t.redis.Set(ctx, key, 1, 0).Err()
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}

	return token, nil
}

func (t *token) GenerateRefreshToken(data []byte) (string, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Add(t.refresh_expiration))
	registeredClaim := jwt.RegisteredClaims{
		Issuer:    "sharebuy",
		ExpiresAt: expiredAt,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Audience:  jwt.ClaimStrings{"sharebuy"},
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	claims := &Claims{data, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	return jwtToken.SignedString(t.privateEd25519Key)
}

func (t *token) BlockToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error

	hash := md5.Sum([]byte(token))
	md5Token := hex.EncodeToString(hash[:])

	key := fmt.Sprintf(redisUserTokenKeyPrefix, md5Token)
	err = t.redis.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	blockKey := fmt.Sprintf(redisUserTokenBlockKeyPrefix, md5Token)
	err = t.redis.Set(ctx, blockKey, 1, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (t *token) ExtractData(token string, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hash := md5.Sum([]byte(token))
	key := fmt.Sprintf(redisUserTokenKeyPrefix, hex.EncodeToString(hash[:]))
	_, err := t.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("token is expired or not exists")
	}
	if err != nil {
		return fmt.Errorf("cannot read from storage")
	}

	checkSigningMethod := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, checkSigningMethod)
	if err != nil {
		return fmt.Errorf("unable to parse token: %w", err)
	}

	if !jwtToken.Valid {
		return fmt.Errorf("%s: token: %v", invalidToken, jwtToken)
	}

	payload, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return fmt.Errorf("%s: %s, token: %v", invalidToken, errorMappingPayload, jwtToken)
	}

	if err := json.Unmarshal([]byte(payload.Data), data); err != nil {
		return fmt.Errorf("%s: %s, data: %s", invalidToken, errorUnmarshalData, payload.Data)
	}

	return nil
}
