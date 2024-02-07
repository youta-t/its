package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// Monotonic tests that the new got value matches with the last got value.
//
// Monotonis is stateful matcher.
// Do not use this in slice related matcher.
//
// # Example
//
// itsNewEvent := its.Monotonic(its.After[time.Time])
// t1 := time.Now()
// itsNewEvent.Match(t1).OrError(t)
// t2 := time.Now()
// t3 := time.Now()
// itsNewEvent.Match(t2).OrError(t) // pass
// t4 := time.Now()
// itsNewEvent.Match(t4).OrError(t) // pass
// itsNewEvent.Match(t3).OrError(t) // fail!
//
// # Args
//
// - matcherFactory: factory function that creates new Matcher with new got.
func Monotonic[T any](matcherFactory func(T) Matcher[T]) itskit.Matcher[T] {
	return &monotonic[T]{
		label:          itskit.NewLabel("// monotonic"),
		matcherFactory: matcherFactory,
		nextMatcher:    Always[T](),
	}
}

type monotonic[T any] struct {
	label          itskit.Label
	matcherFactory func(T) Matcher[T]
	nextMatcher    Matcher[T]
	mathces        []itskit.Match
}

func (mono *monotonic[T]) Match(got T) itskit.Match {
	match := mono.nextMatcher.Match(got)
	mono.mathces = append(mono.mathces, match)
	mono.nextMatcher = mono.matcherFactory(got)

	return itskit.NewMatch(
		match.Ok(),
		mono.label.Fill(got),
		mono.mathces...,
	)
}

func (mono *monotonic[T]) Write(w itsio.Writer) error {
	return itsio.WriteBlock(w, mono.label.String(), mono.mathces)
}

func (mono *monotonic[T]) String() string {
	return itskit.MatcherToString[T](mono)
}

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
func Singuler[T any](matcherFactory func(T) Matcher[T]) itskit.Matcher[T] {
	return &singulerMatcher[T]{
		label:          itskit.NewLabel("//do not match with values have been gotten"),
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
