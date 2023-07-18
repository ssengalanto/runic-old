package exception_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/runic/pkg/exception"
	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *exception.Error
		expected string
	}{
		{
			name:     "NoCause",
			err:      exception.New("custom error message"),
			expected: "custom error message",
		},
		{
			name:     "WithCause",
			err:      exception.Wrap(fmt.Errorf("underlying error"), "custom error message"),
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
		err      *exception.Error
		expected error
	}{
		{
			name:     "NoCause",
			err:      exception.New("custom error message"),
			expected: nil,
		},
		{
			name:     "WithCause",
			err:      exception.Wrap(fmt.Errorf("underlying error"), "custom error message"),
			expected: fmt.Errorf("underlying error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.err.Unwrap())
		})
	}
}
