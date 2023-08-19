package requests

import (
	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

type RegisterStartRequest struct {
	Email      string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password   string `json:"password" xml:"password" form:"password" validate:"required"`
	SuccessUrl string `json:"success_url" xml:"success_url" form:"success_url" validate:"required"`
}

func NewRegisterStartRequest(c echo.Context, v *validator.Validator) (*RegisterStartRequest, error) {
	r := &RegisterStartRequest{}

	err := c.Bind(&r)
	if err != nil {
		return nil, ErrBadRequest
	}

	if v.Validate(r); !v.Valid() {
		return nil, ErrFailedValidation
	}

	return r, nil
}

func (r *RegisterStartRequest) Validate(v *validator.Validator) {
	v.Validate(r)
}
