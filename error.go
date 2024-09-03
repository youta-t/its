package its

import (
	"errors"

	"github.com/youta-t/its/itskit"
)

// Error tests with errors.Is .
func Error(want error) Matcher[error] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher[error](
		func(got error) bool {
			return errors.Is(got, want)
		},
		"errors.Is(%s, %s)",
		itskit.Got, itskit.Want(want),
	)
}
