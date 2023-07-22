package requests

import "github.com/labstack/echo/v4"

type RegisterEndRequest struct {
	FirstName string `json:"email" form:"email" query:"first_name"`
	LastName  string `json:"password" form:"password" query:"last_name"`
}

func (r *RegisterEndRequest) Validate(c *echo.Context) {
	// e.POST("/v1/users", h.RegisterStart)
}
