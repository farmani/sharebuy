package requests

import (
	"github.com/farmani/sharebuy/internal/models/enums"
	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

type ProductUpdateRequest struct {
	UUID     string               `json:"uuid" xml:"uuid" form:"uuid" validate:"required,uuid4"`
	Title    *string              `json:"title" xml:"title" form:"title" validate:"alphanumunicode,min=3,max=255"`
	Price    *int                 `json:"price" xml:"price" form:"price" validate:""`
	Currency *string              `json:"currency" xml:"currency" form:"currency" validate:"iso4217"`
	URL      *string              `json:"url" xml:"url" form:"url" validate:"url"`
	Status   *enums.ProductStatus `json:"status" xml:"status" form:"status" validate:"required"`
	Images   []int64              `json:"images" xml:"images" form:"images" validate:"required"`
}

func NewProductUpdateRequest(c echo.Context, v *validator.Validator) (ProductUpdateRequest, error) {
	r := ProductUpdateRequest{}

	err := c.Bind(&r)
	if err != nil {
		return r, ErrBadRequest
	}

	r.UUID = c.Param("uuid")

	if v.Validate(r); !v.Valid() {
		return r, ErrFailedValidation
	}

	return r, nil
}

func (r *ProductUpdateRequest) Validate(v *validator.Validator) {
	v.Validate(r)
}
