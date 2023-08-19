package models

import (
	"time"

	"github.com/farmani/sharebuy/internal/models/enums"
)

type Product struct {
	ID        int64               `json:"id"` // Unique integer ID for the product
	UUID      string              `json:"uuid"`
	Title     string              `json:"title"`
	Price     int                 `json:"price,omitempty"`
	Currency  string              `json:"currency,omitempty"`
	URL       string              `json:"url,omitempty"`
	Status    enums.ProductStatus `json:"status,omitempty"`
	Images    []Image             `json:"images,omitempty"`
	CreatedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	UpdatedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	DeletedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	Version   int                 `json:"version"` // The version number starts at 1 and is incremented each
	// time the movie information is updated.
}
