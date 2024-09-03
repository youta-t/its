package its_test

import "github.com/youta-t/its"

func ExampleSinguler() {
	itsUniqueId := its.Singuler(its.EqEq[string])
	itsUniqueId.Match("id: a").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: c").OrError(t)
	itsUniqueId.Match("id: b").OrError(t)
	itsUniqueId.Match("id: e").OrError(t)
	// Output:
	// ✘ //do not match with values have been gotten		--- @ ./singuler_test.go:6
	//     ✔ (always pass)		--- @ ./singuler_test.go:7
	//     ✔ // none of:		--- @ ./singuler_test.go:8
	//         ~ /* got */ id: b == /* want */ id: a		--- @ ./singuler_test.go:7
	//     ✔ // none of:		--- @ ./singuler_test.go:9
	//         ~ /* got */ id: c == /* want */ id: a		--- @ ./singuler_test.go:7
	//         ~ /* got */ id: c == /* want */ id: b		--- @ ./singuler_test.go:8
	//     ✘ // none of:		--- @ ./singuler_test.go:10
	//         ✘ /* got */ id: b == /* want */ id: a		--- @ ./singuler_test.go:7
	//         ✔ /* got */ id: b == /* want */ id: b		--- @ ./singuler_test.go:8
	//         ✘ /* got */ id: b == /* want */ id: c		--- @ ./singuler_test.go:9
}
