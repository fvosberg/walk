package walk_test

import (
	"reflect"
	"testing"

	"github.com/fvosberg/walk"
	"github.com/google/go-cmp/cmp"
)

func TestWalkStrings(t *testing.T) {
	tests := []struct {
		title string
		in    interface{}
		out   interface{}
	}{
		{in: pString("foobar"), out: pString("foobarfoobar")},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := walk.Strings(tt.in, func(s string) string {
				return s + s
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if !cmp.Equal(tt.in, tt.out) {
				var in interface{} = tt.in
				var out interface{} = tt.out
				orv := reflect.ValueOf(tt.out)
				if orv.Kind() == reflect.Ptr {
					out = reflect.Indirect(orv.Elem())
				}
				irv := reflect.ValueOf(tt.in)
				if irv.Kind() == reflect.Ptr {
					in = reflect.Indirect(irv.Elem())
				}
				t.Errorf("The result is expected to be %#v, but was %#v", out, in)
			}
		})
	}
}

func pString(s string) *string {
	return &s
}
