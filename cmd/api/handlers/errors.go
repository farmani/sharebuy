package handlers

import (
	"fmt"
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/labstack/echo/v4"
)

func logError(c echo.Context, err error) {
	c.Logger().Error(err, map[string]string{
		"request_method": c.Request().Method,
		"request_url":    c.Request().RequestURI,
	})
}

func errorResponse(c echo.Context, status int, message interface{}) {
	res := app.Envelope{
		Status: "ERROR",
		Code:   status,
		Data: map[string]interface{}{
			"message": message,
		},
	}

	err := c.JSON(status, res)
	if err != nil {
		logError(c, err)
	}
}

func serverErrorResponse(c echo.Context, err error) {
	logError(c, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(c, http.StatusInternalServerError, message)
}

func notFoundResponse(c echo.Context) {
	message := "the requested resource could not be found"
	errorResponse(c, http.StatusNotFound, message)
}

func methodNotAllowedResponse(c echo.Context) {
	message := fmt.Sprintf("the %s method is not supported this resource", c.Request().Method)
	errorResponse(c, http.StatusMethodNotAllowed, message)
}

func badRequestResponse(c echo.Context, err error) {
	errorResponse(c, http.StatusBadRequest, err.Error())
}

func failedValidationResponse(c echo.Context, errors map[string]string) {
	errorResponse(c, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse sends a JSON-formatted error message to the client with a 409 Conflict
// status code.
func editConflictResponse(c echo.Context) {
	message := "unable to update the record due to an edit conflict, please try again"
	errorResponse(c, http.StatusConflict, message)
}

func rateLimitExceededResponse(c echo.Context) {
	message := "rate limited exceeded"
	errorResponse(c, http.StatusTooManyRequests, message)
}

func invalidCredentialsResponse(c echo.Context) {
	message := "invalid authentication credentials"
	errorResponse(c, http.StatusUnauthorized, message)
}

func invalidAuthenticationTokenResponse(c echo.Context) {
	c.Response().Header().Set("WWWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	errorResponse(c, http.StatusUnauthorized, message)
}

func authenticationRequiredResponse(c echo.Context) {
	message := "you must be authenticated to access this resource"
	errorResponse(c, http.StatusUnauthorized, message)
}

func inactiveAccountResponse(c echo.Context) {
	message := "your user account must be activated to access this resource"
	errorResponse(c, http.StatusForbidden, message)
}

func notPermittedResponse(c echo.Context) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	errorResponse(c, http.StatusForbidden, message)
}
