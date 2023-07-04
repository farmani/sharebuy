package app

import (
	"context"
	"net/http"

	"github.com/farmani/sharebuy/internal/data"
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

func (app *Application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *Application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
