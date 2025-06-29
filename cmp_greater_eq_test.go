package its_test

import (
	"math/big"

	"github.com/youta-t/its"
)

func ExampleCmpGreaterEq_ok_equal() {
	want := big.NewInt(42)
	got := big.NewInt(42)
	its.CmpGreaterEq(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpGreaterEq_ok_greater() {
	want := big.NewInt(42)
	got := big.NewInt(43)
	its.CmpGreaterEq(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpGreaterEq_ng() {
	want := big.NewInt(42)
	got := big.NewInt(41)
	its.CmpGreaterEq(want).Match(got).OrError(t)
	// Output:
	// ✘ (/* got */ 41).Cmp(/* want */ 42) >= 0		--- @ ./cmp_greater_eq_test.go:26
}
