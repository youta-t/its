package its_test

import "github.com/youta-t/its"

func ExampleStringContaining_ok() {
	its.StringContaining("bcd").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringContaining_ng() {
	its.StringContaining("bcd").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.Contains(/* got */ "adcbe", /* want */ "bcd")		--- @ ./string_containing_test.go:11
}
