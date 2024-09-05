package its_test

import (
	"math"

	"github.com/youta-t/its"
)

func ExampleInf() {
	its.Inf().Match(math.Inf(1)).OrError(t)
	its.Inf().Match(math.Inf(-1)).OrError(t)
	// Output:
}

func ExampleInf_ng() {
	its.Inf().Match(0).OrError(t)
	// Output:
	// ✘ math.IsInf(/* got */ 0.000000, 0)		--- @ ./inf_test.go:16
}
