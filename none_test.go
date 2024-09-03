package its_test

import "github.com/youta-t/its"

func ExampleNone_ok() {
	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(4).
		OrError(t)
	// Output:
}

func ExampleNone_ng() {
	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(2).
		OrError(t)

	// Output:
	// ✘ // none of:		--- @ ./none_test.go:17
	//     ✘ /* got */ 2 == /* want */ 1		--- @ ./none_test.go:18
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./none_test.go:19
	//     ✘ /* got */ 2 == /* want */ 3		--- @ ./none_test.go:20
}
