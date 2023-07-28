package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func New(cfg *Config) (DB, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil for DB")
	}

	connStr := cfg.Dsn
	if connStr == "" {
		connStr = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	}

	pg, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("open DB Failed: %w", err)
	}

	pg.SetMaxOpenConns(cfg.MaxOpenConns)
	pg.SetMaxIdleConns(cfg.MaxIdleConns)
	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("parse duration Failed: %w", err)
	}

	pg.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping DB Failed: %w", err)
	}

	// boil.SetDB(rdbms)

	return &db{db: pg}, nil
}
