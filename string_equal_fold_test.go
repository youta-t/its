package its_test

import "github.com/youta-t/its"

func ExampleStringEqualFold_ok() {
	its.StringEqualFold("abc").Match("abc").OrError(t)
	its.StringEqualFold("aBc").Match("AbC").OrError(t)
	// Output:
}

func ExampleStringEqualFold_ng() {
	its.StringEqualFold("abc").Match("αβγ").OrError(t)
	// Output:
	// ✘ strings.EqualFold(/* got */ "αβγ", /* want */ "abc")		--- @ ./string_equal_fold_test.go:12
}
