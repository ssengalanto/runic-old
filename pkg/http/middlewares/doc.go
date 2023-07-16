// Package middlewares provides HTTP middleware functions for request/response processing.
//
// The middlewares package provides a collection of HTTP middleware functions that can be used to
// modify or inspect incoming requests and outgoing responses. These middlewares can be applied to
// specific routes or globally to the entire application to perform various tasks.
//
// Each middleware in this package follows the standard HTTP middleware pattern, accepting an
// `http.Handler` and returning a new `http.Handler` that wraps the original handler, allowing it to
// intercept and modify the request/response flow.
package middlewares
