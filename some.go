package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// All tests actual passes at least one spec.
//
// If no matchers are given, it always fail.
func Some[T any](matchers ...Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return someMatcher[T]{
		header: itskit.NewLabelWithLocation(
			"// some: (%d ok / %d matchers)",
			itskit.Placeholder, len(matchers),
		),
		submatch: matchers,
	}
}

type someMatcher[T any] struct {
	header   itskit.Label
	submatch []Matcher[T]
}

func (ss someMatcher[T]) Match(actual T) itskit.Match {
	matches := make([]itskit.Match, len(ss.submatch))
	nOk := 0
	for nth, s := range ss.submatch {
		m := s.Match(actual)
		matches[nth] = m
		if m.Ok() {
			nOk += 1
		}
	}
	return itskit.NewMatch(
		0 < nOk,
		ss.header.Fill(actual, nOk),
		matches...,
	)
}

func (ss someMatcher[T]) Write(ww itsio.Writer) error {
	if err := ww.WriteStringln(ss.header.String()); err != nil {
		return err
	}
	in := ww.Indent()
	for _, s := range ss.submatch {
		if err := s.Write(in); err != nil {
			return err
		}
	}
	return nil
}

func (ss someMatcher[T]) String() string {
	return itskit.MatcherToString[T](ss)
}
