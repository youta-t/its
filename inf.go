package its

import (
	"math"

	"github.com/youta-t/its/itskit"
)

// Inf tests with math.IsInf
//
// This matcher will pass either positive or negative infinity.
func Inf() Matcher[float64] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got float64) bool {
			return math.IsInf(got, 0)
		},
		"math.IsInf(%f, 0)",
		itskit.Got,
	)
}
