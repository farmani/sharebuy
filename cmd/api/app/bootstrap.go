package app

import (
	"fmt"
	"github.com/farmani/sharebuy/pkg/rdbms"
	"go.uber.org/zap"
	"time"

	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/farmani/sharebuy/pkg/redis"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

func (app *Application) Bootstrap() error {
	app.Logger = logger.NewZapLogger(app.Config.Logger.Path, app.Config.App.Env)
	var err error

	err = app.openRedis()
	if err != nil {
		return fmt.Errorf("open Redis Failed: %w", err)
	}
	err = app.openDB()
	if err != nil {
		return fmt.Errorf("open DB Failed: %w", err)
	}

	return nil
	err = app.openNats()
	if err != nil {
		return fmt.Errorf("open NATS Failed: %w", err)
	}

	return nil
}

func (app *Application) openNats() error {
	nc, err := nats.Connect(
		nats.DefaultURL,
		nats.Timeout(5*time.Second),
		nats.MaxReconnects(5),
		nats.ReconnectWait(5*time.Second),
		nats.ReconnectBufSize(8388608),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {}),
		nats.ReconnectHandler(func(nc *nats.Conn) {}),
		nats.ClosedHandler(func(nc *nats.Conn) {}),
		nats.DiscoveredServersHandler(func(nc *nats.Conn) {}),
		nats.PingInterval(5*time.Second),
		nats.MaxPingsOutstanding(2),
		nats.UseOldRequestStyle(),
		nats.NoEcho(),
		nats.UserCredentials(""),
		nats.Token(""),
	)
	if err != nil {
		return err
	}

	app.Nats = nc

	return nil
}

func (app *Application) openDB() error {
	var err error
	app.Db, err = rdbms.New(app.Config.Db)
	if err != nil {
		app.Logger.Panic("database connection establishment failed", zap.Error(err))
	}

	app.Logger.Info("database connection pool established")

	return nil
}

func (app *Application) openRedis() error {
	var err error
	app.Redis, err = redis.New(app.Config.Redis)
	if err != nil {
		app.Logger.Panic("redis connection establishment failed", zap.Error(err))
	}

	app.Logger.Info("redis connection pool established")

	return nil
}

func (app *Application) Shutdown() error {
	err := app.Db.Close()
	if err != nil {
		return err
	}

	err = app.Redis.Close()
	if err != nil {
		return err
	}

	app.Nats.Close()

	return nil
}
