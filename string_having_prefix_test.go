package its_test

import "github.com/youta-t/its"

func ExampleStringHavingPrefix_ok() {
	its.StringHavingPrefix("abc").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringHavingPrefix_ng() {
	its.StringHavingPrefix("abc").Match("adcbe").OrError(t)

	// Output:
	// ✘ strings.HasPrefix(/* got */ "adcbe", /* want */ "abc")		--- @ ./string_having_prefix_test.go:11
}
