package its_test

import "github.com/youta-t/its"

func ExampleStringHavingSuffix_ok() {
	its.StringHavingSuffix("cde").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringHavingSuffix_ng() {
	its.StringHavingSuffix("cde").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.HasSuffix(/* got */ "adcbe", /* want */ "cde")		--- @ ./string_having_suffix_test.go:11
}
