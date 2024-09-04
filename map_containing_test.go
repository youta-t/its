package its_test

import "github.com/youta-t/its"

func ExampleMapContaining_ok() {
	its.MapContaining(map[string]its.Matcher[int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		// "c": its.EqEq(99),
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)

	// same as above
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
		// "c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)
	// Output:
}

func ExampleMapContaining_different_value() {
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 99,
		}).
		OrError(t)

	// Output:
	// ✘ map[string]int{ ... (contain; keys /* got */ 2, /* want */ 2; -1)		--- @ ./map_containing_test.go:34
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_containing_test.go:34
	//     ✘ b:
	//         ✘ /* got */ 99 == /* want */ 98		--- @ ./map_containing_test.go:34
}

func ExampleMapContaining_different_key() {
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
	}).
		Match(map[string]int{
			"a": 97,
			"x": 98,
			"c": 99,
		}).
		OrError(t)
	// Output:
	// ✘ map[string]int{ ... (contain; keys /* got */ 3, /* want */ 2; -1)		--- @ ./map_containing_test.go:53
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_containing_test.go:53
	//     ✘ b: (not in got)
	//         ✘ /* got */ ?? == /* want */ 98		--- @ ./map_containing_test.go:53
	//     ✘ c: (not in want)
	//         ✘ /* got */ 99, /* want */ ??
	//     ✘ x: (not in want)
	//         ✘ /* got */ 98, /* want */ ??
}
