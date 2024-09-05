package its

import (
	"bytes"

	"github.com/youta-t/its/itskit"
)

// BytesEqual tests with bytes.Equal
func BytesEqual(want []byte) Matcher[[]byte] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.Equal(got, want)
		},
		`bytes.Equal(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
