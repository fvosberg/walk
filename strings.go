package walk

import (
	"errors"
	"reflect"
)

// Strings walks over all strings of the first parameter and executes the given function on it
func Strings(i interface{}, fn func(string) string) error {
	rv := reflect.ValueOf(i)
	if rv.Kind() != reflect.Ptr && rv.Kind() != reflect.Slice {
		return errors.New("Couldn't set the value - need pointer or slice as argument")
	}
	return strings(rv, fn)
}

func strings(rv reflect.Value, fn func(string) string) error {
	switch rv.Kind() {
	case reflect.String:
		applyString(rv, fn)
	case reflect.Ptr:
		return strings(rv.Elem(), fn)
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			iv := rv.Index(i)
			if iv.Kind() == reflect.String {
				applyString(iv, fn)
			}
		}
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			if rv.Field(i).CanSet() {
				applyString(rv.Field(i), fn)
			}
		}
	}
	return nil
}

func applyString(rv reflect.Value, fn func(string) string) {
	rv.SetString(fn(rv.String()))
}
