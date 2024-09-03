package its

import (
	"math"

	"github.com/youta-t/its/itskit"
)

// NaN tests with math.IsNaN
func NaN() Matcher[float64] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		math.IsNaN,
		"math.IsNaN(%f)",
		itskit.Got,
	)
}
