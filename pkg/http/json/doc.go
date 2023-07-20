// Package json provides utilities for encoding and decoding JSON data in Go.
//
// This package wraps the standard encoding/json package to add additional features
// and error handling for JSON encoding and decoding operations. It includes functions
// to handle JSON decoding from HTTP request bodies, enforce strict JSON parsing,
// and limit the size of the request body to prevent potential attacks.
package json
