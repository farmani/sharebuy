package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/farmani/sharebuy/internal/common/config"
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/internal/jsonlog"
	"github.com/farmani/sharebuy/internal/mailer"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

type application struct {
	config config.Config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
	db     *sql.DB
	redis  *redis.Client
	nats   *nats.Conn
}

func NewApiApplication(cfg config.Config) *application {
	return &application{
		config: cfg,
	}
}

func (app *application) serve() error {
	e := echo.New()
	e.Server.IdleTimeout = time.Minute
	e.Server.ReadTimeout = time.Second * 15
	e.Server.WriteTimeout = time.Second * 30
	app.routes(e)

	addr := fmt.Sprintf(":%d", app.config.App.Port)

	shutdownError := make(chan error)
	// Start a background goroutine.
	go func() {
		// Create a quit channel which carries os.Signal values. Use buffered
		quit := make(chan os.Signal, 1)
		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and relay
		// them to the quit channel. Any other signal will not be caught by signal.Notify()
		// and will retain their default behavior.
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		// Read the signal from the quit channel. This code will block until a signal is
		// received.
		s := <-quit
		// Log a message to say we caught the signal. Notice that we also call the
		// String() method on the signal to get the signal name and include it in the log
		// entry properties.
		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		// Create a context with a 5-second timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// call Shutdown on the server, and only send on the shutdownError channel if it returns
		// an error
		// Create a shutdownError channel. We will use this to receive any errors returned
		// by the graceful Shutdown() function.
		if err := e.Shutdown(ctx); err != nil {
			shutdownError <- err
		}
		// Log a message to say that we're waiting for any background goroutines to complete
		// their tasks.
		app.logger.PrintInfo("completing background tasks", map[string]string{
			"addr": e.Server.Addr,
		})

		// Call Wait() to block until our WaitGroup counter is zero. This essentially blocks
		// until the background goroutines have finished. Then we return nil on the shutdownError
		// channel to indicate that the shutdown as compleeted without any issues.
		app.wg.Wait()
		shutdownError <- nil

	}()

	// Log a "starting server" message.
	app.logger.PrintInfo("starting server", map[string]string{
		"addr": e.Server.Addr,
		"env":  app.config.App.Env,
	})

	// Calling Shutdown() on our server will cause ListenAndServer() to immediately
	// return a http.ErrServerClosed error. So, if we see this error, it is actually a good thing
	// and an indication that the graceful shutdown has started. So, we specifically check for this,
	// only returning the error if it is NOT http.ErrServerClosed.

	e.Logger.Debug("Starting %s server on %s", app.config.App.Env, e.Server.Addr)
	if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

	// Otherwise, we wait to receive the return value from Shutdown() on the shutdownErr
	// channel. If the return value is an error, we know that there was a problem with the
	// graceful shutdown, and we return the error.
	err := <-shutdownError
	if err != nil {
		return err
	}

	// At this point we know that the graceful shutdown completed successfully, and we log
	// a "stopped server" message.
	app.logger.PrintInfo("Stopped server", map[string]string{
		"addr": e.Server.Addr,
	})

	return nil
}
