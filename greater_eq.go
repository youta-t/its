package its

import "github.com/youta-t/its/itskit"

// GreaterEq tests of numeric value with
//
//	want <= got
func GreaterEq[T Numeric | ~string](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return want <= got },
		"%+v <= %+v",
		itskit.Want(want), itskit.Got,
	)
}
