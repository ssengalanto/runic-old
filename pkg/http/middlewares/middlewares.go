package middlewares

import "net/http"

// JSONContentType is a middleware that sets the "Content-Type" header of the response to "application/json".
// It wraps the provided http.Handler and delegates the request/response processing
// to the next middleware or handler in the chain.
func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
