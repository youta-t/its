package its_test

import "github.com/youta-t/its"

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
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:34
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:34
	//     ✘ b:
	//         ✘ /* got */ 99 == /* want */ 98		--- @ ./map_test.go:34
	//     ✔ c:
	//         ✔ /* got */ 99 == /* want */ 99		--- @ ./map_test.go:34
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
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:57
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:57
	//     ✔ b:
	//         ✔ /* got */ 98 == /* want */ 98		--- @ ./map_test.go:57
	//     ✘ c: (not in got)
	//         ✘ /* got */ ?? == /* want */ 99		--- @ ./map_test.go:57
	//     ✘ d: (not in want)
	//         ✘ /* got */ 99, /* want */ ??
}
