package its_test

import "github.com/youta-t/its"

func ExampleNot_ok() {
	its.Not(its.EqEq(42)).Match(24).OrError(t)
	its.Not(its.EqEq(42)).Match(35).OrError(t)
	// Output:
}

func ExampleNot_ng() {
	its.Not(its.EqEq(42)).Match(42).OrError(t)
	// Output:
	// ✘ // not:		--- @ ./not_test.go:12
	//     ✔ /* got */ 42 == /* want */ 42		--- @ ./not_test.go:12
}
