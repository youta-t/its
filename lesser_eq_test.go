package its_test

import "github.com/youta-t/its"

func ExampleLesserEq_int() {
	//      >=
	its.LesserEq(10).Match(11).OrError(t)
	its.LesserEq(11).Match(10).OrError(t)
	its.LesserEq(10).Match(10).OrError(t)

	// Output:
	// ✘ /* want */ 10 >= /* got */ 11		--- @ ./lesser_eq_test.go:7
}

func ExampleLesserEq_float() {
	its.LesserEq(1.0).Match(1.1).OrError(t)
	its.LesserEq(1.0).Match(1.0).OrError(t)
	its.LesserEq(1.1).Match(1.0).OrError(t)

	// Output:
	//　✘ /* want */ 1 >= /* got */ 1.1		--- @ ./lesser_eq_test.go:16
}

func ExampleLesserEq_string() {
	its.LesserEq("aaa").Match("aab").OrError(t)
	its.LesserEq("aaa").Match("aaa").OrError(t)
	its.LesserEq("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aaa >= /* got */ aab		--- @ ./lesser_eq_test.go:25
}
