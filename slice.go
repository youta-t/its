package its

import (
	"strings"

	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
	"github.com/youta-t/its/internal/set"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type sliceSpec[T any] struct {
	matchers []Matcher[T]
	template itskit.Label
}

func (eq sliceSpec[T]) Write(w itsio.Writer) error {
	if err := eq.template.Write(w); err != nil {
		return err
	}

	iw := w.Indent()
	for _, s := range eq.matchers {
		if err := s.Write(iw); err != nil {
			return err
		}
	}

	return nil
}

func (eq sliceSpec[T]) String() string {
	w := new(strings.Builder)
	eq.Write(itsio.Wrap(w))
	return w.String()
}

type sliceMatcher[T any] struct{ sliceSpec[T] }

// Slice tests actual slice elements satisfies specs for same index.
//
// # Args
//
// - mathcers: matchers for each element.
// Each matchers is tried multiple times with many element.
// Do not use stateful matcher.
func Slice[T any](matchers ...Matcher[T]) Matcher[[]T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return sliceMatcher[T]{
		sliceSpec: sliceSpec[T]{
			matchers: matchers,
			template: itskit.NewLabelWithLocation(
				"[]%T{ ... (len: %d, %d; +%d, -%d)",
				*new(T),
				itskit.Got,
				itskit.Want(len(matchers)),
				itskit.Placeholder,
				itskit.Placeholder,
			),
		},
	}
}

func (eq sliceMatcher[T]) Match(actual []T) itskit.Match {

	m := make([]itskit.Matcher[T], len(eq.matchers))
	for i, s := range eq.matchers {
		m[i] = s
	}

	diffs := editorialgraph.NewWithMatcher(m, actual)
	nMiss := 0
	nExtra := 0
	submatches := []itskit.Match{}

	for _, d := range diffs {
		submatches = append(submatches, d.Value)
		switch d.Mode {
		case diff.Missing:
			nMiss += 1
		case diff.Extra:
			nExtra += 1
		}
	}
	return itskit.NewMatch(
		nMiss == 0 && nExtra == 0,
		eq.sliceSpec.template.Fill(len(actual), nExtra, nMiss),
		submatches...,
	)
}

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
