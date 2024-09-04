package its

import (
	"bytes"

	"github.com/youta-t/its/itskit"
)

// BytesContaining tests with bytes.Contains
func BytesContaining(want []byte) Matcher[[]byte] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.Contains(got, want)
		},
		`bytes.Contains(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
