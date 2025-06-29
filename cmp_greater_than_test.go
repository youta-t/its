package its_test

import (
	"math/big"

	"github.com/youta-t/its"
)

func ExampleCmpGreaterThan_ok() {
	want := big.NewInt(42)
	got := big.NewInt(43)
	its.CmpGreaterThan(want).Match(got).OrError(t)
	// Output:
}
func ExampleCmpGreaterThan_ng() {
	want := big.NewInt(42)
	got := big.NewInt(41)
	its.CmpGreaterThan(want).Match(got).OrError(t)
	// Output:
	// ✘ (/* got */ 41).Cmp(/* want */ 42) > 0		--- @ ./cmp_greater_than_test.go:18
}
