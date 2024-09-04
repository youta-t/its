package its

import "github.com/youta-t/its/itskit"

// EquivWith tests with
//
//	equiv(want, got)
//
// # Args
//
// - want T: expectation
//
// - equiv: function returns true if want matches with got.
func EquivWith[T, U any](want T, equiv func(want T, got U) bool) Matcher[U] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got U) bool { return equiv(want, got) },
		"(%+v) equiv. (%+v)",
		itskit.Want(want), itskit.Got,
	)
}
