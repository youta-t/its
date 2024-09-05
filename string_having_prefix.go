package its

import (
	"strings"

	"github.com/youta-t/its/itskit"
)

// StringHavingPrefix tests with strings.HasPrefix
func StringHavingPrefix(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.HasPrefix((string)(got), want)
		},
		`strings.HasPrefix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
