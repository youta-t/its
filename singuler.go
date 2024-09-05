package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// Singuler tests that new got value DO NOT match with any former got values.
//
// Singuler is stateful matcher.
// Do not use this in slice related matcher.
//
// # Example
//
//	// match when all got strings are different
//	itsUniqueId := its.Singuler(its.EqEq[string])
//	itsUniqueId.Match("id=a").OrError(t)
//	itsUniqueId.Match("id=b").OrError(t)  // pass
//	itsUniqueId.Match("id=c").OrError(t)  // pass
//	itsUniqueId.Match("id=b").OrError(t)  // fail!
//
// # Args
//
// - matcherFactory: factory function creates a new mathcer for new got.
func Singuler[T any](matcherFactory func(T) Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return &singulerMatcher[T]{
		label:          itskit.NewLabelWithLocation("//do not match with values have been gotten"),
		matcherFactory: matcherFactory,
	}
}

type singulerMatcher[T any] struct {
	label          itskit.Label
	matcherFactory func(T) Matcher[T]
	matchers       []Matcher[T]
	matches        []itskit.Match
}

func (uniq *singulerMatcher[T]) Match(got T) itskit.Match {
	cancel := itskit.SkipStack()
	defer cancel()

	var match itskit.Match
	if len(uniq.matchers) == 0 {
		match = Always[T]().Match(got)
	} else {
		match = None(uniq.matchers...).Match(got)
	}
	uniq.matches = append(uniq.matches, match)

	uniq.matchers = append(uniq.matchers, uniq.matcherFactory(got))

	return itskit.NewMatch(
		match.Ok(),
		uniq.label.Fill(itskit.Missing),
		uniq.matches...,
	)
}

func (uniq *singulerMatcher[T]) Write(w itsio.Writer) error {
	return itsio.WriteBlock(w, uniq.label.String(), uniq.matches)
}

func (uniq *singulerMatcher[T]) String() string {
	return itskit.MatcherToString[T](uniq)
}
