package models

import "time"

type LotteryParticipant struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	LotteryID    int64     `json:"lottery_id"`
	ProductID    int64     `json:"product_id"`
	Amount       int64     `json:"amount"`
	TicketCounts int64     `json:"ticket_counts"`
	CreatedAt    time.Time `json:"-"`
	// time the movie information is updated.
}
