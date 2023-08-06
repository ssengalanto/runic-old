package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ssengalanto/runic/pkg/http/exceptions"
)

// Success sends an HTTP success json back to the client with the specified status code.
func Success(w http.ResponseWriter, statusCode int, payload any) {
	var res []byte
	var err error

	if payload != nil {
		res, err = json.Marshal(payload)
		if err != nil {
			panic(fmt.Sprintf("json.Success: %s", err))
		}
	}

	w.WriteHeader(statusCode)

	if len(res) > 0 {
		w.Write(res) //nolint:errcheck //unnecessary
	}
}

// Error sends an HTTP error json back to the client with the specified status code.
func Error(w http.ResponseWriter, err error) {
	if err == nil {
		err = errors.New("provided error is nil")
	}

	httpError := exceptions.NewHTTPError(err)

	res, err := json.Marshal(httpError)
	if err != nil {
		panic(fmt.Sprintf("json.Error: %s", err))
	}

	w.WriteHeader(httpError.Error.Code)

	if len(res) > 0 {
		w.Write(res) //nolint:errcheck //unnecessary
	}
}
