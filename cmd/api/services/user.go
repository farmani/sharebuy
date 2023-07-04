package services

import (
	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewUserService(a *app.Application) *UserService {
	return &UserService{
		app: a,
	}
}

func (s *UserService) Cast() interface{} {
	return s
}

func (s *UserService) RegisterStart(c echo.Context) error {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	return nil
}
