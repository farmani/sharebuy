package handlers

import (
	"net/http"
	"strconv"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/responses"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/data"
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
	//e.POST("/v1/users", h.RegisterStart)
	e.PUT("/v1/users/activated", h.RegisterEnd)
	// e.POST("/v1/users/login", h.createAuthenticationTokenHandler)
	// e.POST("/v1/users/logout", h.createAuthenticationTokenHandler)
	// e.POST("/v1/users/forget", h.createAuthenticationTokenHandler)
	e.GET("/v1/users/:id", h.UserView)
}

func (h *UserHandler) UserView(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.FindById(int64(id))
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			notFoundResponse(c)
		default:
			serverErrorResponse(c, err)
		}
		return nil
	}

	return responses.WriteJSON(c, responses.NewUserViewResponse(&user))
}

/*
	func (h *UserHandler) RegisterStart(c echo.Context) error {
		r := requests.NewRegisterStartRequest(c)
		err := c.Bind(&r)
		if err != nil {
			badRequestResponse(c, err)
			return nil
		}

		if v, err := r.Validate(c); err != nil {
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
*/
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
