package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrDocumentNotFound = errors.New("DocumentNotFound")

func NewErrorStatusCodeMaps() map[error]int {

	var errorStatusCodeMaps = make(map[error]int)
	errorStatusCodeMaps[ErrDocumentNotFound] = http.StatusNotFound
	return errorStatusCodeMaps
}

type httpErrorHandler struct {
	statusCodes map[error]int
}

func NewHttpErrorHandler(errorStatusCodeMaps map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		statusCodes: errorStatusCodeMaps,
	}
}
func (self *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range self.statusCodes {
		if errors.Is(err, key) {
			return value
		}
	}

	return http.StatusInternalServerError
}

func unwrapRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}

func (self *httpErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: unwrapRecursive(err).Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}

func (app *Application) logError(c echo.Context, err error) {
	c.Logger().Error(err, map[string]string{
		"request_method": c.Request().Method,
		"request_url":    c.Request().RequestURI,
	})
}

func (app *Application) errorResponse(c echo.Context, status int, message interface{}) {
	res := Envelope{
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

func (app *Application) serverErrorResponse(c echo.Context, err error) {
	app.logError(c, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(c, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(c echo.Context) {
	message := "the requested resource could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(c echo.Context) {
	message := fmt.Sprintf("the %s method is not supported this resource", c.Request().Method)
	app.errorResponse(c, http.StatusMethodNotAllowed, message)
}

func (app *Application) badRequestResponse(c echo.Context, err error) {
	app.errorResponse(c, http.StatusBadRequest, err.Error())
}

func (app *Application) failedValidationResponse(c echo.Context, errors map[string]string) {
	app.errorResponse(c, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse sends a JSON-formatted error message to the client with a 409 Conflict
// status code.
func (app *Application) editConflictResponse(c echo.Context) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(c, http.StatusConflict, message)
}

func (app *Application) rateLimitExceededResponse(c echo.Context) {
	message := "rate limited exceeded"
	app.errorResponse(c, http.StatusTooManyRequests, message)
}

func (app *Application) invalidCredentialsResponse(c echo.Context) {
	message := "invalid authentication credentials"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) invalidAuthenticationTokenResponse(c echo.Context) {
	c.Response().Header().Set("WWWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) authenticationRequiredResponse(c echo.Context) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) inactiveAccountResponse(c echo.Context) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(c, http.StatusForbidden, message)
}

func (app *Application) notPermittedResponse(c echo.Context) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorResponse(c, http.StatusForbidden, message)
}
