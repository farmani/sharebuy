package handlers

import (
	"errors"
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/requests"
	"github.com/farmani/sharebuy/cmd/api/responses"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/repository"
	"github.com/farmani/sharebuy/pkg/jwt"
	"github.com/farmani/sharebuy/pkg/validator"
	golangJwt "github.com/golang-jwt/jwt/v5"
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

	/*
		privateEd25519Key, err := golangJwt.ParseEdPrivateKeyFromPEM([]byte(h.app.Config.Jwt.PrivatePem))
		if err != nil {
			h.app.Logger.Warn("unable to parse Ed25519 private key")
		}
		config := echojwt.Config{
			SigningKey:    privateEd25519Key,
			SigningMethod: "EdDSA",
			TokenLookup:   "header:Authorization:Bearer ,cookie:" + h.app.Config.Jwt.CookieTokenName,
			BeforeFunc: func(c echo.Context) {
				// Extract the token from the Authorization header, and load the
				// key pair from the config struct.
				token := c.Request().Header.Get("Authorization")
				token = strings.Replace(token, "Bearer ", "", 1)
				token = strings.Replace(token, "bearer ", "", 1)
				token = strings.Replace(token, " ", "", 1)
				if token != "" {
					return
				}
				//
				// claims := jwt.Claims{}
				// _, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
				//	return privateEd25519Key, nil
				// })
				// if err != nil {
				//	h.app.Logger.Warn("unable to parse token")
				//	return
				// }
				//
				// if claims.ExpiresAt < time.Now().Unix() {
				//	h.app.Logger.Warn("token expired")
				//	return
				// }
				//
				// c.Set("user", claims)
			},
			TokenLookupFuncs: []middleware.ValuesExtractor{
				func(ectx echo.Context) ([]string, error) {
					token := ectx.Request().Header.Get("Authorization")
					token = strings.Replace(token, "Bearer ", "", 1)
					token = strings.Replace(token, "bearer ", "", 1)
					token = strings.Replace(token, " ", "", 1)

					return []string{token}, nil
				},
				func(ectx echo.Context) ([]string, error) {
					enc := encryption.New(h.app.Config.Encryption)

					c := cookie.New(h.app.Config.Cookie)
					token, err := c.ReadEncryptedCookies(ectx, enc, h.app.Config.Jwt.CookieTokenName)
					return []string{token}, err
				},
			},
			NewClaimsFunc: func(ectx echo.Context) golangJwt.Claims {
				return new(jwt.Claims)
			},
			ErrorHandler: func(ectx echo.Context, err error) error {
				h.app.Logger.Warn(err.Error())
				return err
			},
			ParseTokenFunc: func(ectx echo.Context, auth string) (interface{}, error) {
				jwtToken, err := jwt.New(h.app.Config.Jwt, h.app.Redis)
				if err != nil {
					return nil, err
				}

				data := new(map[string]interface{})

				err = jwtToken.ExtractData(auth, data)
				if err != nil {
					return nil, err
				}

				ectx.Set("user", data)
				return data, nil
			},
		}
	*/

	// Users handlers
	e.POST("/users/v1/signup", h.RegisterStart)
	e.POST("/users/v1/active", h.RegisterEnd)
	e.POST("/users/v1/login", h.Login)
	// e.POST("/v1/users/forget", h.createAuthenticationTokenHandler)

	// e.GET("/users/v1/me", h.Me, echojwt.WithConfig(config))
	// e.POST("/users/v1/logout", h.Logout, echojwt.WithConfig(config))
	// e.GET("/users/v1/:id", h.UserView, echojwt.WithConfig(config))
	// e.PUT("/users/v1/activate", h.RegisterEnd, echojwt.WithConfig(config))
}

func (h *UserHandler) RegisterStart(ectx echo.Context) error {
	v := validator.New(ectx)
	r, err := requests.NewRegisterStartRequest(ectx, v)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBadRequest):
			badRequestResponse(ectx, err)
			return nil
		case errors.Is(err, requests.ErrFailedValidation):
			failedValidationResponse(ectx, v.Errors)
			return nil
		default:
			serverErrorResponse(ectx, err)
			return nil
		}
	}

	user, err := models.NewUser(r.Email, r.Password)
	if err != nil {
		serverErrorResponse(ectx, err)
		return nil
	}

	err = h.userService.RegisterStart(user)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			failedValidationResponse(ectx, v.Errors)
		default:
			serverErrorResponse(ectx, err)
		}
		return nil
	}

	ectx.Response().Header().Add("Location", r.SuccessUrl)

	return ectx.NoContent(http.StatusAccepted)
}

func (h *UserHandler) RegisterEnd(ectx echo.Context) error {
	v := validator.New(ectx)
	r, err := requests.NewRegisterEndRequest(ectx, v)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBadRequest):
			badRequestResponse(ectx, err)
			return nil
		case errors.Is(err, requests.ErrFailedValidation):
			failedValidationResponse(ectx, v.Errors)
			return nil
		default:
			serverErrorResponse(ectx, err)
			return nil
		}
	}

	user, err := h.userService.ActivateUser(r.Token, r.UserId)
	if err != nil {
		serverErrorResponse(ectx, err)
		return nil
	}

	t, err := h.userService.Authorize(ectx, user)
	if err != nil {
		serverErrorResponse(ectx, err)
		return nil
	}

	return responses.WriteJSON(ectx, responses.NewMeResponse(user, t))
}

func (h *UserHandler) Login(ectx echo.Context) error {

	// Create token with claims
	jwtToken, err := jwt.New(h.app.Config.Jwt, h.app.Redis)
	if err != nil {
		return err
	}

	mapClaims := golangJwt.MapClaims{
		"username": "test",
		"id":       12,
	}
	t, err := jwtToken.GenerateToken(mapClaims)
	if err != nil {
		return err
	}

	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: echo.Map{
			"token": t,
		},
	}

	return ectx.JSON(http.StatusOK, res)
}

func (h *UserHandler) Logout(ectx echo.Context) error {
	jwtToken, err := jwt.New(h.app.Config.Jwt, h.app.Redis)
	if err != nil {
		return err
	}
	jwtToken.BlockToken("token")

	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"env":     h.app.Config.App.Env,
			"version": h.app.Config.App.Version,
		},
	}

	return ectx.JSON(http.StatusOK, res)
}
