### What is walk?

Walk provides an easy and type safe way to manipulate all fields of the same
type of a variable in Go (Golang). Look at the *How to use walk*.

### Installing

The easiest way to get walk is go getting it
```
go get github.com/fvosberg/walk
```

### How to use walk

#### Walk over strings

To manipulate all strings in a given variable you can use the following code
```
u := User{
	FirstName: "*** Random ***",
	LastName: "*** Guy ***",
	Email: "guy@random.com",
}
walk.Strings(u, func(s string) string {
	return strings.Trim(s, " *")
})
fmt.Printf("After processing: %#v")
// User{FirstName:"Random", LastName:"Guy", Email:"guy@random.com"}
```
