package its

import (
	"github.com/youta-t/its/itskit"
)

// CmpLesserThan tests with
//
//	got.Cmp(want) < 0
//
// want value can be big.Int, for example, but whatever okay if it has Cmp().
func CmpLesserThan[T interface{ Cmp(T) int }](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool {
			return got.Cmp(want) < 0
		},
		"(%v).Cmp(%v) < 0",
		itskit.Got, itskit.Want(want),
	)
}
