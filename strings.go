package walk

import (
	"errors"
	"reflect"
)

// Strings walks over all strings of the first parameter and executes the given function on it
func Strings(i interface{}, fn func(string) string) error {
	rv := reflect.ValueOf(i)
	if !rv.Elem().CanSet() {
		return errors.New("Couldn't set the value of the parameter - need a pointer or a slice")
	}
	rv.Elem().SetString(fn(rv.Elem().String()))
	return nil
}
