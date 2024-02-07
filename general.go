package its

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// EqEq tests of comparable with
//
//	want == got
func EqEq[T comparable](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return got == want },
		"%+v == %+v",
		itskit.Got, itskit.Want(want),
	)
}

type eqeqPtrMatcher[T comparable] struct {
	label itskit.Label
	want  *T
}

func ptrLabel[T any](v *T) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("&(%+v)", *v)
}

func (epm eqeqPtrMatcher[T]) Match(got *T) itskit.Match {
	ok := false

	if got == nil || epm.want == nil {
		ok = got == nil && epm.want == nil
	} else {
		ok = *got == *epm.want
	}

	return itskit.NewMatch(
		ok,
		epm.label.Fill(ptrLabel(got)),
	)
}

func (epm eqeqPtrMatcher[T]) Write(w itsio.Writer) error {
	return epm.label.Write(w)
}

func (epm eqeqPtrMatcher[T]) String() string {
	return itskit.MatcherToString[*T](epm)
}

// EqEqPtr tests of pointer for comparable with
//
//	(want == got) || (*want == *got)
func EqEqPtr[T comparable](want *T) Matcher[*T] {
	return eqeqPtrMatcher[T]{
		label: itskit.NewLabel(
			"%+v == %+v",
			itskit.Got, itskit.Want(ptrLabel(want)),
		),
		want: want,
	}
}

// GreaterThan tests of numeric value with
//
//	want < got
func GreaterThan[T Numeric | string](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return want < got },
		"%+v < %+v",
		itskit.Want(want), itskit.Got,
	)
}

// GreaterEq tests of numeric value with
//
//	want <= got
func GreaterEq[T Numeric | ~string](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return want <= got },
		"%+v <= %+v",
		itskit.Want(want), itskit.Got,
	)
}

// LesserThan tests of numeric value with
//
//	want > got
func LesserThan[T Numeric | ~string](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return want > got },
		"%+v > %+v",
		itskit.Want(want), itskit.Got,
	)
}

// LesserEq tests of numeric value with
//
//	want > got
func LesserEq[T Numeric | ~string](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return want >= got },
		"%+v >= %+v",
		itskit.Want(want), itskit.Got,
	)
}

// Before tests with
//
//	got.Before(want)
//
// want value can be time.Time, for example, but whatever okay if it has Before().
func Before[T interface{ Before(T) bool }](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return got.Before(want) },
		"(%+v).Before(%+v)",
		itskit.Got, itskit.Want(want),
	)
}

// After tests with
//
//	got.After(want)
//
// want value can be time.Time, for example, but whatever okay if it has After().
func After[T interface{ After(T) bool }](want T) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool { return got.After(want) },
		"(%+v).After(%+v)",
		itskit.Got, itskit.Want(want),
	)
}

// Equal tests with
//
//	expcted.Equal(got)
//
// want value can be time.Time, for example, but whatever okay if it has Equal().
func Equal[T any](want interface{ Equal(T) bool }) Matcher[T] {
	return itskit.SimpleMatcher(
		want.Equal,
		"(%+v).Equal(%+v)",
		itskit.Want(want), itskit.Got,
	)
}

// EquivWith tests with
//
//	equiv(want, got)
//
// # Args
//
// - want T: expectation
//
// - equiv: function returns true if want matches with got.
func EquivWith[T, U any](want T, equiv func(want T, got U) bool) Matcher[U] {
	return itskit.SimpleMatcher(
		func(got U) bool { return equiv(want, got) },
		"(%+v) equiv. (%+v)",
		itskit.Want(want), itskit.Got,
	)
}

// Error tests with errors.Is .
func Error(want error) Matcher[error] {
	return itskit.SimpleMatcher[error](
		func(got error) bool {
			return errors.Is(got, want)
		},
		"errors.Is(%s, %s)",
		itskit.Got, itskit.Want(want),
	)
}

// ErrorAs tests with errors.As .
func ErrorAs[T error]() Matcher[error] {
	return itskit.SimpleMatcher[error](
		func(got error) bool {
			want := new(T)
			return errors.As(got, want)
		},
		"want := new(%T); errors.As(%s, want)",
		*new(T), itskit.Got,
	)
}

// always pass.
func Always[T any]() Matcher[T] {
	return itskit.SimpleMatcher(
		func(T) bool { return true },
		"(always pass)",
	)
}

