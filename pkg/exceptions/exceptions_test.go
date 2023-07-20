package exceptions_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *exceptions.Error
		expected string
	}{
		{
			name:     "no cause",
			err:      exceptions.New("custom error message"),
			expected: "custom error message",
		},
		{
			name:     "with cause",
			err:      exceptions.Wrap(fmt.Errorf("underlying error"), "custom error message"),
			expected: "custom error message: underlying error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.err.Error())
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	tests := []struct {
		name     string
		err      *exceptions.Error
		expected error
	}{
		{
			name:     "no cause",
			err:      exceptions.New("custom error message"),
			expected: nil,
		},
		{
			name:     "with cause",
			err:      exceptions.Wrap(fmt.Errorf("underlying error"), "custom error message"),
			expected: fmt.Errorf("underlying error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.err.Unwrap())
		})
	}
}
