package services

import (
	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/labstack/echo/v4"
)

type SiteService struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewSiteService(a *app.Application) *SiteService {
	return &SiteService{
		app: a,
	}
}

func (us *UserService) HealthCheck(c echo.Context) error {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	return nil
}
