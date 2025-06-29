package its_test

import (
	"math/big"

	"github.com/youta-t/its"
)

func ExampleCmpLesserThan_ok() {
	want := big.NewInt(42)
	got := big.NewInt(41)
	its.CmpLesserThan(want).Match(got).OrError(t)
	// Output:
}

func ExampleCmpLesserThan_ng() {
	want := big.NewInt(42)
	got := big.NewInt(43)
	its.CmpLesserThan(want).Match(got).OrError(t)
	// Output:
	// ✘ (/* got */ 43).Cmp(/* want */ 42) < 0		--- @ ./cmp_lesser_than_test.go:19
}
