package testexample

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := map[string]struct{ A, B, Expected int }{
		"foo": {1, 1, 2},
		"bar": {1, -1, 0},
	}
	for k, tc := range cases {
		actual := tc.A + tc.B
		if actual != tc.Expected{
			t.Errorf(
				"%s: %d + %d = %d, expected %d",
				k, tc.A, tc.B, actual, tc.Expected)
		}
	}
}

func TestAdd1(t *testing.T) {
	cases := map[string]struct{ A, B, Expected int }{
		"foo": {1, 1, 2},
		"bar": {1, -1, 0},
	}
	for k, tc := range cases {
		t.Run("",func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected{
				t.Errorf(
					"%s: %d + %d = %d, expected %d",
					k, tc.A, tc.B, actual, tc.Expected)
			}
		})
	}
}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hell o
}
