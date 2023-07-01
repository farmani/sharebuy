package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) meHandler(c echo.Context) error {
	res := envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     app.config.env,
			"version": version,
		},
	}

	return c.JSON(http.StatusOK, res)
}
