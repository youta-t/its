package its_test

import "github.com/youta-t/its"

func ExampleAll() {
	its.All(
		its.StringHavingPrefix[string]("abc"),
		its.StringContaining[string]("hij"),
		its.StringHavingSuffix[string]("xyz"),
	).Match("abc...fghijkl...xyz").OrError(t)

	its.All(
		its.StringHavingPrefix[string]("abc"),
		its.StringContaining[string]("hij"),
		its.StringHavingSuffix[string]("xyz"),
	).Match("abc...fghIJkl...xyz").OrError(t)

	// Output:
	// ✘ // all: (2 ok / 3 matchers)
	//     ✔ strings.HasPrefix(/* got */ "abc...fghIJkl...xyz", /* want */ "abc")
	//     ✘ strings.Contains(/* got */ "abc...fghIJkl...xyz", /* want */ "hij")
	//     ✔ strings.HasSuffix(/* got */ "abc...fghIJkl...xyz", /* want */ "xyz")
}

func ExampleSome() {
	its.Some(
		its.StringHavingPrefix[string]("abc"),
		its.StringContaining[string]("hij"),
		its.StringHavingSuffix[string]("xyz"),
	).Match("abc...fghIJkl...xyz").OrError(t)

	its.Some(
		its.StringHavingPrefix[string]("abc"),
		its.StringContaining[string]("hij"),
		its.StringHavingSuffix[string]("xyz"),
	).Match("The quick brown fox jumps over the lazy dog").OrError(t)
	// Output:
	// ✘ // some: (0 ok / 3 matchers)
	//     ✘ strings.HasPrefix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "abc")
	//     ✘ strings.Contains(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "hij")
	//     ✘ strings.HasSuffix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "xyz")
}

func ExampleNot() {
	its.Not(its.EqEq(42)).Match(24).OrError(t)
	its.Not(its.EqEq(42)).Match(35).OrError(t)
	its.Not(its.EqEq(42)).Match(42).OrError(t)
	// Output:
	// ✘ // not:
	//     ✔ /* got */ 42 == /* want */ 42
}
