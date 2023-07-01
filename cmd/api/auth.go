package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) loginHandler(c echo.Context) error {
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

func (app *application) logoutHandler(c echo.Context) error {
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

func (app *application) forgetPasswordStartHandler(c echo.Context) error {
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

func (app *application) forgetPasswordFinishHandler(c echo.Context) error {
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

func (app *application) changePasswordHandler(c echo.Context) error {
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

func (app *application) registerStartHandler(c echo.Context) error {
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

func (app *application) registerFinishHandler(c echo.Context) error {
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
