//nolint:gochecknoglobals,gochecknoinits // unnecessary rules for this package
package validator

import (
	"fmt"
	"reflect"
	"strings"

	v "github.com/go-playground/validator/v10"
	"github.com/ssengalanto/runic/pkg/exceptions"
)

var validator *v.Validate

type fieldErr struct {
	key   string
	tag   string
	value any
}

func init() {
	validator = v.New()
	registerCustomValidators()
}

// Var validates a single field using a specified tag.
// It returns an error if the validation fails.
func Var(field any, tag string) error {
	err := validator.Var(field, tag)
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.ErrInvalid, varErrMsg(err, field))
	}

	return nil
}

// Struct validates a struct using the struct field tags.
// It returns an error if any of the fields fail validation.
func Struct(s any) error {
	err := validator.Struct(s)
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.ErrInvalid, structErrMsg(err))
	}

	return nil
}

// varErrMsg generates an error message for a single field validation failure.
func varErrMsg(err error, value any) string {
	var errMsg strings.Builder
	fieldErrs := fieldErrors(err)
	valueType := reflect.TypeOf(value)

	idx := 0
	for _, field := range fieldErrs {
		msg := fmt.Sprintf("`%s` field with value of `%v` failed on `%s` validation", valueType.Name(), value, field.tag)
		if idx > 0 {
			errMsg.WriteString(", ")
		}
		errMsg.WriteString(msg)
		idx++
	}

	return errMsg.String()
}

// structErrMsg generates an error message for struct validation failure.
func structErrMsg(err error) string {
	var errMsg strings.Builder
	fieldErrs := fieldErrors(err)

	idx := 0
	for _, field := range fieldErrs {
		msg := fmt.Sprintf(
			"`%s` field with value of `%v` failed on `%s` validation",
			field.key,
			field.value,
			field.tag,
		)
		if idx > 0 {
			errMsg.WriteString(", ")
		}
		errMsg.WriteString(msg)
		idx++
	}

	return errMsg.String()
}

// fieldErrors extracts field-level validation errors from the validation error.
func fieldErrors(err error) map[string]fieldErr {
	fieldErrs := make(map[string]fieldErr)

	for _, err := range err.(v.ValidationErrors) { //nolint:errorlint //intentional type assertion
		if _, ok := fieldErrs[err.StructNamespace()]; !ok {
			fieldErrs[err.StructNamespace()] = fieldErr{
				key:   err.StructNamespace(),
				tag:   err.Tag(),
				value: err.Value(),
			}
		}
	}

	return fieldErrs
}
