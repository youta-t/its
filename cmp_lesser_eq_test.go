package its_test

import (
	"math/big"

	"github.com/youta-t/its"
)

func ExampleCmpLesserEq_ok_equal() {
	want := big.NewInt(42)
	got := big.NewInt(42)
	its.CmpLesserEq(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpLesserEq_ok_lesser() {
	want := big.NewInt(42)
	got := big.NewInt(41)
	its.CmpLesserEq(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpLesserEq_ng() {
	want := big.NewInt(42)
	got := big.NewInt(43)
	its.CmpLesserEq(want).Match(got).OrError(t)
	// Output:
	// ✘ (/* got */ 43).Cmp(/* want */ 42) <= 0		--- @ ./cmp_lesser_eq_test.go:26
}
