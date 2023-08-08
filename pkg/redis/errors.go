package redis

import (
	"fmt"
)

// ErrConnectionFailed represents an error that is returned when a redis client connection attempt fails.
var ErrConnectionFailed = fmt.Errorf("redis client connection failed")
