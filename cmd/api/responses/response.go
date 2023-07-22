package responses

import (
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/labstack/echo/v4"
)

// readString is a helper method on Application type that returns a string value from the URL query
// string, or the provided default value if no matching key is found.
func WriteJSON(c echo.Context, data interface{}) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data:   data,
	}

	return c.JSON(http.StatusOK, res)
}
