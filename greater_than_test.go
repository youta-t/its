package its_test

import "github.com/youta-t/its"

func ExampleGreaterThan_int() {
	its.GreaterThan(10).Match(10).OrError(t)
	its.GreaterThan(11).Match(10).OrError(t)
	its.GreaterThan(10).Match(11).OrError(t)
	// Output:
	// ✘ /* want */ 10 < /* got */ 10		--- @ ./greater_than_test.go:6
	//
	// ✘ /* want */ 11 < /* got */ 10		--- @ ./greater_than_test.go:7
}

func ExampleGreaterThan_float() {
	its.GreaterThan(1.0).Match(1.0).OrError(t)
	its.GreaterThan(1.1).Match(1.0).OrError(t)
	its.GreaterThan(1.0).Match(1.1).OrError(t)
	// Output:
	// ✘ /* want */ 1 < /* got */ 1		--- @ ./greater_than_test.go:16
	//
	// ✘ /* want */ 1.1 < /* got */ 1		--- @ ./greater_than_test.go:17
}

func ExampleGreaterThan_string() {
	its.GreaterThan("aaa").Match("aaa").OrError(t)
	its.GreaterThan("aab").Match("aaa").OrError(t)
	its.GreaterThan("aaa").Match("aab").OrError(t)

	// Output:
	// ✘ /* want */ aaa < /* got */ aaa		--- @ ./greater_than_test.go:26
	//
	// ✘ /* want */ aab < /* got */ aaa		--- @ ./greater_than_test.go:27
}
