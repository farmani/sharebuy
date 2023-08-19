package models

import (
	"time"

	"github.com/farmani/sharebuy/internal/models/enums"
)

type Image struct {
	ID        int64             `json:"id"` // Unique integer ID for the product
	UUID      string            `json:"uuid"`
	ProductID int64             `json:"product_id"`
	File      string            `json:"file"`
	Mime      string            `json:"mime"`
	Url       string            `json:"url"`
	Status    enums.ImageStatus `json:"status"`
	CreatedAt time.Time         `json:"-"`       // Use the - directive to never export in JSON output
	UpdatedAt time.Time         `json:"-"`       // Use the - directive to never export in JSON output
	DeletedAt time.Time         `json:"-"`       // Use the - directive to never export in JSON output
	Version   int               `json:"version"` // The version number starts at 1 and is incremented each
}
