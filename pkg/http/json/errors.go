package json

import (
	"errors"
)

var (
	// ErrBadlyFormedJSON is returned when the request body contains a JSON that is not properly formatted.
	ErrBadlyFormedJSON = errors.New("request body contains badly formed json")

	// ErrIncorrectJSONType is returned when the request body contains a JSON value of incorrect type.
	ErrIncorrectJSONType = errors.New("request body contains incorrect json type")

	// ErrEmptyBody is returned when the request body is empty.
	ErrEmptyBody = errors.New("request body is empty")

	// ErrInvalidJSON is returned when the request body is an invalid JSON.
	ErrInvalidJSON = errors.New("invalid json")

	// ErrRequestBodyTooLarge is returned when the request body size exceeds the maximum allowed limit.
	ErrRequestBodyTooLarge = errors.New("request body is too large")

	// ErrUnexpectedEOF is the error returned when the JSON decoding encounters an unexpected EOF.
	ErrUnexpectedEOF = errors.New("unexpected end of JSON input")
)
