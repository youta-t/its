package its

import (
	"strings"

	"github.com/youta-t/its/itskit"
)

// StringContaining tests with strings.Contains
func StringContaining(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.Contains((string)(got), want)
		},
		`strings.Contains(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
