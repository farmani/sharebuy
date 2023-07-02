package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/farmani/sharebuy/internal/jsonlog"
	"github.com/redis/go-redis/v9"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (app *application) bootstrap() error {
	app.logger = jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

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

func (app *application) openNats() error {
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

	app.nats = nc

	return nil
}

func (app *application) openDB() error {
	db, err := sql.Open("mysql", app.config.Db.Dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(app.config.Db.MaxOpenConns)
	db.SetMaxIdleConns(app.config.Db.MaxIdleConns)
	duration, err := time.ParseDuration(app.config.Db.MaxIdleTime)
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

	app.db = db

	err = app.db.Ping()
	if err != nil {
		return err
	}

	app.logger.PrintInfo("database connection pool established", nil)

	return nil
}

func (app *application) openRedis() error {
	app.redis = redis.NewClient(&redis.Options{
		Addr:     app.config.Redis.Addr,
		Password: app.config.Redis.Password, // no password set
		DB:       app.config.Redis.Db,       // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.redis.Ping(ctx)

	return nil
}

func (app *application) Shutdown() error {
	err := app.db.Close()
	if err != nil {
		return err
	}

	err = app.redis.Close()
	if err != nil {
		return err
	}

	app.nats.Close()

	return nil
}
