package models

import (
	"time"

	"github.com/farmani/sharebuy/internal/models/enums"
)

type Lottery struct {
	ID                    int64               `json:"id"`
	UUID                  string              `json:"uuid"`
	ProductID             int64               `json:"product_id"`
	TicketPrice           int64               `json:"ticket_price"`
	ProductPrice          int64               `json:"product_price"`
	TicketCounts          int64               `json:"ticket_counts"`
	AvailableTicketCounts int64               `json:"available_ticket_counts"`
	Status                enums.LotteryStatus `json:"status"`
	StartedAt             time.Time           `json:"-"`
	EndedAt               time.Time           `json:"-"`
	CreatedAt             time.Time           `json:"-"`
	UpdatedAt             time.Time           `json:"-"`
	DeletedAt             time.Time           `json:"-"`
	Version               int32               `json:"version"` // The version number starts at 1 and is incremented each
	// time the movie information is updated.
}
