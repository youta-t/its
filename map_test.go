package its_test

import (
	"github.com/youta-t/its"
)

func ExampleMap_ok() {
	its.Map(map[string]its.Matcher[int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)

	// same as above
	its.ForEntries(its.Map, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)
	// Output:
}

func ExampleMap_different_value() {
	its.ForEntries(its.Map, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 99,
			"c": 99,
		}).
		OrError(t)

	// Output:
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:36
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:36
	//     ✘ b:
	//         ✘ /* got */ 99 == /* want */ 98		--- @ ./map_test.go:36
	//     ✔ c:
	//         ✔ /* got */ 99 == /* want */ 99		--- @ ./map_test.go:36
}

func ExampleMap_different_key() {
	its.ForEntries(its.Map, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"d": 99,
		}).
		OrError(t)
	// Output:
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:59
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:59
	//     ✔ b:
	//         ✔ /* got */ 98 == /* want */ 98		--- @ ./map_test.go:59
	//     ✘ c: (not in got)
	//         ✘ /* got */ ?? == /* want */ 99		--- @ ./map_test.go:59
	//     ✘ d: (not in want)
	//         ✘ /* got */ 99, /* want */ ??
}

func ExampleMapContaining_ok() {

	its.MapContaining(map[string]its.Matcher[int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)

	// same above
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)

	// less entries is ok.
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 97,
		"c": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)
	// Output:
}

func ExampleMapContaining_different_entries() {
	its.ForEntries(its.MapContaining, its.EqEq, map[string]int{
		"a": 96,
		"b": 98,
		"d": 99,
	}).
		Match(map[string]int{
			"a": 97,
			"b": 98,
			"c": 99,
		}).
		OrError(t)
	// Output:
	// ✘ map[string]int{ ... (contain; keys /* got */ 3, /* want */ 3; -2)		--- @ ./map_test.go:124
	//     ✘ a:
	//         ✘ /* got */ 97 == /* want */ 96		--- @ ./map_test.go:124
	//     ✔ b:
	//         ✔ /* got */ 98 == /* want */ 98		--- @ ./map_test.go:124
	//     ✘ c: (not in want)
	//         ✘ /* got */ 99, /* want */ ??
	//     ✘ d: (not in got)
	//         ✘ /* got */ ?? == /* want */ 99		--- @ ./map_test.go:124
}
