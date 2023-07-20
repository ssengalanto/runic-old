package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ssengalanto/runic/pkg/constants"
)

// DecodeRequest decodes the JSON payload from an HTTP request and stores it in the given destination.
// The function also enforces a maximum request body size to prevent large payloads.
// It returns an error if the decoding fails.
func DecodeRequest(w http.ResponseWriter, r *http.Request, dst any) error {
	// Limit the request body size to prevent potential abuse or attacks.
	r.Body = http.MaxBytesReader(w, r.Body, int64(constants.MaxHeaderBytes))

	// Create a JSON decoder and disallow unknown fields to ensure strict JSON decoding.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the JSON payload into the provided destination (dst).
	err := dec.Decode(dst)
	if err != nil {
		err = handleDecodeError(err)
		return err
	}

	return nil
}

// handleDecodeError handles the error from decoding JSON request body.
func handleDecodeError(err error) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		return fmt.Errorf("%w (at character %d)", ErrBadlyFormedJSON, syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		return ErrUnexpectedEOF

	case errors.As(err, &unmarshalTypeError):
		return fmt.Errorf("%w for field %q", ErrIncorrectJSONType, unmarshalTypeError.Field)

	case errors.Is(err, io.EOF):
		return ErrEmptyBody

	case err.Error() == "http: request body too large":
		return fmt.Errorf("%w, it must not be larger than %d bytes", ErrRequestBodyTooLarge, constants.MaxHeaderBytes)

	default:
		return fmt.Errorf("%w: %s", ErrInvalidJSON, err.Error())
	}
}
