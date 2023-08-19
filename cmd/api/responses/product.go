package responses

import (
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/models/enums"
	"time"
)

type ProductCollection struct {
	Products []ProductResponse `json:"products"`
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Page     int               `json:"page"`
	Pages    int               `json:"pages"`
}

type ProductResponse struct {
	ID        int64               `json:"id"` // Unique integer ID for the product
	UUID      string              `json:"uuid"`
	Title     string              `json:"title"`
	Price     int                 `json:"price,omitempty"`
	Currency  string              `json:"currency,omitempty"`
	URL       string              `json:"url,omitempty"`
	Status    enums.ProductStatus `json:"status,omitempty"`
	Images    []models.Image      `json:"images,omitempty"`
	CreatedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	UpdatedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	DeletedAt time.Time           `json:"-"`       // Use the - directive to never export in JSON output
	Version   int                 `json:"version"` // The version number starts at 1 and is incremented each
}

func NewProductResponse(p *models.Product) *ProductResponse {
	return &ProductResponse{
		UUID:     p.UUID,
		Title:    p.Title,
		Price:    p.Price,
		Currency: p.Currency,
		URL:      p.URL,
		Status:   p.Status,
		Images:   p.Images,
		Version:  p.Version,
	}
}
