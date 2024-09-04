package its

import (
	"bytes"

	"github.com/youta-t/its/itskit"
)

// BytesHavingSuffix tests with bytes.HasSuffix
func BytesHavingSuffix(want []byte) Matcher[[]byte] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.HasSuffix(got, want)
		},
		`bytes.HasSuffix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
