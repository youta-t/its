package its

import (
	"strings"

	"github.com/youta-t/its/itskit"
)

// StringHavingSuffix tests with strings.HasSuffix
func StringHavingSuffix(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.HasSuffix((string)(got), want)
		},
		`strings.HasSuffix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
