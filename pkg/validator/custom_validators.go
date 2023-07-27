package validator

import (
	"log"
	"reflect"
	"time"

	v "github.com/go-playground/validator/v10"
)

func registerCustomValidators() {
	err := validator.RegisterValidation("nz", nonZeroed)
	if err != nil {
		log.Fatal(err)
	}
}

// nonZeroed is a custom validation function for non-zero values of all primitive types.
func nonZeroed(fl v.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	switch kind { //nolint:exhaustive //unnecessary
	case reflect.String:
		return field.String() != ""
	case reflect.Bool:
		return field.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return field.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return field.Float() != 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return field.Len() != 0
	case reflect.Chan:
		return !field.IsNil()
	case reflect.Struct:
		// handling time.Time
		if field.Type() == reflect.TypeOf(time.Time{}) {
			return !field.Interface().(time.Time).IsZero()
		}
	}

	return false
}
