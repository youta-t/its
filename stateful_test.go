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
	// ✘ // monotonic
	//     ✔ (always pass)
	//     ✔ /* want */ apple < /* got */ banana
	//     ✔ /* want */ banana < /* got */ cherry
	//     ✘ /* want */ cherry < /* got */ bacon
}

func ExampleSinguler() {
	itsUniqueId := its.Singuler(its.EqEq[string])
	itsUniqueId.Match("id: a").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: c").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: e").OrError(t)
	// Output:
	// ✘ //do not match with values have been gotten
	//     ✔ (always pass)
	//     ✔ // none of:
	//         ~ /* got */ id: b == /* want */ id: a
	//     ✔ // none of:
	//         ~ /* got */ id: c == /* want */ id: a
	//         ~ /* got */ id: c == /* want */ id: b
	//     ✘ // none of:
	//         ✘ /* got */ id: b == /* want */ id: a
	//         ✔ /* got */ id: b == /* want */ id: b
	//         ✘ /* got */ id: b == /* want */ id: c
}
