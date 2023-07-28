package repository

import (
	"github.com/farmani/sharebuy/pkg/rdbms"
	"go.uber.org/zap"
)

type Repository interface {
}

type repository struct {
	logger *zap.Logger
	config *Config
	db     rdbms.DB
}

func New(logger *zap.Logger, cfg *Config, db rdbms.DB) Repository {
	return &repository{
		logger: logger,
		config: cfg,
		db:     db,
	}
}
