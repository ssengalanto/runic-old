package validator_test

import (
	"testing"
	"time"

	"github.com/ssengalanto/runic/pkg/validator"
	"github.com/stretchr/testify/require"
)

func TestIsNonZeroed(t *testing.T) {
	type Data struct {
		StringField string         `validate:"nz"`
		BoolField   bool           `validate:"nz"`
		IntField    int            `validate:"nz"`
		UintField   uint           `validate:"nz"`
		FloatField  float64        `validate:"nz"`
		ArrayField  [3]int         `validate:"nz"`
		SliceField  []string       `validate:"nz"`
		MapField    map[string]int `validate:"nz"`
		ChanField   chan int       `validate:"nz"`
		TimeField   time.Time      `validate:"nz"`
	}

	testCases := []struct {
		name   string
		input  Data
		assert func(t *testing.T, err error)
	}{
		{
			name: "valid input",
			input: Data{
				StringField: "Hello",
				BoolField:   true,
				IntField:    42,
				UintField:   123,
				FloatField:  3.14,
				ArrayField:  [3]int{1, 2, 3},
				SliceField:  []string{"a", "b", "c"},
				MapField:    map[string]int{"x": 10, "y": 20},
				ChanField:   make(chan int),
				TimeField:   time.Now(),
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:  "invalid input",
			input: Data{},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Struct(tc.input)
			tc.assert(t, err)
		})
	}
}
