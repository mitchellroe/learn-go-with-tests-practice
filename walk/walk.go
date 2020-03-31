package walk

// golang challenge: write a function walk(x interface{}, fn func(string)) which
// takes a struct x and calls fn for all strings fields found inside. difficulty
// level: recursively.

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// So get the value of the interface we were passed.
	val := reflect.ValueOf(x)

	// Then loop through each of the fields that are in it
	for i := 0; i < val.NumField(); i++ {
		// Grab the i-th
		field := val.Field(i)
		fn(field.String())
	}
}
