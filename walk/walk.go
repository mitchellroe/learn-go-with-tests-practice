package walk

// golang challenge: write a function walk(x interface{}, fn func(string)) which
// takes a struct x and calls fn for all strings fields found inside. difficulty
// level: recursively.

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// So get the value of the interface we were passed.
	val := reflect.ValueOf(x)
	// Then store the first field in that interface into variable "field"
	field := val.Field(0)
	// Then pass that field to the function we were passed.
	fn(field.String())
}
