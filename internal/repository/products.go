package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"

	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/pkg/rdbms"
	"go.uber.org/zap"
)

type ProductRepository interface {
	Insert(product *models.Product) error
	GetById(id int64) (models.Product, error)
	GetByUuid(uuid string) (models.Product, error)
	Update(product *models.Product) error
}

type Products struct {
	db     rdbms.DB
	config *Config
	logger *zap.Logger
}

func NewProductRepository(lg *zap.Logger, cfg *Config, db rdbms.DB) Products {
	return Products{
		db:     db,
		config: cfg,
		logger: lg,
	}
}

func (p Products) GetById(id int64) (models.Product, error) {
	query := `
		SELECT id, created_at, title, status, version
		FROM products
		WHERE id = $1
		`

	var product models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := []any{id}
	out := []any{
		&product.ID,
		&product.CreatedAt,
		&product.Status,
		&product.Version,
	}

	err := p.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return product, ErrRecordNotFound
		default:
			return product, err
		}
	}

	return product, nil
}

func (p Products) GetByUuid(uuid string) (models.Product, error) {
	query := `
		SELECT id, uuid, title, price, currency, url, images, status, created_at, updated_at, version
		FROM products
		WHERE uuid = $1 AND status = 'active'
		`

	var product models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	in := []any{uuid}
	out := []any{
		&product.ID,
		&product.UUID,
		&product.Title,
		&product.Price,
		&product.Currency,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Version,
	}

	err := p.db.QueryRowContext(ctx, query, in, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return product, ErrRecordNotFound
		default:
			return product, err
		}
	}

	return product, nil
}

func (p Products) Insert(product *models.Product) error {
	query := `
		INSERT INTO products (title, price, currency, url, images, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, uuid, created_at, updated_at, version
		`

	in := []any{product.Title, product.Price, product.Currency, pq.Array(product.Images), product.Status}
	out := []any{&product.ID, &product.UUID, &product.CreatedAt, &product.UpdatedAt, &product.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.db.QueryRowContext(ctx, query, in, out)
}

func (p Products) Update(product *models.Product) error {
	query := `
		UPDATE products
		SET title = $1, price = $2, currency = $3, url = $4, images = $5, status = $6, version = version + 1
		WHERE id = $7 AND version = $8
		RETURNING version
		`

	in := []any{
		product.Title,
		product.Price,
		product.Currency,
		product.URL,
		pq.Array(product.Images),
		product.Status,
		product.ID,
		product.Version,
	}
	out := []any{&product.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.db.QueryRowContext(ctx, query, in, out)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "products_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
