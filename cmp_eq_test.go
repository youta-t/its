package its_test

import (
	"math/big"

	"github.com/youta-t/its"
)

func ExampleCmpEq_ok() {
	want := big.NewInt(42)
	got := big.NewInt(42)
	its.CmpEq(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpEq_ng() {
	want := big.NewInt(42)
	got := big.NewInt(43)
	its.CmpEq(want).Match(got).OrError(t)
	// Output:
	// ✘ (/* got */ 43).Cmp(/* want */ 42) == 0		--- @ ./cmp_eq_test.go:19
}
