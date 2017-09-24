package walk_test

import (
	"errors"
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
		err   error
	}{
		{title: "pointer to string", in: pString("foobar"), out: pString("foobarfoobar")},
		{title: "string", in: "foobar", out: "foobar", err: errors.New("Couldn't set the value - need pointer or slice as argument")},
		{
			title: "slice of strings",
			in:    []string{"foobar", "baarfoo"},
			out:   []string{"foobarfoobar", "baarfoobaarfoo"},
		},
		{
			title: "struct with strings",
			in:    &twoString{A: "foo", B: "bar", c: "small"},
			out:   &twoString{A: "foofoo", B: "barbar", c: "small"},
		},
		{
			title: "struct with strings and ints",
			in:    &twoStringInt{A: "foo", B: "bar", I: 42},
			out:   &twoStringInt{A: "foofoo", B: "barbar", I: 42},
		},
		{
			title: "struct with a ptr to a string",
			in:    &ptrStringStruct{A: "foo", B: pString("bar")},
			out:   &ptrStringStruct{A: "foofoo", B: pString("barbar")},
		},
		{
			title: "struct with structs",
			in:    &deepStrings{A: "foo", D: twoString{A: "bar"}, P: &deepStrings{A: "baz"}},
			out:   &deepStrings{A: "foofoo", D: twoString{A: "barbar"}, P: &deepStrings{A: "bazbaz"}},
		},
		{
			title: "non pointer struct",
			in:    twoString{A: "bar"},
			out:   twoString{A: "bar"},
			err:   errors.New("Couldn't set the value - need pointer or slice as argument"),
		},
		{
			title: "string nil pointer",
			in:    (*string)(nil),
			out:   (*string)(nil),
		},
		{
			title: "struct with a nil pointer",
			in:    &deepStrings{A: "foo", D: twoString{A: "bar"}, P: nil},
			out:   &deepStrings{A: "foofoo", D: twoString{A: "barbar"}, P: nil},
		},
		// TODO map with strings
		// TODO map with structs
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := walk.Strings(tt.in, func(s string) string {
				return s + s
			})
			if (err == nil) != (tt.err == nil) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
				t.Fatalf("Unexpected error: %s - expected %s", err, tt.err)
			}
			if !cmp.Equal(tt.in, tt.out, cmp.AllowUnexported(twoString{})) {
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

type twoString struct {
	A string
	B string
	c string
}

type twoStringInt struct {
	A string
	B string
	I int
}

type ptrStringStruct struct {
	A string
	B *string
}

type deepStrings struct {
	A string
	D twoString
	P *deepStrings
}
