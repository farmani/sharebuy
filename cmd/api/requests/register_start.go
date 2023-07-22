package requests

import (
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/internal/validator"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type RegisterStartRequest struct {
	Email    string `param:"email" json:"email" xml:"email" form:"email"`
	Password string `param:"password" json:"password" xml:"password" form:"password"`
}

func NewRegisterStartRequest(c echo.Context) *RegisterStartRequest {
	return &RegisterStartRequest{}
}

func (r *RegisterStartRequest) Validate(c echo.Context) (*validator.Validator, error) {
	user := &data.User{
		Email:     r.Email,
		Activated: false,
	}

	err := user.Password.Set(r.Password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to encrypt password")
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		return v, errors.Wrap(err, "invalid request body")
	}

	return v, nil
}
