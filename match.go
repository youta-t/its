package its

import "github.com/youta-t/its/itskit"

// Match matches with Match(T)bool method.
//
// # Example
//
//	Match[[]byte](regexp.MustCompile(`[Mm]atcher`))
func Match[T any, M interface{ Match(T) bool }](m M) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		m.Match,
		"(%+v).Match(%+v)",
		itskit.Want(m), itskit.Got,
	)
}
