package its_test

import "github.com/youta-t/its"

func ExampleMonotonic() {
	itsDictinoalyOrder := its.Monotonic(its.GreaterThan[string])

	itsDictinoalyOrder.Match("apple").OrError(t)
	itsDictinoalyOrder.Match("banana").OrError(t)
	itsDictinoalyOrder.Match("cherry").OrError(t)
	itsDictinoalyOrder.Match("bacon").OrError(t)
	itsDictinoalyOrder.Match("castard").OrError(t)
	// Output:
	// ✘ // monotonic		--- @ ./monotonic_test.go:6
	//     ✔ (always pass)		--- @ ./monotonic_test.go:6
	//     ✔ /* want */ apple < /* got */ banana		--- @ ./monotonic_test.go:8
	//     ✔ /* want */ banana < /* got */ cherry		--- @ ./monotonic_test.go:9
	//     ✘ /* want */ cherry < /* got */ bacon		--- @ ./monotonic_test.go:10
}
