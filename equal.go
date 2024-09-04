package its

import "github.com/youta-t/its/itskit"

// Equal tests with
//
//	expcted.Equal(got)
//
// want value can be time.Time, for example, but whatever okay if it has Equal().
func Equal[T any, E interface{ Equal(T) bool }](want E) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		want.Equal,
		"(%+v).Equal(%+v)",
		itskit.Want(want), itskit.Got,
	)
}
