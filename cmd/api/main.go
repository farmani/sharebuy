package main

import (
	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/handlers"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/common/config"
)

var version = "1.0.0"

func main() {
	application := app.NewApiApplication(config.NewConfig())

	// Add the handlers to the Application
	application.Handlers = []app.Handler{
		handlers.NewSiteHandler(application, services.NewSiteService(application)),
		handlers.NewUserHandler(application, services.NewUserService(application)),
		// Add other handlers here
	}

	if err := application.Start(); err != nil {
		panic(err)
	}
}
