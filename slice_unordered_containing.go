package its

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/internal/set"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type sliceUnorderedContainingMatcher[T any] sliceSpec[T]

// SliceContainsUnordered test that actual slice contain elements satisfy each specs.
//
// # Args
//
// - mathcers: matchers for each element.
// Each matchers is tried multiple times with many element.
// Do not use stateful matcher.
func SliceUnorderedContaining[T any](spec ...Matcher[T]) Matcher[[]T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return sliceUnorderedContainingMatcher[T]{
		matchers: spec,
		template: itskit.NewLabelWithLocation(
			"[]%T{ ... (unordered, contain; len: %d, %d; -%d)",
			*new(T),
			itskit.Got,
			itskit.Want(len(spec)),
			itskit.Placeholder,
		),
	}
}

func (ss sliceUnorderedContainingMatcher[T]) Match(actual []T) itskit.Match {
	ms := make([]itskit.Matcher[T], len(ss.matchers))
	for i, m := range ss.matchers {
		ms[i] = m
	}
	diffs := set.CompareWithMatcher(actual, ms)
	matches := []itskit.Match{}
	miss := 0
	for _, d := range diffs {
		matches = append(matches, d.Value)
		switch d.Mode {
		case diff.Missing:
			miss += 1
		}
	}

	return itskit.NewMatch(
		miss == 0,
		ss.template.Fill(len(actual), miss),
		matches...,
	)
}

func (ss sliceUnorderedContainingMatcher[T]) Write(ww itsio.Writer) error {
	if err := ss.template.Write(ww); err != nil {
		return err
	}
	iw := ww.Indent()
	for _, s := range ss.matchers {
		if err := s.Write(iw); err != nil {
			return err
		}
	}
	return nil
}

func (ss sliceUnorderedContainingMatcher[T]) String() string {
	return itskit.MatcherToString[[]T](ss)
}
