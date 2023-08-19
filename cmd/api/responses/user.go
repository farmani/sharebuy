package responses

import (
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/models/enums"
	"time"
)

type UserViewResponse struct {
	ID           int64            `json:"id"`
	Email        string           `json:"email"`
	Name         string           `json:"name"`
	CreatedAt    time.Time        `json:"created_at,omitempty"`
	UpdatedAt    time.Time        `json:"updated_at,omitempty"`
	RegisteredAt time.Time        `json:"registered_at,omitempty"`
	Status       enums.UserStatus `json:"status,omitempty"`
	Version      int              `json:"-"`
}

type MeResponse struct {
	ID           int64            `json:"id"`
	Email        string           `json:"email"`
	Name         string           `json:"name"`
	CreatedAt    time.Time        `json:"created_at,omitempty"`
	UpdatedAt    time.Time        `json:"updated_at,omitempty"`
	RegisteredAt time.Time        `json:"registered_at,omitempty"`
	Status       enums.UserStatus `json:"status,omitempty"`
	Version      int              `json:"-"`
	Token        string           `json:"token,omitempty"`
}

func NewUserViewResponse(u *models.User) *UserViewResponse {
	return &UserViewResponse{
		ID:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		RegisteredAt: u.CreatedAt,
		Status:       u.Status,
		Version:      u.Version,
	}
}

func NewMeResponse(u *models.User, t string) *MeResponse {
	return &MeResponse{
		ID:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		RegisteredAt: u.CreatedAt,
		Status:       u.Status,
		Token:        t,
		Version:      u.Version,
	}
}
