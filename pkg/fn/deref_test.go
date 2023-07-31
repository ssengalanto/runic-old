package fn_test

import (
	"testing"
	"time"

	"github.com/ssengalanto/runic/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestDeref_NonPointerValue(t *testing.T) {
	assert.Panics(t, func() {
		fn.Deref("not a pointer")
	})
}

func TestDeref_DifferentTypes(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "deref nil *int (nil pointer)",
			input:    (*int)(nil),
			expected: 0,
		},
		{
			name:     "deref *int (pointer)",
			input:    func() *int { val := 42; return &val }(),
			expected: 42,
		},
		{
			name:     "deref *string (non-nil pointer)",
			input:    func() *string { val := "hello"; return &val }(),
			expected: "hello",
		},
		{
			name:     "deref nil *string (pointer)",
			input:    (*string)(nil),
			expected: "",
		},
		{
			name:     "deref nil *bool (nil pointer)",
			input:    (*bool)(nil),
			expected: false,
		},
		{
			name:     "deref *bool (pointer)",
			input:    func() *bool { val := true; return &val }(),
			expected: true,
		},
		{
			name:     "deref nil *uint (nil pointer)",
			input:    (*uint)(nil),
			expected: uint(0),
		},
		{
			name:     "deref *uint (pointer)",
			input:    func() *uint { val := uint(42); return &val }(),
			expected: uint(42),
		},
		{
			name:     "deref nil *float64 (nil pointer)",
			input:    (*float64)(nil),
			expected: float64(0),
		},
		{
			name:     "deref *float64 (pointer)",
			input:    func() *float64 { val := float64(42); return &val }(),
			expected: float64(42),
		},
		{
			name:     "deref nil *[]int (nil pointer)",
			input:    (*[]int)(nil),
			expected: []int(nil),
		},
		{
			name:     "deref *[]int (pointer)",
			input:    func() *[]int { val := []int{1, 2, 3}; return &val }(),
			expected: []int{1, 2, 3},
		},
		{
			name:     "deref *struct time.Time (nil-pointer)",
			input:    (*time.Time)(nil),
			expected: time.Time{},
		},
		{
			name:     "deref *struct time.Time (pointer)",
			input:    func() *time.Time { t := time.Date(2023, time.July, 29, 12, 0, 0, 0, time.UTC); return &t }(),
			expected: time.Date(2023, time.July, 29, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fn.Deref(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
