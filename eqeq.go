package its

import "github.com/youta-t/its/itskit"

// EqEq tests of comparable with
//
//	want == got
func EqEq[T comparable](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return got == want },
		"%+v == %+v",
		itskit.Got, itskit.Want(want),
	)
}
