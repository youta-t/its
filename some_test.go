package its_test

import "github.com/youta-t/its"

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
	// ✘ // some: (0 ok / 3 matchers)		--- @ ./some_test.go:17
	//     ✘ strings.HasPrefix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "abc")		--- @ ./some_test.go:18
	//     ✘ strings.Contains(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "hij")		--- @ ./some_test.go:19
	//     ✘ strings.HasSuffix(/* got */ "The quick brown fox jumps over the lazy dog", /* want */ "xyz")		--- @ ./some_test.go:20
}
