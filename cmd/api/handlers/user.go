package handlers

import (
	"errors"
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// Dependencies or state for the UserHandler
	app *app.Application
	userService *services.UserService
}

func NewUserHandler(a *app.Application, u *services.UserService) *UserHandler {
	return &UserHandler{
		app: a,
		userService: u,
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
	// Create an anonymous struct to hold the expected data from the request body.
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := s.app.ReadJSON(c, &input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}


	v := validator.New()

	// Validate the user struct and return the error messages to the client if
	// any of the checks fail.
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	us := h.app.GetService("UserService").(*services.UserService)
	user, err := us.RegisterStart(c)
	if err != nil {
		return failedValidationResponse(c, errors.New("failed to register user"))
	}

	}
	// us.RegisterStart(c)
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"user": user,
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
