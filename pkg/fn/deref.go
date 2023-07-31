package fn

import "reflect"

func Deref(ptr any) any {
	// Ensure the input is a pointer
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("deref: input must be a pointer")
	}

	field := reflect.ValueOf(ptr)

	// If the pointer is nil, return zeroed value
	if field.IsNil() {
		return reflect.Zero(field.Type().Elem()).Interface()
	}

	// Dereference the pointer and return the value
	return field.Elem().Interface()
}
