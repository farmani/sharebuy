package models

import (
	"github.com/farmani/sharebuy/internal/models/enums"
	"time"
)

type Token struct {
	Plaintext string           `json:"token"`
	Hash      []byte           `json:"-"`
	UserID    int64            `json:"-"`
	Scope     enums.TokenScope `json:"-"`
	ExpiredAt time.Time        `json:"expired_at"`
}
