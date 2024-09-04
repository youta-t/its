package its

import "github.com/youta-t/its/itskit"

// After tests with
//
//	got.After(want)
//
// want value can be time.Time, for example, but whatever okay if it has After().
func After[T interface{ After(T) bool }](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return got.After(want) },
		"(%+v).After(%+v)",
		itskit.Got, itskit.Want(want),
	)
}
