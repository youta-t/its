package its

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// EqEq tests of comparable with
//
//	want == got
func EqEq[T comparable](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
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
//
// Deprecated: Use Pointer(EqEq(...)) .
func EqEqPtr[T comparable](want *T) Matcher[*T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return eqeqPtrMatcher[T]{
		label: itskit.NewLabelWithLocation(
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
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return got.After(want) },
		"(%+v).After(%+v)",
		itskit.Got, itskit.Want(want),
	)
}

// DeepEqual tests with
//
//	reflect.DeepEqual(want, got)
func DeepEqual[T any](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return reflect.DeepEqual(got, want) },
		"reflect.DeepEqual(%+v, %+v)",
		itskit.Got, itskit.Want(want),
	)
}

// Equal tests with
//
//	expcted.Equal(got)
//
// want value can be time.Time, for example, but whatever okay if it has Equal().
func Equal[T any, E interface{ Equal(T) bool }](want E) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got U) bool { return equiv(want, got) },
		"(%+v) equiv. (%+v)",
		itskit.Want(want), itskit.Got,
	)
}

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

// always pass.
func Always[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return itskit.SimpleMatcher(
		func(T) bool { return true },
		"(always pass)",
	)
}

// always fail.
func Never[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(T) bool { return false },
		"(never pass)",
	)
}

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

// StringHavingSuffix tests with strings.HasSuffix
func StringHavingSuffix(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.HasSuffix((string)(got), want)
		},
		`strings.HasSuffix(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// StringContaining tests with strings.Contains
func StringContaining(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.Contains((string)(got), want)
		},
		`strings.Contains(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

// StringEqualFold tests with strings.EqualFold
func StringEqualFold(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got string) bool {
			return strings.EqualFold((string)(got), want)
		},
		`strings.EqualFold(%#v, %#v)`,
		itskit.Got, itskit.Want(want),
	)
}

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

// NaN tests with math.IsNaN
func NaN() Matcher[float64] {
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got float64) bool {
			return math.IsInf(got, 0)
		},
		"math.IsInf(%f, 0)",
		itskit.Got,
	)
}

type chanMatcher[T any, C chan T | <-chan T] struct {
	label itskit.Label
}

// ClosedChan tests wheather channel is closed or not.
//
// This matcher tries to receive from channel, it may cause sideeffect.
func ClosedChan[C chan T | <-chan T, T any]() Matcher[C] {
	cancel := itskit.SkipStack()
	defer cancel()
	return chanMatcher[T, C]{
		label: itskit.NewLabelWithLocation("chan %T is %s.", *new(T), itskit.Placeholder),
	}
}

func (c chanMatcher[T, C]) Match(ch C) itskit.Match {
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

func (c chanMatcher[T, C]) Write(ww itsio.Writer) error {
	return c.label.Write(ww)
}

func (c chanMatcher[T, C]) String() string {
	return itskit.MatcherToString(c)
}

// Type tests got value is a type.
func Type[T any]() Matcher[any] {
	cancel := itskit.SkipStack()
	defer cancel()
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
func Match[T any, M interface{ Match(T) bool }](m M) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
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
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		m.MatchString,
		"(%+v).MatchString(%#v)",
		itskit.Want(m), itskit.Got,
	)
}

type nilMatcher[T any] struct {
	label itskit.Label
}

func (n nilMatcher[T]) Match(got T) itskit.Match {
	rg := reflect.ValueOf(got)
	typ := reflect.ValueOf(*new(T))
	switch k := typ.Kind(); k {
	case reflect.Pointer:
		isnil := rg.IsNil()
		if isnil {
			return itskit.OK(n.label.Fill("<nil>"))
		}
		return itskit.NG(n.label.Fill(rg.Elem().Interface()))
	case reflect.Chan, reflect.Func:
		return itskit.NewMatch(
			rg.IsNil(),
			n.label.Fill(fmt.Sprintf("%T", rg.Interface())),
		)
	case reflect.Interface, reflect.Map, reflect.Slice:
		return itskit.NewMatch(
			rg.IsNil(),
			n.label.Fill(rg.Interface()),
		)
	default:
		isnil := any(got) == nil
		if isnil {
			return itskit.OK(n.label.Fill("<nil>"))
		}
		return itskit.NG(n.label.Fill(rg.Interface()))
	}
}

func (n nilMatcher[T]) Write(w itsio.Writer) error {
	return n.label.Write(w)
}

func (n nilMatcher[T]) String() string {
	return itskit.MatcherToString(n)
}

