package its_test

import "github.com/youta-t/its"

func ExampleGreaterEq_int() {
	//      <=
	its.GreaterEq(10).Match(11).OrError(t)
	its.GreaterEq(11).Match(10).OrError(t)
	its.GreaterEq(10).Match(10).OrError(t)
	// Output:
	// ✘ /* want */ 11 <= /* got */ 10		--- @ ./greater_eq_test.go:8
}

func ExampleGreaterEq_float() {
	its.GreaterEq(1.0).Match(1.1).OrError(t)
	its.GreaterEq(1.0).Match(1.0).OrError(t)
	its.GreaterEq(1.1).Match(1.0).OrError(t)
	// Output:
	// ✘ /* want */ 1.1 <= /* got */ 1		--- @ ./greater_eq_test.go:17
}

func ExampleGreaterEq_sring() {
	its.GreaterEq("aaa").Match("aab").OrError(t)
	its.GreaterEq("aaa").Match("aaa").OrError(t)
	its.GreaterEq("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aab <= /* got */ aaa		--- @ ./greater_eq_test.go:25
}
