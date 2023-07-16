// Package mux provides a convenience wrapper around the go-chi/chi router to create an HTTP route multiplexer with predefined configurations.
//
// The mux package offers a simple and opinionated way to create a chi router with commonly used middlewares and configurations.
// It provides a New function that returns a *chi.Mux instance with predefined middleware and settings.
//
// The returned router has the following middleware already applied:
// - Recovery middleware to recover from panics and handle internal server errors.
// - RequestID middleware to generate unique request IDs for each incoming request.
// - RealIP middleware to extract the real IP address from X-Real-IP and X-Forwarded-For headers.
// - Logger middleware to log request and response details.
// - Compression middleware to compress response bodies using gzip with a specified compression level.
// - CORS (Cross-Origin Resource Sharing) middleware to enable cross-origin requests with fine-grained configuration options.
// - IP rate limiting middleware to limit the number of requests per IP address within a specified duration.
// - Timeout middleware to set a maximum timeout duration for requests.
// - JSON content type middleware to set the "Content-Type" header of the response to "application/json".
// - Heartbeat middleware to provide a heartbeat route for basic health checks.
package mux
