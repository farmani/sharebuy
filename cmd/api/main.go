package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}

	e := echo.New()
	e.Server.IdleTimeout = time.Minute
	e.Server.ReadTimeout = time.Second * 15
	e.Server.WriteTimeout = time.Second * 30

	e.GET("/health", app.healthHandler)

	addr := fmt.Sprintf(":%d", cfg.port)
	e.Logger.Debug("Starting %s server on %s", cfg.env, addr)
	err := e.Start(addr)
	e.Logger.Fatal(err)
}
