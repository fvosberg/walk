package walk_test

import (
	"fmt"
	"strings"

	"github.com/fvosberg/walk"
)

func ExampleWalkStrings() {
	type user struct {
		FirstName string
		LastName  string
		Email     string
	}
	u := user{
		FirstName: "*** Random ***",
		LastName:  "*** Guy ***",
		Email:     "guy@random.com",
	}
	walk.Strings(&u, func(s string) string {
		return strings.Trim(s, " *")
	})
	fmt.Printf("%#v\n", u)
	// Output: walk_test.user{FirstName:"Random", LastName:"Guy", Email:"guy@random.com"}
}
