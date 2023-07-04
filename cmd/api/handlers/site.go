package handlers

import (
	"net/http"

	"github.com/farmani/sharebuy/cmd/api/app"
	echoPrometheus "github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type SiteHandler struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewSiteHandler(a *app.Application) *SiteHandler {
	return &SiteHandler{
		app: a,
	}
}

func (h *SiteHandler) RegisterRoutes(e *echo.Echo) {
	// Define routes specific to the UserHandler

	// healthcheck
	e.GET("/v1/health", h.Health)

	// adds middleware to gather metrics from those endpoints and expose them on /metrics
	e.GET("/metrics", echoPrometheus.NewHandler())
}

func (h *SiteHandler) Health(c echo.Context) error {
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

func (h *SiteHandler) Metrics(c echo.Context) error {
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
