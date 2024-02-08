package its

import (
	"github.com/youta-t/its/itskit"
)

// Matcher does match got an want values,
// determines success or not and writes error message.
//
// To run matching, use Matcher.Match.
type Matcher[T any] itskit.Matcher[T]
