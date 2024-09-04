package its_test

import (
	"math"

	"github.com/youta-t/its"
)

func ExampleNaN_ok() {
	its.NaN().Match(math.NaN()).OrError(t)
	// Output:
}

func ExampleNaN_ng() {
	its.NaN().Match(42).OrError(t)
	// Output:
	// ✘ math.IsNaN(/* got */ 42.000000)		--- @ ./nan_test.go:15
}
