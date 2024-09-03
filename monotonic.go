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
func Monotonic[T any](matcherFactory func(T) Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return &monotonic[T]{
		label:          itskit.NewLabelWithLocation("// monotonic"),
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
	cancel := itskit.SkipStack()
	defer cancel()
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
