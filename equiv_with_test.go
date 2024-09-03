package its_test

import (
	"fmt"

	"github.com/youta-t/its"
)

func ExampleEquivWith_ok() {
	its.EquivWith(
		42,
		func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).
		Match("42").
		OrError(t)
	// Output:
}

func ExampleEquivWith_ng() {
	its.EquivWith(
		42,
		func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).
		Match("40").
		OrError(t)

	// Output:
	// ✘ (/* want */ 42) equiv. (/* got */ 40)		--- @ ./equiv_with_test.go:20
}
