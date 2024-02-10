package its_test

import (
	"github.com/youta-t/its"
)

func ExampleMap() {
	its.Map(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).OrError(t)

	its.Map(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 99,
		"c": 99,
	}).OrError(t)

	its.Map(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 98,
		"d": 99,
	}).OrError(t)
	// Output:
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:18
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:19
	//     ✘ b:
	//         ✘ /* got */ 99 == /* want */ 98		--- @ ./map_test.go:20
	//     ✔ c:
	//         ✔ /* got */ 99 == /* want */ 99		--- @ ./map_test.go:21
	//
	// ✘ map[string]int{... ( keys: /* got */ 3, /* want */ 3; +1, -1 )		--- @ ./map_test.go:28
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:29
	//     ✔ b:
	//         ✔ /* got */ 98 == /* want */ 98		--- @ ./map_test.go:30
	//     ✘ c: (not in got)
	//         ✘ /* got */ ?? == /* want */ 99		--- @ ./map_test.go:31
	//     ✘ d: (not in want)
	//         ✘ /* got */ 99, /* want */ ??
}

func ExampleMapContaining() {

	its.MapContaining(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).OrError(t)

	its.MapContaining(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).OrError(t)

	its.MapContaining(its.MapSpec[string, int]{
		"a": its.EqEq(96),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}).OrError(t)

	its.MapContaining(its.MapSpec[string, int]{
		"a": its.EqEq(97),
		"b": its.EqEq(98),
		"c": its.EqEq(99),
	}).Match(map[string]int{
		"a": 97,
		"c": 99,
	}).OrError(t)

	// Output:
	// ✘ map[string]int{ ... (contain; keys /* got */ 3, /* want */ 2; -1)		--- @ ./map_test.go:78
	//     ✘ a:
	//         ✘ /* got */ 97 == /* want */ 96		--- @ ./map_test.go:79
	//     ✘ b: (not in want)
	//         ✘ /* got */ 98, /* want */ ??
	//     ✔ c:
	//         ✔ /* got */ 99 == /* want */ 99		--- @ ./map_test.go:80
	//
	// ✘ map[string]int{ ... (contain; keys /* got */ 2, /* want */ 3; -1)		--- @ ./map_test.go:87
	//     ✔ a:
	//         ✔ /* got */ 97 == /* want */ 97		--- @ ./map_test.go:88
	//     ✘ b: (not in got)
	//         ✘ /* got */ ?? == /* want */ 98		--- @ ./map_test.go:89
	//     ✔ c:
	//         ✔ /* got */ 99 == /* want */ 99		--- @ ./map_test.go:90
}
