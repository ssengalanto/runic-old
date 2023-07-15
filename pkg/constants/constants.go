// Package constants provides constant values that are used throughout the application.
package constants

import "time"

const (
	MaxHeaderBytes       = 1 << 20
	IdleTimeout          = time.Minute
	ReadTimeout          = 10 * time.Second
	WriteTimeout         = 30 * time.Second
	RateLimit            = 100
	Timeout              = time.Minute
	ShutdownGracePeriod  = 15 * time.Second
	CorsMaxAge           = 300
	GzipCompressionLevel = 5
)
