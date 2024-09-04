package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// All tests actual passes all specs.
//
// If no matchers are given, it always pass.
func All[T any](matchers ...Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return allMatcher[T]{
		header: itskit.NewLabelWithLocation(
			"// all: (%d ok / %d matchers)",
			itskit.Placeholder, len(matchers),
		),
		specs: matchers,
	}
}

type allMatcher[T any] struct {
	header itskit.Label
	specs  []Matcher[T]
}

func (as allMatcher[T]) Match(actual T) itskit.Match {
	matches := make([]itskit.Match, len(as.specs))
	nSpec := len(as.specs)
	nOk := 0
	for nth, s := range as.specs {
		m := s.Match(actual)
		matches[nth] = m
		if m.Ok() {
			nOk += 1
		}
	}
	return itskit.NewMatch(
		nOk == nSpec,
		as.header.Fill(actual, nOk),
		matches...,
	)
}

func (as allMatcher[T]) Write(ww itsio.Writer) error {
	if err := ww.WriteStringln(as.header.String()); err != nil {
		return err
	}
	in := ww.Indent()
	for _, m := range as.specs {
		if err := m.Write(in); err != nil {
			return err
		}
	}
	return nil
}

func (as allMatcher[T]) String() string {
	return itskit.MatcherToString[T](as)
}
