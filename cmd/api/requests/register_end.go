package requests

import (
	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

type RegisterEndRequest struct {
	FirstName string `json:"email" form:"email" query:"first_name"`
	LastName  string `json:"password" form:"password" query:"last_name"`
	UserId    int64  `json:"user_id" form:"user_id" query:"user_id"`
	Token     string `json:"token" form:"token" query:"token"`
}

func NewRegisterEndRequest(c echo.Context, v *validator.Validator) (*RegisterEndRequest, error) {
	r := &RegisterEndRequest{}

	err := c.Bind(&r)
	if err != nil {
		return nil, ErrBadRequest
	}

	if v.Validate(r); !v.Valid() {
		return nil, ErrFailedValidation
	}

	return r, nil
}

func (r *RegisterEndRequest) Validate(c *echo.Context) {
	// e.POST("/v1/users", h.RegisterStart)
}
