package its_test

import "github.com/youta-t/its"

func ExampleAll_ok() {
	its.All(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).
		Match("abc...fghijkl...xyz").
		OrError(t)
	// Output:
}

func ExampleAll_ng() {

	its.All(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).
		Match("abc...fghIJkl...xyz").
		OrError(t)

	// Output:
	// ✘ // all: (2 ok / 3 matchers)		--- @ ./logical_test.go:18
	//     ✔ strings.HasPrefix(/* got */ "abc...fghIJkl...xyz", /* want */ "abc")		--- @ ./logical_test.go:19
	//     ✘ strings.Contains(/* got */ "abc...fghIJkl...xyz", /* want */ "hij")		--- @ ./logical_test.go:20
	//     ✔ strings.HasSuffix(/* got */ "abc...fghIJkl...xyz", /* want */ "xyz")		--- @ ./logical_test.go:21
}

func ExampleSome_ok() {
	its.Some(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).
		Match("abc...fghIJkl...xyz").
		OrError(t)
	// Output:
}

func ExampleSome_ng() {
	its.Some(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).
		Match("The quick brown fox jumps over the lazy dog").
		OrError(t)
	// Output:
	// ✘ // some: (0 ok / 3 matchers)		--- @ ./logical_test.go:45
	//     ✘ strings.HasPrefix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "abc")		--- @ ./logical_test.go:46
	//     ✘ strings.Contains(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "hij")		--- @ ./logical_test.go:47
	//     ✘ strings.HasSuffix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "xyz")		--- @ ./logical_test.go:48
}

func ExampleNot_ok() {
	its.Not(its.EqEq(42)).Match(24).OrError(t)
	its.Not(its.EqEq(42)).Match(35).OrError(t)
	// Output:
}

func ExampleNot_ng() {
	its.Not(its.EqEq(42)).Match(42).OrError(t)
	// Output:
	// ✘ // not:		--- @ ./logical_test.go:66
	//     ✔ /* got */ 42 == /* want */ 42		--- @ ./logical_test.go:66
}

func ExampleNone_ok() {
	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(4).
		OrError(t)
	// Output:
}

func ExampleNone_ng() {
	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(2).
		OrError(t)

	// Output:
	// ✘ // none of:		--- @ ./logical_test.go:84
	//     ✘ /* got */ 2 == /* want */ 1		--- @ ./logical_test.go:85
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./logical_test.go:86
	//     ✘ /* got */ 2 == /* want */ 3		--- @ ./logical_test.go:87
}
