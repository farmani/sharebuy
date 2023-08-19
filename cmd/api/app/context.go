package app

import (
	"context"
	"github.com/farmani/sharebuy/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiContext struct {
	echo.Context
	app *Application
}

func (c *ApiContext) App() *Application {
	return c.app
}

type contextKey string

const userContextKey = contextKey("user")

func (app *Application) contextSetUser(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *Application) contextGetUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(userContextKey).(*models.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
