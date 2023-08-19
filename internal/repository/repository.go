package repository

import (
	"errors"
	"go.uber.org/zap"

	"github.com/farmani/sharebuy/pkg/rdbms"
)

var (
	ErrRecordNotFound   = errors.New("record not found")
	ErrDuplicateEmail   = errors.New("duplicate email")
	ErrEditConflict     = errors.New("edit conflict")
	ErrSavingDataFailed = errors.New("cannot save data")
)

type Repository interface {
	User() UserRepository
	Product() ProductRepository
	Lottery() LotteryRepository
	Token() TokenRepository
}

type repository struct {
	logger *zap.Logger
	config *Config
	db     rdbms.DB
}

func New(lg *zap.Logger, cfg *Config, rdbms rdbms.DB) Repository {
	return &repository{
		logger: lg,
		config: cfg,
		db:     rdbms,
	}
}

func (r *repository) User() UserRepository {
	return NewUserRepository(r.logger, r.config, r.db)
}
func (r *repository) Product() ProductRepository {
	return NewProductRepository(r.logger, r.config, r.db)
}
func (r *repository) Lottery() LotteryRepository {
	return NewLotteryRepository(r.logger, r.config, r.db)
}
func (r *repository) Token() TokenRepository {
	return NewTokenRepository(r.logger, r.config, r.db)
}