func Nil[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return nilMatcher[T]{
		label: itskit.NewLabelWithLocation("(%+v) is nil", itskit.Got),
	}
}

type ptrMatcher[T any] struct {
	label itskit.Label
	m     Matcher[T]
}

func (p ptrMatcher[T]) Match(got *T) itskit.Match {
	if got == nil {
		return itskit.NG(
			p.label.Fill("nil"),
			itskit.NG(p.m.String()),
		)
	}
	match := p.m.Match(*got)

	return itskit.NewMatch(
		match.Ok(), p.label.Fill(fmt.Sprintf("%T", got)), match,
	)
}

func (p ptrMatcher[T]) Write(w itsio.Writer) error {
	if err := p.label.Write(w); err != nil {
		return err
	}
	in := w.Indent()
	if err := p.m.Write(in); err != nil {
		return err
	}
	return nil
}

func (p ptrMatcher[T]) String() string {
	return itskit.MatcherToString(p)
}

// Pointer wraps matcher as matcher to pointer.
//
// It checks got value is not nil and matches with m.
//
// To check got is nil, use Nil[T]().
//
// # Args
//
// - m Matcher[T]: matcher for dereferenced value.
func Pointer[T any](m Matcher[T]) Matcher[*T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return ptrMatcher[T]{
		label: itskit.NewLabelWithLocation(
			"%s is not nil,",
			itskit.Got,
		),
		m: m,
	}
}

type textMatcher struct {
	label itskit.Label
	want  string
}

func (tm textMatcher) Match(got string) itskit.Match {
	gs := strings.SplitAfter(got, "\n")
	ws := strings.SplitAfter(tm.want, "\n")

	diffs := editorialgraph.New(
		gs, ws,
		func(s1, s2 string) (string, bool) {
			return s1, s1 == s2
		},
		func(s string) string { return s },
		func(s string) string { return s },
	)

	message := new(strings.Builder)
	message.WriteString(tm.label.String())
	message.WriteString("\n")

	unmatch := 0
	for _, d := range diffs {
		header := "      | "
		switch d.Mode {
		case diff.Extra:
			unmatch += 1
			header = "    + | "
		case diff.Missing:
			unmatch += 1
			header = "    - | "
		default:
		}

		val := d.Value
		if len(val) == 0 {
			message.WriteString(header[:len(header)-1])
			message.WriteString("\n")
			continue
		} else if val == "\n" {
			message.WriteString(header[:len(header)-1])
		} else {
			message.WriteString(header)
		}

		message.WriteString(d.Value)
	}

	return itskit.NewMatch(
		unmatch == 0,
		message.String(),
	)
}

func (tm textMatcher) Write(w itsio.Writer) error {
	if err := tm.label.Write(w); err != nil {
		return err
	}
	in := w.Indent()
	if err := in.WriteStringln(tm.want); err != nil {
		return err
	}
	return nil
}

func (tm textMatcher) String() string {
	return itskit.MatcherToString(tm)
}

// Text returns a matcher for a long text.
//
// When it get unmatch, it shows diff of text.
func Text(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()

	return textMatcher{
		label: itskit.NewLabelWithLocation("(+ = got, - = want)"),
		want:  want,
	}
}

type propMatcher[T, U any] struct {
	description itskit.Label
	prop        func(T) U
	m           Matcher[U]
}

// Property creates a matcher for property U calcurated from type T.
//
// # Args
//
// - description: description of property.
// It can be string for a static message, or itskit.Label for a dinamic message.
//
// - prop: calcuration extracting U from T
//
// - m: matcher for U
func Property[T, U any, D string | itskit.Label](
	description D,
	prop func(T) U,
	m Matcher[U],
) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	var label itskit.Label
	switch d := any(description).(type) {
	case string:
		label = itskit.NewLabelWithLocation(d + " :")
	case itskit.Label:
		label = d
	}

	return propMatcher[T, U]{
		description: label, prop: prop, m: m,
	}
}

func (k propMatcher[T, U]) Match(actual T) itskit.Match {
	p := k.prop(actual)
	match := k.m.Match(p)
	return itskit.NewMatch(match.Ok(), k.description.Fill(actual), match)
}

func (k propMatcher[T, U]) Write(w itsio.Writer) error {
	if err := k.description.Write(w); err != nil {
		return err
	}
	return k.m.Write(w.Indent())
}

func (k propMatcher[T, U]) String() string {
	return itskit.MatcherToString(k)
}