// always fail.
func Never[T any]() Matcher[T] {
	return itskit.SimpleMatcher(
		func(T) bool { return false },
		"(never pass)",
	)
}

// StringHavingPrefix tests with strings.HasPrefix
func StringHavingPrefix[T ~string](want string) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool {
			return strings.HasPrefix((string)(got), want)
		},
		`strings.HasPrefix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// StringHavingSuffix tests with strings.HasSuffix
func StringHavingSuffix[T ~string](want string) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool {
			return strings.HasSuffix((string)(got), want)
		},
		`strings.HasSuffix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// StringContaining tests with strings.Contains
func StringContaining[T ~string](want string) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool {
			return strings.Contains((string)(got), want)
		},
		`strings.Contains(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// StringEqualFold tests with strings.EqualFold
func StringEqualFold[T ~string](want string) Matcher[T] {
	return itskit.SimpleMatcher(
		func(got T) bool {
			return strings.EqualFold((string)(got), want)
		},
		`strings.EqualFold(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// BytesEqual tests with bytes.Equal
func BytesEqual(want []byte) Matcher[[]byte] {
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.Equal(got, want)
		},
		`bytes.Equal(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// BytesHavingPrefix tests with bytes.HasPrefix
func BytesHavingPrefix(want []byte) Matcher[[]byte] {
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.HasPrefix(got, want)
		},
		`bytes.HasPrefix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// BytesHavingSuffix tests with bytes.HasSuffix
func BytesHavingSuffix(want []byte) Matcher[[]byte] {
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.HasSuffix(got, want)
		},
		`bytes.HasSuffix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// BytesContaining tests with bytes.Contains
func BytesContaining(want []byte) Matcher[[]byte] {
	return itskit.SimpleMatcher(
		func(got []byte) bool {
			return bytes.Contains(got, want)
		},
		`bytes.Contains(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// NaN tests with math.IsNaN
func NaN() Matcher[float64] {
	return itskit.SimpleMatcher(
		math.IsNaN,
		"math.IsNaN(%f)",
		itskit.Got,
	)
}

// Inf tests with math.IsInf
//
// This matcher will pass either positive or negative infinity.
func Inf() Matcher[float64] {
	return itskit.SimpleMatcher(
		func(got float64) bool {
			return math.IsInf(got, 0)
		},
		"math.IsInf(%f, 0)",
		itskit.Got,
	)
}

type chanMatcher[T any] struct {
	label itskit.Label
}

// Closed tests wheather channel is closed or not.
//
// This matcher tries to receive from channel, it may cause sideeffect.
func Closed[T any]() Matcher[<-chan T] {
	return chanMatcher[T]{
		label: itskit.NewLabel("chan %T is %s.", *new(T), itskit.Placeholder),
	}
}

func (c chanMatcher[T]) Match(ch <-chan T) itskit.Match {
	var closed bool
	select {
	case _, ok := <-ch:
		closed = !ok
	default:
		closed = false
	}

	message := "not closed"
	if closed {
		message = "closed"
	}
	return itskit.NewMatch(
		closed,
		c.label.Fill(itskit.Missing, message),
	)
}

func (c chanMatcher[T]) Write(ww itsio.Writer) error {
	return c.label.Write(ww)
}

func (c chanMatcher[T]) String() string {
	return itskit.MatcherToString[<-chan T](c)
}

// Type tests got value is a type.
func Type[T any]() Matcher[any] {
	return itskit.SimpleMatcher(
		func(got any) bool {
			_, ok := got.(T)
			return ok
		},
		"%+v is a %T",
		itskit.Got, *new(T),
	)
}

// Match matches with Match(T)bool method.
//
// # Example
//
//	Match[[]byte](regexp.MustCompile(`[Mm]atcher`))
func Match[T any](m interface{ Match(T) bool }) Matcher[T] {
	return itskit.SimpleMatcher(
		m.Match,
		"(%+v).Match(%+v)",
		itskit.Want(m), itskit.Got,
	)
}

// Match matches with Match(T)bool method.
//
// # Example
//
//	MatchString(regexp.MustCompile(`[Mm]atcher`))
func MatchString(m interface{ MatchString(string) bool }) Matcher[string] {
	return itskit.SimpleMatcher(
		m.MatchString,
		"(%+v).MatchString(%+v)",
		itskit.Want(m), itskit.Got,
	)
}
