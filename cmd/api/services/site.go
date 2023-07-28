package services

import (
	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/internal/repository"
	"github.com/labstack/echo/v4"
)

type SiteService struct {
	// Dependencies or state for the UserHandler
	app  *app.Application
	repo repository.Repository
}

func NewSiteService(a *app.Application, repo repository.Repository) *SiteService {
	return &SiteService{
		app:  a,
		repo: repo,
	}
}

func (s *UserService) HealthCheck(c echo.Context) error {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	return nil
}
