package its

import "github.com/youta-t/its/itskit"

// Match matches with Match(T)bool method.
//
// # Example
//
//	MatchString(regexp.MustCompile(`[Mm]atcher`))
func MatchString(m interface{ MatchString(string) bool }) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		m.MatchString,
		"(%+v).MatchString(%#v)",
		itskit.Want(m), itskit.Got,
	)
}
