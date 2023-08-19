package repository

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"errors"
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/models/enums"
	"github.com/farmani/sharebuy/pkg/rdbms"
	"go.uber.org/zap"
	"time"
)

type TokenRepository interface {
	GenerateToken(userID int64, ttl time.Duration, scope enums.TokenScope) (*models.Token, error)
	Insert(token *models.Token) error
	GetByHash(hash string, userId int64, scope enums.TokenScope) (*models.Token, error)
	DeleteAllForUser(scope enums.TokenScope, userID int64) error
}

type Tokens struct {
	db     rdbms.DB
	config *Config
	logger *zap.Logger
}

func NewTokenRepository(lg *zap.Logger, cfg *Config, db rdbms.DB) Tokens {
	return Tokens{
		db:     db,
		config: cfg,
		logger: lg,
	}
}

func (t Tokens) GetByHash(hash string, userId int64, scope enums.TokenScope) (*models.Token, error) {
	query := `
		SELECT hash, user_id, expired_at, scope
		FROM tokens
		WHERE user_id = $1 AND scope = $2 AND hash = $3 AND expired_at > NOW()
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var token models.Token

	in := []any{hash, userId, scope.String()}
	out := []any{
		&token.Hash,
		&token.UserID,
		&token.ExpiredAt,
		&token.Scope,
	}

	err := t.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &token, nil
}

func (t Tokens) Insert(token *models.Token) error {
	query := `
		INSERT INTO tokens (hash, user_id, expired_at, scope)
		VALUES ($1, $2, $3, $4);
		`

	args := []any{token.Hash, token.UserID, token.ExpiredAt, token.Scope.String()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.db.ExecContext(ctx, query, args)
}

func (t Tokens) DeleteAllForUser(scope enums.TokenScope, userID int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{scope, userID}

	err := t.db.ExecContext(ctx, query, args)
	return err
}

func (t Tokens) GenerateToken(userID int64, ttl time.Duration, scope enums.TokenScope) (*models.Token, error) {
	token := &models.Token{
		UserID:    userID,
		ExpiredAt: time.Now().Add(ttl),
		Scope:     scope,
	}

	// Initialize a zero-valued byte slice with a length of 16 bytes.
	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Encode the byte slice to a base-32 encoded string and assign it to the token Plaintext field.
	// This will be the token string that we send to the user in their welcome email. They will
	// look similar to this:
	//
	// Y3QMGX3PJ3WLRL2YRTQGQ6KRHU
	//
	// Note that by default base-32 strings may be padded at the end with the = character.
	// However, we don't need this padding character for the purpose of our tokens, so we use
	// the WithPadding(base32.NoPadding) method in the line below to omit them.
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Generate a SHA-256 hash of the plaintext token string. This will be the value
	// that we store in the `hash` field of our tokens table.
	// Note, that the sha256.Sum256() function returns an *array* of length 32,
	// so to make it easier to work with we convert it to a slice using the [:]
	// operator before operating on it.
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}
