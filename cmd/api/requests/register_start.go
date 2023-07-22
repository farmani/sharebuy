package requests

import (
	"fmt"

	"github.com/farmani/sharebuy/internal/data"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type RegisterStartRequest struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}

func NewRegisterStartRequest(c echo.Context) *RegisterStartRequest {
	return &RegisterStartRequest{}
}

func (r *RegisterStartRequest) Validate(v *validator.Validate, c echo.Context) error {

	err := v.Struct(r)
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return errors.Wrap(err, "unable to process request")
	}

	user := &data.User{
		Email:     r.Email,
		Activated: false,
	}

	err = user.Password.Set(r.Password)
	if err != nil {
		return errors.Wrap(err, "unable to encrypt password")
	}

	return nil
}
