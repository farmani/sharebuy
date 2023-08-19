package repository

import (
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/pkg/rdbms"
	"go.uber.org/zap"
)

type LotteryRepository interface {
	Insert(lottery *models.Lottery) error
	GetById(id int64) (*models.Lottery, error)
	Update(lottery *models.Lottery) error
}

type Lotteries struct {
	db     rdbms.DB
	config *Config
	logger *zap.Logger
}

func NewLotteryRepository(lg *zap.Logger, cfg *Config, db rdbms.DB) Lotteries {
	return Lotteries{
		db:     db,
		config: cfg,
		logger: lg,
	}
}

func (l Lotteries) GetById(id int64) (*models.Lottery, error) {
	return nil, nil
}

func (l Lotteries) GetByUuid(uuid string) (*models.Lottery, error) {
	return nil, nil
}

func (l Lotteries) Insert(lottery *models.Lottery) error {
	return nil
}

func (l Lotteries) Update(lottery *models.Lottery) error {

	return nil
}
