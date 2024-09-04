package its_test

import "github.com/youta-t/its"

func ExampleLesserThan_int() {
	//       >
	its.LesserThan(10).Match(11).OrError(t)
	its.LesserThan(11).Match(10).OrError(t)
	its.LesserThan(10).Match(10).OrError(t)
	// Output:
	// ✘ /* want */ 10 > /* got */ 11		--- @ ./lesser_than_test.go:7
	//
	// ✘ /* want */ 10 > /* got */ 10		--- @ ./lesser_than_test.go:9
}

func ExampleLesserThan_float() {
	its.LesserThan(1.0).Match(1.1).OrError(t)
	its.LesserThan(1.0).Match(1.0).OrError(t)
	its.LesserThan(1.1).Match(1.0).OrError(t)
	// Output:
	// ✘ /* want */ 1 > /* got */ 1.1		--- @ ./lesser_than_test.go:17
	//
	// ✘ /* want */ 1 > /* got */ 1		--- @ ./lesser_than_test.go:18
}

func ExampleLesserThan_string() {
	its.LesserThan("aaa").Match("aab").OrError(t)
	its.LesserThan("aaa").Match("aaa").OrError(t)
	its.LesserThan("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aaa > /* got */ aab		--- @ ./lesser_than_test.go:27
	//
	// ✘ /* want */ aaa > /* got */ aaa		--- @ ./lesser_than_test.go:28
}
