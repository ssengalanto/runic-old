package redis

import (
	"context"
	"fmt"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/ssengalanto/runic/pkg/constants"
)

// NewUniversalClient creates a new redis universal client.
func NewUniversalClient(host, port, pwd string, db int) (redis.UniversalClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	hp := net.JoinHostPort(host, port)
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       []string{hp},
		Password:    pwd,
		DB:          db,
		MaxRetries:  maxRetries,
		DialTimeout: dialTimeout,
	})

	_, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	return c, nil
}
