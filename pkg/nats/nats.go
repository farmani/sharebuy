package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func New(cfg *Config) (*nats.Conn, error) {

	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		return nil, fmt.Errorf("parse nats timeout duration Failed: %w", err)
	}

	drainTimeout, err := time.ParseDuration(cfg.DrainTimeout)
	if err != nil {
		return nil, fmt.Errorf("parse nats drain timeout duration Failed: %w", err)
	}

	flusherTimeout, err := time.ParseDuration(cfg.FlusherTimeout)
	if err != nil {
		return nil, fmt.Errorf("parse nats flusher timeout duration Failed: %w", err)
	}

	pingInterval, err := time.ParseDuration(cfg.PingInterval)
	if err != nil {
		return nil, fmt.Errorf("parse nats ping interval duration Failed: %w", err)
	}

	reconnectWait, err := time.ParseDuration(cfg.ReconnectWait)
	if err != nil {
		return nil, fmt.Errorf("parse nats reconnect wait duration Failed: %w", err)
	}

	nc, err := nats.Connect(
		cfg.Url,
		nats.Timeout(timeout),
		nats.DrainTimeout(drainTimeout),
		nats.FlusherTimeout(flusherTimeout),
		nats.PingInterval(pingInterval),
		nats.ReconnectWait(reconnectWait),
		nats.MaxReconnects(cfg.MaxReconnects),
		nats.MaxPingsOutstanding(cfg.MaxPingsOutstanding),
		nats.ReconnectBufSize(8388608),
		nats.Compression(cfg.Compression),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {}),
		nats.ReconnectHandler(func(nc *nats.Conn) {}),
		nats.ClosedHandler(func(nc *nats.Conn) {}),
		nats.DiscoveredServersHandler(func(nc *nats.Conn) {}),
		nats.UseOldRequestStyle(),
		nats.NoEcho(),
		nats.UserCredentials(""),
		nats.Token(""),
	)

	if err != nil {
		return nil, err
	}

	return nc, nil
}
