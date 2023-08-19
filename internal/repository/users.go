package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/farmani/sharebuy/pkg/rdbms"

	"github.com/farmani/sharebuy/internal/models"
)

type UserRepository interface {
	GetById(id int64) (*models.User, error)
	GetByUuid(uuid string) (*models.User, error)
	Insert(user *models.User) error
	Update(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByEmailAndPassword(email, password string) (*models.User, error)
}

type Users struct {
	db     rdbms.DB
	config *Config
	logger *zap.Logger
}

func NewUserRepository(lg *zap.Logger, cfg *Config, db rdbms.DB) *Users {
	return &Users{
		db:     db,
		config: cfg,
		logger: lg,
	}
}

func (u *Users) GetById(id int64) (*models.User, error) {
	query := `
		SELECT id, created_at, name, email, password, status, version
		FROM users
		WHERE id = $1
		`

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := []any{id}
	out := []any{
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Status,
		&user.Version,
	}

	err := u.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *Users) GetByUuid(uuid string) (*models.User, error) {
	query := `
		SELECT id, uuid, name, email, password, status, version, created_at
		FROM users
		WHERE uuid = $1
		`

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := []any{uuid}
	out := []any{
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Status,
		&user.Version,
		&user.CreatedAt,
	}

	err := u.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *Users) Insert(user *models.User) error {
	query := `
		INSERT INTO users (name, email, username, password, strip_id, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, version;
		`

	in := []any{user.Name, user.Email, user.Username, user.Password.Hash, user.StripeId, user.Status.String()}
	out := []any{&user.ID, &user.CreatedAt, &user.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, in, out)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), `pq: duplicate key value violates unique constraint "users_email_key"`):
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (u *Users) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, created_at, name, email, password, status, version
		FROM users
		WHERE email = $1
		`

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := []any{email}
	out := []any{&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Status,
		&user.Version,
	}

	err := u.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *Users) GetByEmailAndPassword(email, password string) (*models.User, error) {
	return nil, nil
}

func (u *Users) Update(user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password = $3, status = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version
		`

	in := []any{
		user.Name,
		user.Email,
		user.Password.Hash,
		user.Status,
		user.ID,
		user.Version,
	}
	out := []any{&user.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, in, out)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (u *Users) GetByToken(tokenScope, tokenPlaintext string) (*models.User, error) {
	return nil, nil
}
