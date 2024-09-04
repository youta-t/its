package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// None tests got value does NOT match for all given matchers.
func None[T any](matchers ...Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return noneMathcer[T]{
		label:   itskit.NewLabelWithLocation("// none of:"),
		matcher: matchers,
	}
}

type noneMathcer[T any] struct {
	label   itskit.Label
	matcher []Matcher[T]
}

func (none noneMathcer[T]) Match(got T) itskit.Match {
	ok := false
	matches := make([]itskit.Match, len(none.matcher))

	for nth, m := range none.matcher {
		match := m.Match(got)
		ok = ok || match.Ok()
		matches[nth] = match
	}

	return itskit.NewMatch(
		!ok,
		none.label.Fill(got),
		matches...,
	)
}

func (none noneMathcer[T]) Write(w itsio.Writer) error {
	return itsio.WriteBlock(w, none.label.String(), none.matcher)
}

func (none noneMathcer[T]) String() string {
	return itskit.MatcherToString[T](none)
}
