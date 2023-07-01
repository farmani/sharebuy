package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) logError(c echo.Context, err error) {
	c.Logger().Error(err, map[string]string{
		"request_method": c.Request().Method,
		"request_url":    c.Request().RequestURI,
	})
}

func (app *application) errorResponse(c echo.Context, status int, message interface{}) {
	res := envelope{
		Status: "ERROR",
		Code:   status,
		Data: map[string]interface{}{
			"message": message,
		},
	}

	err := c.JSON(status, res)
	if err != nil {
		app.logError(c, err)
	}
}

func (app *application) serverErrorResponse(c echo.Context, err error) {
	app.logError(c, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(c, 500, message)
}

func (app *application) notFoundResponse(c echo.Context) {
	message := "the requested resource could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(c echo.Context) {
	message := fmt.Sprintf("the %s method is not supported this resource", c.Request().Method)
	app.errorResponse(c, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(c echo.Context, err error) {
	app.errorResponse(c, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(c echo.Context, errors map[string]string) {
	app.errorResponse(c, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse sends a JSON-formatted error message to the client with a 409 Conflict
// status code.
func (app *application) editConflictResponse(c echo.Context) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(c, http.StatusConflict, message)
}

func (app *application) rateLimitExceededResponse(c echo.Context) {
	message := "rate limited exceeded"
	app.errorResponse(c, http.StatusTooManyRequests, message)
}

func (app *application) invalidCredentialsResponse(c echo.Context) {
	message := "invalid authentication credentials"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) invalidAuthenticationTokenResponse(c echo.Context) {
	c.Response().Header().Set("WWWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) authenticationRequiredResponse(c echo.Context) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(c echo.Context) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(c, http.StatusForbidden, message)
}

func (app *application) notPermittedResponse(c echo.Context) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorResponse(c, http.StatusForbidden, message)
}
