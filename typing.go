package its

import "github.com/youta-t/its/itskit"

// Matcher does match got an want values,
// determines success or not and writes error message.
//
// To run matching, use Matcher.Match.
type Matcher[T any] itskit.Matcher[T]

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Numeric interface {
	Integer | Float
}
