package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/farmani/sharebuy/internal/logger"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (app *Application) Bootstrap() error {
	app.Logger = logger.NewZapLogger(app.Config.LogPath, app.Config.Env)

	err := app.openDB()
	if err != nil {
		return fmt.Errorf("open DB Failed: %w", err)
	}

	err = app.openRedis()
	if err != nil {
		return fmt.Errorf("open Redis Failed: %w", err)
	}

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
	db, err := sql.Open("mysql", app.Config.Db.Dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(app.Config.Db.MaxOpenConns)
	db.SetMaxIdleConns(app.Config.Db.MaxIdleConns)
	duration, err := time.ParseDuration(app.Config.Db.MaxIdleTime)
	if err != nil {
		return err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	boil.SetDB(db)

	app.Db = db

	err = app.Db.Ping()
	if err != nil {
		return err
	}

	app.Logger.Info("database connection pool established")

	return nil
}

func (app *Application) openRedis() error {
	app.Redis = redis.NewClient(&redis.Options{
		Addr:     app.Config.Redis.Addr,
		Password: app.Config.Redis.Password, // no password set
		DB:       app.Config.Redis.Db,       // use default DB
	})

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	app.Redis.Ping()

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
