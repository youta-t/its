package its_test

import (
	"github.com/youta-t/its"
	"github.com/youta-t/its/itskit"
)

func ExampleProperty_ok() {
	got := "abcde"
	its.Property(
		itskit.NewLabel("len(%s) == 5", itskit.Got),
		func(got string) int { return len(got) },
		its.EqEq(5),
	).
		Match(got).
		OrError(t)
	// Output:
}

func ExampleProperty_ok_no_placeholder() {
	got := "abcde"
	its.Property(
		"len == 5",
		func(got string) int { return len(got) },
		its.EqEq(5),
	).
		Match(got).
		OrError(t)
	// Output:
}

func ExampleProperty_ng() {
	got := "abcde"
	its.Property(
		itskit.NewLabelWithLocation("len(%s) == 4", itskit.Got),
		func(got string) int { return len(got) },
		its.EqEq(4),
	).
		Match(got).
		OrError(t)
	// Output:
	// ✘ len(/* got */ abcde) == 4		--- @ ./property_test.go:35
	//     ✘ /* got */ 5 == /* want */ 4		--- @ ./property_test.go:37
}

func ExampleProperty_ng_no_placeholder() {
	got := "abcde"
	its.Property(
		"len == 4",
		func(got string) int { return len(got) },
		its.EqEq(4),
	).
		Match(got).
		OrError(t)
	// Output:
	// ✘ len == 4 :		--- @ ./property_test.go:48
	//     ✘ /* got */ 5 == /* want */ 4		--- @ ./property_test.go:51
}
