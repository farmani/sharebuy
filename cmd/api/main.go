package main

import (
	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/handlers"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/internal/common/config"
)

var version = "1.0.0"

func main() {
	// Declare an instance of the config struct.
	var cfg config.Config
	cfg.ParseFlags()

	application, err := app.NewApiApplication(cfg)
	if err != nil {
		panic(err)
	}

	if err := application.Bootstrap(); err != nil {
		// panic(err)
	}

	// Add the handlers to the Application
	application.Services = map[string]app.Service{
		"UserService": services.NewUserService(application),
	}

	// Add the handlers to the Application
	application.Handlers = []app.Handler{
		handlers.NewSiteHandler(application),
		handlers.NewUserHandler(application),
		// Add other handlers here
	}

	if err := application.Start(); err != nil {
		panic(err)
	}
}
