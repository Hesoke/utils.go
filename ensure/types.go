package ensure

import (
	"reflect"
	"strings"
)

// Returns true if obj is not nil.
//
// IMPORTANT: if obj is int, float, complex, array, struct or string,
// it will return true, because this types are not nillable
func NotNil[T any](obj T) bool {
	if !isNillable(obj) {
		return true
	}
	v := reflect.ValueOf(obj)
	kind := v.Kind()
	return (kind == reflect.Func ||
		kind == reflect.Chan ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Interface ||
		kind == reflect.Pointer) && !v.IsNil()
}

// Returns true if obj is not empty.
//
// IMPORTANT: if obj is int, float, complex, struct,
// it will return true, because this types do not have length
func NotEmpty[T any](obj T) bool {
	if !canEmpty(obj) {
		return true
	}
	v := reflect.ValueOf(obj)
	kind := v.Kind()
	return (kind == reflect.Chan ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.String ||
		kind == reflect.Array ||
		kind == reflect.Struct) &&
		func() bool {
			if kind == reflect.String {
				return strings.TrimSpace(any(obj).(string)) != ""
			}
			if kind == reflect.Struct {
				return !reflect.DeepEqual(obj, reflect.Zero(reflect.TypeOf(obj)).Interface())
			}
			return v.Len() != 0
		}()
}

func NotNilOrEmpty[T any](obj T) bool {
	return NotNil(obj) && NotEmpty(obj)
}

func isNillable(obj any) bool {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return false
	case reflect.Float32, reflect.Float64:
		return false
	case reflect.Complex64, reflect.Complex128:
		return false
	case reflect.String, reflect.Array, reflect.Struct, reflect.Bool:
		return false
	default:
		return true
	}
}

func canEmpty(obj any) bool {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice, reflect.Map, reflect.String, reflect.Struct:
		return true
	default:
		return false
	}
}
