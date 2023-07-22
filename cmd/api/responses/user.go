package responses

import (
	"time"

	"github.com/farmani/sharebuy/internal/data"
)

type UserViewResponse struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	RegisteredAt time.Time `json:"registered_at,omitempty"`
	Activated    bool      `json:"activated,omitempty"`
	Version      int       `json:"-"`
}

func NewUserViewResponse(u *data.User) *UserViewResponse {
	return &UserViewResponse{
		ID:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		RegisteredAt: u.CreatedAt,
		Activated:    u.Activated,
		Version:      u.Version,
	}
}
