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

type namedMatcher[T any] struct {
	name Label
	m    Matcher[T]
}

func Named[T any, M Matcher[T], L string | Label](
	name L,
	m M,
) Matcher[T] {
	var label Label
	switch n := (any(name)).(type) {
	case string:
		if !strings.HasSuffix(n, ":") {
			n += " :"
		}
		label = NewLabel(n)
	case Label:
		label = n
	}
	return namedMatcher[T]{
		name: label,
		m:    m,
	}
}

func (n namedMatcher[T]) Match(got T) Match {
	m := n.m.Match(got)
	return NewMatch(
		m.Ok(), n.name.Fill(got), m,
	)
}

func (k namedMatcher[T]) Write(w itsio.Writer) error {
	if err := k.name.Write(w); err != nil {
		return err
	}
	return k.m.Write(w.Indent())
}

func (k namedMatcher[T]) String() string {
	return MatcherToString[T](k)
}
