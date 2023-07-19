// Package exceptions provides custom error types and utilities for handling HTTP errors.
//
// The exceptions package offers a set of custom error types that represent various common HTTP errors.
// These error types allow developers to create well-defined error instances with specific HTTP status
// codes, messages, and causes. The package also includes a utility function, NewHTTPError, that takes
// a standard error and converts it into an HTTPError with an appropriate error structure based on the
// type of the original error.
package exceptions
