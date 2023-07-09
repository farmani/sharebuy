package handlers

import (
	"errors"
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/internal/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// Dependencies or state for the UserHandler
	app         *app.Application
	userService *services.UserService
}

func NewUserHandler(a *app.Application, u *services.UserService) *UserHandler {
	return &UserHandler{
		app:         a,
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
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Parse the request body into the anonymous struct
	err := h.app.ReadJSON(c, &input)
	if err != nil {
		badRequestResponse(c, err)
		return nil
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		serverErrorResponse(c, err)
		return nil
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		failedValidationResponse(c, v.Errors)
		return nil
	}

	err = h.userService.RegisterStart(user)
	if err != nil {
		switch {
		// If we get an ErrDuplicateEmail error, use the v.AddError() method to manually add
		// a message to the validator instance, and then call our failedValidationResponse
		// helper().
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			failedValidationResponse(c, v.Errors)
		default:
			serverErrorResponse(c, err)
		}
		return nil
	}

	// us.RegisterStart(c)
	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"name":     input.Name,
			"email":    input.Email,
			"password": input.Password,
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
