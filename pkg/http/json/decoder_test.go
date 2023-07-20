package json_test

import (
	"bytes"
	stdjson "encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/runic/pkg/http/json"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a new HTTP request with the given payload.
func newRequest(payload string) (*http.Request, error) {
	return http.NewRequest("POST", "/example", bytes.NewBufferString(payload))
}

// Test payload struct.
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Helper function to create a destination struct for decoding.
func newDestination() *Data {
	return &Data{}
}

// TestDecodeRequest tests the DecodeRequest function.
func TestDecodeRequest(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		payload := `{"name": "John", "age": 30}`
		req, err := newRequest(payload)
		assert.NoError(t, err)

		data := newDestination()

		err = json.DecodeRequest(nil, req, data)

		assert.NoError(t, err)
		assert.Equal(t, "John", data.Name)
		assert.Equal(t, 30, data.Age)
	})

	t.Run("incorrect json type", func(t *testing.T) {
		payload := `{"name": "John", "age": "thirty"}`
		req, err := newRequest(payload)
		assert.NoError(t, err)

		data := newDestination()

		err = json.DecodeRequest(nil, req, data)

		assert.True(t, errors.Is(err, json.ErrIncorrectJSONType))
	})

	t.Run("empty request body", func(t *testing.T) {
		payload := ``
		req, err := newRequest(payload)
		assert.NoError(t, err)

		data := newDestination()

		err = json.DecodeRequest(nil, req, data)

		assert.True(t, errors.Is(err, json.ErrEmptyBody))
	})

	t.Run("badly-formed json", func(t *testing.T) {
		payload := `{"name": "John, "age": 30, "occupation": "developer", "city": "New York"}`
		req, err := newRequest(payload)
		assert.NoError(t, err)

		data := newDestination()

		err = json.DecodeRequest(nil, req, data)

		assert.True(t, errors.Is(err, json.ErrBadlyFormedJSON))
	})

	t.Run("unexpected end of file", func(t *testing.T) {
		payload := `{"name": "John", "age":`
		req, err := newRequest(payload)
		assert.NoError(t, err)

		data := newDestination()

		err = json.DecodeRequest(nil, req, data)

		assert.True(t, errors.Is(err, json.ErrUnexpectedEOF))
	})

	t.Run("request body too large", func(t *testing.T) {
		gofakeit.Seed(0)
		paragraph := gofakeit.Paragraph(100, 10, 1000, ".")

		payloadObj := map[string]string{"data": paragraph}
		largeJSONPayload, err := stdjson.Marshal(payloadObj)
		assert.NoError(t, err)

		req, err := newRequest(string(largeJSONPayload))
		assert.NoError(t, err)

		data := struct {
			Data string `json:"data"`
		}{}

		err = json.DecodeRequest(httptest.NewRecorder(), req, &data)

		assert.True(t, errors.Is(err, json.ErrRequestBodyTooLarge))
	})
}
