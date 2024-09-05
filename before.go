package its

import "github.com/youta-t/its/itskit"

// Before tests with
//
//	got.Before(want)
//
// want value can be time.Time, for example, but whatever okay if it has Before().
func Before[T interface{ Before(T) bool }](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return got.Before(want) },
		"(%+v).Before(%+v)",
		itskit.Got, itskit.Want(want),
	)
}
