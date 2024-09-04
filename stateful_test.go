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
	// ✘ // monotonic		--- @ ./stateful_test.go:6
	//     ✔ (always pass)		--- @ ./stateful_test.go:6
	//     ✔ /* want */ apple < /* got */ banana		--- @ ./stateful_test.go:8
	//     ✔ /* want */ banana < /* got */ cherry		--- @ ./stateful_test.go:9
	//     ✘ /* want */ cherry < /* got */ bacon		--- @ ./stateful_test.go:10
}

func ExampleSinguler() {
	itsUniqueId := its.Singuler(its.EqEq[string])
	itsUniqueId.Match("id: a").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: c").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: e").OrError(t)
	// Output:
	// ✘ //do not match with values have been gotten		--- @ ./stateful_test.go:22
	//     ✔ (always pass)		--- @ ./stateful_test.go:23
	//     ✔ // none of:		--- @ ./stateful_test.go:24
	//         ~ /* got */ id: b == /* want */ id: a		--- @ ./stateful_test.go:23
	//     ✔ // none of:		--- @ ./stateful_test.go:25
	//         ~ /* got */ id: c == /* want */ id: a		--- @ ./stateful_test.go:23
	//         ~ /* got */ id: c == /* want */ id: b		--- @ ./stateful_test.go:24
	//     ✘ // none of:		--- @ ./stateful_test.go:26
	//         ✘ /* got */ id: b == /* want */ id: a		--- @ ./stateful_test.go:23
	//         ✔ /* got */ id: b == /* want */ id: b		--- @ ./stateful_test.go:24
	//         ✘ /* got */ id: b == /* want */ id: c		--- @ ./stateful_test.go:25
}
