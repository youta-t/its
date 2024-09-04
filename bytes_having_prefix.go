package its

import (
	"bytes"

	"github.com/youta-t/its/itskit"
)

// BytesHavingPrefix tests with bytes.HasPrefix
func BytesHavingPrefix(want []byte) Matcher[[]byte] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.HasPrefix(got, want)
		},
		`bytes.HasPrefix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}
