package its

import (
	"strings"

	"github.com/youta-t/its/itskit"
)

// StringEqualFold tests with strings.EqualFold
func StringEqualFold(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.EqualFold((string)(got), want)
		},
		`strings.EqualFold(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
