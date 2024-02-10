package itskit

import (
	"strings"

	"github.com/youta-t/its/itskit/itsio"
)

// Build string expression of Matcher
func MatcherToString[T any](m Matcher[T]) string {
	sb := new(strings.Builder)
	ww := itsio.Wrap(sb)
	m.Write(ww)
	s := sb.String()
	if last := len(s) - 1; s[last] == '\n' {
		s = s[:last]
	}
	return s
}

// Matcher tests wheather given value matches with critelia or not.
type Matcher[A any] interface {

	// Match test value and return result.
	Match(A) Match

	// String expression of this matcher.
	//
	// You can implement this method as
	//
	// 	func (m youtMatcher) String() string {
	// 		return itskit.MatcherToString(m)
	// 	}
	//
	String() string

	// Write writes its string expression into itsio.Writer
	Write(itsio.Writer) error
}

func SimpleMatcher[T any](
	predicator func(T) bool,
	msg string,
	params ...any,
) Matcher[T] {
	cancel := SkipStack()
	defer cancel()

	return &simpleMatcher[T]{
		fn:    predicator,
		label: NewLabelWithLocation(msg, params...),
	}
}

type simpleMatcher[T any] struct {
	fn    func(T) bool
	label Label
}

func (ss *simpleMatcher[T]) Match(actual T) Match {
	return NewMatch(ss.fn(actual), ss.label.Fill(actual))
}

func (ss *simpleMatcher[T]) Write(ww itsio.Writer) error {
	return ss.label.Write(ww)
}

func (ss *simpleMatcher[T]) String() string {
	return MatcherToString[T](ss)
}

type propMatcher[T, U any] struct {
	description string
	prop        func(T) U
	m           Matcher[U]
}

// Property creates a matcher for property U calcurated from type T.
//
// # Args
//
// - description: description of property
//
// - prop: calcuration extractin U from T
//
// - m: matcher for U
func Property[T, U any](
	description string,
	prop func(T) U,
	m Matcher[U],
) Matcher[T] {
	return propMatcher[T, U]{description: description, prop: prop, m: m}
}

func (k propMatcher[T, U]) Match(actual T) Match {
	p := k.prop(actual)
	match := k.m.Match(p)
	return NewMatch(match.Ok(), k.description+" :", match)
}

func (k propMatcher[T, U]) Write(w itsio.Writer) error {
	if err := w.WriteStringln(k.description + " :"); err != nil {
		return err
	}
	return k.m.Write(w.Indent())
}

func (k propMatcher[T, U]) String() string {
	return MatcherToString[T](k)
}
