package walk

import (
	"reflect"

	"github.com/pkg/errors"
)

// Strings walks over all strings of the first parameter and executes the given function on it
func Strings(i interface{}, fn onStrings) error {
	rv := reflect.ValueOf(i)
	if rv.Kind() != reflect.Ptr && rv.Kind() != reflect.Slice {
		return errors.New("Couldn't set the value - need pointer or slice as argument")
	}
	return fn.walk(rv)
}

type onStrings func(string) string

func (fn onStrings) walk(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.String:
		fn.apply(rv)
	case reflect.Ptr:
		return fn.walk(rv.Elem())
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			iv := rv.Index(i)
			if iv.Kind() == reflect.String {
				fn.apply(iv)
			}
		}
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			err := fn.walk(rv.Field(i))
			if err != nil {
				return errors.Wrapf(err, "processing field %#v failed", rv)
			}
		}
	}
	return nil
}

func (fn onStrings) apply(rv reflect.Value) {
	if !rv.CanSet() {
		return
	}
	rv.SetString(fn(rv.String()))
}
