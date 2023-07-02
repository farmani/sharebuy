package main

import (
	"github.com/farmani/sharebuy/internal/common/config"
)

var version = "1.0.0"

func main() {
	// Declare an instance of the config struct.
	var cfg config.Config
	cfg.ParseFlags()

	app := NewApiApplication(cfg)
	app.bootstrap()

	app.serve()
}
