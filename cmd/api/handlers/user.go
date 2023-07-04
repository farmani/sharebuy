package handlers

import (
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewUserHandler(a *app.Application) *UserHandler {
	return &UserHandler{
		app: a,
	}
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	// Define routes specific to the UserHandler

	// Users handlers
	e.POST("/v1/users", h.RegisterStart)
	e.PUT("/v1/users/activated", h.RegisterEnd)
	// e.POST("/v1/users/login", h.createAuthenticationTokenHandler)
	// e.POST("/v1/users/logout", h.createAuthenticationTokenHandler)
	// e.POST("/v1/users/forget", h.createAuthenticationTokenHandler)
}

func (h *UserHandler) RegisterStart(c echo.Context) error {
	// us := h.app.GetService("UserService").(*services.UserService)
	// err := us.RegisterStart(c)
	// us.RegisterStart(c)
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) RegisterEnd(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Me(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Login(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Logout(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) ForgetPasswordStart(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) ForgetPasswordEnd(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return c.JSON(http.StatusOK, res)
}
