package its_test

import "github.com/youta-t/its"

func ExamplePointer_ok() {
	got := 42
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)
	// Output:
}

func ExamplePointer_ng() {
	got := 40
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)

	var ptrgot *int
	its.Pointer(its.EqEq(42)).Match(ptrgot).OrError(t)
	// Output:
	// ✘ /* got */ *int is not nil,		--- @ ./pointer_test.go:13
	//     ✘ /* got */ 40 == /* want */ 42		--- @ ./pointer_test.go:13
	//
	// ✘ /* got */ nil is not nil,		--- @ ./pointer_test.go:16
	//     ✘ /* got */ ?? == /* want */ 42		--- @ ./pointer_test.go:16
}
