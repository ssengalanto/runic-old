package fn

import "reflect"

func Deref[T any](ptr *T) T {
	field := reflect.ValueOf(ptr)

	// If the pointer is nil, return zeroed value
	if field.IsNil() {
		return reflect.Zero(field.Type().Elem()).Interface().(T)
	}

	// Dereference the pointer and return the value
	return field.Elem().Interface().(T)
}
