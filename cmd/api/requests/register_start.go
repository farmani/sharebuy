package requests

import (
	"errors"

	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

var (
	ErrBadRequest       = errors.New("bad request error")
	ErrFailedValidation = errors.New("failed validation error")
)

type RegisterStartRequest struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email_dns"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
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
