package requests

import (
	"github.com/farmani/sharebuy/internal/models/enums"
	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

type ProductCreateRequest struct {
	Title    string              `json:"title" xml:"title" form:"title" validate:"alphanumunicode,min=3,max=255"`
	Price    int                 `json:"price" xml:"price" form:"price" validate:"gt=0,lt=1000000"`
	Currency string              `json:"currency" xml:"currency" form:"currency" validate:"iso4217"`
	URL      string              `json:"url" xml:"url" form:"url" validate:"url"`
	Status   enums.ProductStatus `json:"status" xml:"status" form:"status" validate:"required"`
	Images   []int64             `json:"images" xml:"images" form:"images" validate:"required"`
}

func NewProductCreateRequest(c echo.Context, v *validator.Validator) (*ProductCreateRequest, error) {
	r := &ProductCreateRequest{}

	err := c.Bind(&r)
	if err != nil {
		return nil, ErrBadRequest
	}

	if v.Validate(r); !v.Valid() {
		return nil, ErrFailedValidation
	}

	return r, nil
}

func (r *ProductCreateRequest) Validate(v *validator.Validator) {
	v.Validate(r)
}
