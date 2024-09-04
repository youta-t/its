package its

import (
	"errors"

	"github.com/youta-t/its/itskit"
)

// ErrorAs tests with errors.As .
func ErrorAs[T error]() Matcher[error] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher[error](
		func(got error) bool {
			want := new(T)
			return errors.As(got, want)
		},
		"want := new(%T); errors.As(%s, want)",
		*new(T), itskit.Got,
	)
}
