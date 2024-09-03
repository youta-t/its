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
	// ✘ // all: (2 ok / 3 matchers)		--- @ ./all_test.go:18
	//     ✔ strings.HasPrefix(/* got */ "abc...fghIJkl...xyz", /* want */ "abc")		--- @ ./all_test.go:19
	//     ✘ strings.Contains(/* got */ "abc...fghIJkl...xyz", /* want */ "hij")		--- @ ./all_test.go:20
	//     ✔ strings.HasSuffix(/* got */ "abc...fghIJkl...xyz", /* want */ "xyz")		--- @ ./all_test.go:21
}
