package its

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/internal/set"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type sliceUnorderedMatcher[T any] sliceSpec[T]

// Set test that for each element in actual slice matches each spec.
//
// # Args
//
// - mathcers: matchers for each element.
// Each matchers is tried multiple times with many element.
// Do not use stateful matcher.
func SliceUnordered[T any](specs ...Matcher[T]) Matcher[[]T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return sliceUnorderedMatcher[T]{
		matchers: specs,
		template: itskit.NewLabelWithLocation(
			"[]%T{ ... (unordered; len: %d, %d; +%d, -%d)",
			*new(T),
			itskit.Want(len(specs)),
			itskit.Got,
			itskit.Placeholder,
			itskit.Placeholder,
		),
	}
}

func (ss sliceUnorderedMatcher[T]) Match(got []T) itskit.Match {
	ms := make([]itskit.Matcher[T], len(ss.matchers))
	for i, m := range ss.matchers {
		ms[i] = m
	}
	diffs := set.CompareWithMatcher(got, ms)
	matches := []itskit.Match{}
	extra := 0
	miss := 0
	for _, d := range diffs {
		matches = append(matches, d.Value)
		switch d.Mode {
		case diff.Missing:
			miss += 1
		case diff.Extra:
			extra += 1
		}
	}

	return itskit.NewMatch(
		extra == 0 && miss == 0,
		ss.template.Fill(len(got), extra, miss),
		matches...,
	)
}

func (ss sliceUnorderedMatcher[T]) Write(ww itsio.Writer) error {
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

func (ss sliceUnorderedMatcher[T]) String() string {
	return itskit.MatcherToString[[]T](ss)
}
