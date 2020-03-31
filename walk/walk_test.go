package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	// We create a slice of structs...
	cases := []struct {
		// ... that have three values in them: the name of the struct, the
		// interface that contains the input, and the things we expect to get
		// out of the function call.
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		// And here they are.
		{
			// Item 1.
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		// And that's it.
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
