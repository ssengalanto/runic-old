package constants

import "time"

const (
	// MaxHeaderBytes represents the maximum number of bytes allowed for HTTP headers.
	MaxHeaderBytes = 1 << 20

	// IdleTimeout represents the maximum duration of idle keep-alive connections.
	IdleTimeout = time.Minute

	// ReadTimeout represents the maximum duration for reading the entire request.
	ReadTimeout = 10 * time.Second

	// WriteTimeout represents the maximum duration for writing the entire response.
	WriteTimeout = 30 * time.Second

	// RateLimit represents the maximum number of requests allowed per second.
	RateLimit = 100

	// Timeout represents the default duration for request timeouts.
	Timeout = 10 * time.Second

	// ShutdownGracePeriod represents the duration to allow for graceful shutdown of the server.
	ShutdownGracePeriod = 15 * time.Second

	// CorsMaxAge represents the maximum duration to cache CORS (Cross-Origin Resource Sharing) responses.
	CorsMaxAge = 300

	// GzipCompressionLevel represents the compression level for Gzip compression.
	GzipCompressionLevel = 5

	// DevEnv represents the development environment.
	DevEnv = "development"

	// StgEnv represents the staging environment.
	StgEnv = "staging"

	// ProdEnv represents the production environment.
	ProdEnv = "production"
)
