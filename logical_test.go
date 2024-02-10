package its_test

import "github.com/youta-t/its"

func ExampleAll() {
	its.All(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).Match("abc...fghijkl...xyz").OrError(t)

	its.All(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).Match("abc...fghIJkl...xyz").OrError(t)

	// Output:
	// ✘ // all: (2 ok / 3 matchers)		--- @ ./logical_test.go:12
	//     ✔ strings.HasPrefix(/* got */ "abc...fghIJkl...xyz", /* want */ "abc")		--- @ ./logical_test.go:13
	//     ✘ strings.Contains(/* got */ "abc...fghIJkl...xyz", /* want */ "hij")		--- @ ./logical_test.go:14
	//     ✔ strings.HasSuffix(/* got */ "abc...fghIJkl...xyz", /* want */ "xyz")		--- @ ./logical_test.go:15
}

func ExampleSome() {
	its.Some(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).Match("abc...fghIJkl...xyz").OrError(t)

	its.Some(
		its.StringHavingPrefix("abc"),
		its.StringContaining("hij"),
		its.StringHavingSuffix("xyz"),
	).Match("The quick brown fox jumps over the lazy dog").OrError(t)
	// Output:
	// ✘ // some: (0 ok / 3 matchers)		--- @ ./logical_test.go:32
	//     ✘ strings.HasPrefix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "abc")		--- @ ./logical_test.go:33
	//     ✘ strings.Contains(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "hij")		--- @ ./logical_test.go:34
	//     ✘ strings.HasSuffix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "xyz")		--- @ ./logical_test.go:35
}

func ExampleNot() {
	its.Not(its.EqEq(42)).Match(24).OrError(t)
	its.Not(its.EqEq(42)).Match(35).OrError(t)
	its.Not(its.EqEq(42)).Match(42).OrError(t)
	// Output:
	// ✘ // not:		--- @ ./logical_test.go:47
	//     ✔ /* got */ 42 == /* want */ 42		--- @ ./logical_test.go:47
}

func ExampleNone() {
	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(4).OrError(t)

	its.None(
		its.EqEq(1),
		its.EqEq(2),
		its.EqEq(3),
	).
		Match(2).OrError(t)

	// Output:
	// ✘ // none of:		--- @ ./logical_test.go:61
	//     ✘ /* got */ 2 == /* want */ 1		--- @ ./logical_test.go:62
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./logical_test.go:63
	//     ✘ /* got */ 2 == /* want */ 3		--- @ ./logical_test.go:64
}
