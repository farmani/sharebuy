package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/requests"
	"github.com/farmani/sharebuy/cmd/api/responses"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/pkg/cookie"
	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/farmani/sharebuy/pkg/jwt"
	"github.com/farmani/sharebuy/pkg/validator"
	golangJwt "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
			//claims := jwt.Claims{}
			//_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			//	return privateEd25519Key, nil
			//})
			//if err != nil {
			//	h.app.Logger.Warn("unable to parse token")
			//	return
			//}
			//
			//if claims.ExpiresAt < time.Now().Unix() {
			//	h.app.Logger.Warn("token expired")
			//	return
			//}
			//
			//c.Set("user", claims)
		},
		TokenLookupFuncs: []middleware.ValuesExtractor{
			func(c echo.Context) ([]string, error) {
				token := c.Request().Header.Get("Authorization")
				token = strings.Replace(token, "Bearer ", "", 1)
				token = strings.Replace(token, "bearer ", "", 1)
				token = strings.Replace(token, " ", "", 1)

				return []string{token}, nil
			},
			func(c echo.Context) ([]string, error) {
				enc := encryption.New(h.app.Config.Encryption)

				cookie := cookie.New(h.app.Config.Cookie)
				token, err := cookie.ReadEncryptedCookies(c, enc, h.app.Config.Jwt.CookieTokenName)
				return []string{token}, err
			},
		},
		NewClaimsFunc: func(c echo.Context) golangJwt.Claims {
			return new(jwt.Claims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			h.app.Logger.Warn(err.Error())
			return err
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			jwtToken, err := jwt.New(h.app.Config.Jwt, h.app.Redis)
			if err != nil {
				return nil, err
			}

			data := new(map[string]interface{})

			err = jwtToken.ExtractData(auth, data)
			if err != nil {
				return nil, err
			}

			c.Set("user", data)
			return data, nil
		},
	}

	// Users handlers
	e.POST("/v1/users", h.RegisterStart)
	e.PUT("/v1/users/activated", h.RegisterEnd)
	e.POST("/v1/users/login", h.Login)
	// e.POST("/v1/users/forget", h.createAuthenticationTokenHandler)

	e.GET("/v1/users/:id", h.UserView, echojwt.WithConfig(config))
	e.GET("/v1/users/me", h.Me, echojwt.WithConfig(config))
	e.POST("/v1/users/logout", h.Logout, echojwt.WithConfig(config))
}

func (h *UserHandler) UserView(c echo.Context) error {
	//jwtUser := c.Get("user").(*jwt.Token)
	//jwt.Token.ExtractData(c.Get("user").(*jwt.Token), jwtUser)
	//id, _ := strconv.Atoi(claims["id"].(string))
	user, err := h.userService.FindById(1)
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

func (h *UserHandler) RegisterStart(c echo.Context) error {
	v := validator.New(c)
	r, err := requests.NewRegisterStartRequest(c, v)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBadRequest):
			badRequestResponse(c, err)
			return nil
		case errors.Is(err, requests.ErrFailedValidation):
			failedValidationResponse(c, v.Errors)
			return nil
		default:
			serverErrorResponse(c, err)
			return nil
		}
	}

	/*
		us.RegisterStart(c)
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
	*/

	// Create token with claims
	jwtToken, err := jwt.New(h.app.Config.Jwt, h.app.Redis)
	if err != nil {
		serverErrorResponse(c, err)
		return nil
	}

	mapClaims := golangJwt.MapClaims{
		"email": r.Email,
		"id":    12,
	}
	t, err := jwtToken.GenerateToken(mapClaims)
	if err != nil {
		serverErrorResponse(c, err)
		return nil
	}

	res := app.Envelope{
		Status: "OK",
		Code:   200,
		Data: map[string]interface{}{
			"token": t,
		},
	}

	enc := encryption.New(h.app.Config.Encryption)

	cookie := cookie.New(h.app.Config.Cookie)
	cookie.SetEncryptedCookies(c, enc, h.app.Config.Jwt.CookieTokenName, t)

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
	user := c.Get("user")
	fmt.Printf("%v %T", user, user)
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

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Logout(c echo.Context) error {
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
