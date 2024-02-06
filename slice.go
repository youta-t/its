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
	spec     []itskit.Matcher[T]
	template itskit.Label
}

func (eq sliceSpec[T]) Write(w itsio.Writer) error {
	if err := eq.template.Write(w); err != nil {
		return err
	}

	iw := w.Indent()
	for _, s := range eq.spec {
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
func Slice[T any](specs ...itskit.Matcher[T]) itskit.Matcher[[]T] {
	return sliceMatcher[T]{
		sliceSpec: sliceSpec[T]{
			spec: specs,
			template: itskit.NewLabel(
				"[]%T{ ... (len: %d, %d; +%d, -%d)",
				*new(T),
				itskit.Got,
				itskit.Want(len(specs)),
				itskit.Placeholder,
				itskit.Placeholder,
			),
		},
	}
}

func (eq sliceMatcher[T]) Match(actual []T) itskit.Match {
	diffs := editorialgraph.New(eq.spec, actual)
	nMiss := 0
	nExtra := 0
	submatches := []itskit.Match{}

	for _, d := range diffs {
		submatches = append(submatches, d.Match)
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
func SliceUnordered[T any](specs ...itskit.Matcher[T]) itskit.Matcher[[]T] {
	return sliceUnorderedMatcher[T]{
		spec: specs,
		template: itskit.NewLabel(
			"[]%T{ ... (unordered; len: %d, %d; +%d, -%d)",
			*new(T),
			itskit.Want(len(specs)),
			itskit.Got,
			itskit.Placeholder,
			itskit.Placeholder,
		),
	}
}

func (ss sliceUnorderedMatcher[T]) Match(actual []T) itskit.Match {
	diffs := set.Compare(actual, ss.spec)
	matches := []itskit.Match{}
	extra := 0
	miss := 0
	for _, d := range diffs {
		matches = append(matches, d.Match)
		switch d.Mode {
		case diff.Missing:
			miss += 1
		case diff.Extra:
			extra += 1
		}
	}

	return itskit.NewMatch(
		extra == 0 && miss == 0,
		ss.template.Fill(len(actual), extra, miss),
		matches...,
	)
}

func (ss sliceUnorderedMatcher[T]) Write(ww itsio.Writer) error {
	if err := ss.template.Write(ww); err != nil {
		return err
	}
	iw := ww.Indent()
	for _, s := range ss.spec {
		if err := s.Write(iw); err != nil {
			return err
		}
	}
	return nil
}

func (ss sliceUnorderedMatcher[T]) String() string {
	return itskit.MatcherToString[[]T](ss)
}

type sliceContainedUnorderedMatcher[T any] sliceSpec[T]

// SliceContainsUnordered test that actual slice contain elements satisfy each specs.
func SliceUnorderedContaining[T any](spec ...itskit.Matcher[T]) itskit.Matcher[[]T] {
	return sliceContainedUnorderedMatcher[T]{
		spec: spec,
		template: itskit.NewLabel(
			"[]%T{ ... (unordered, contain; len: %d, %d; -%d)",
			*new(T),
			itskit.Got,
			itskit.Want(len(spec)),
			itskit.Placeholder,
		),
	}
}

func (ss sliceContainedUnorderedMatcher[T]) Match(actual []T) itskit.Match {
	diffs := set.Compare(actual, ss.spec)
	matches := []itskit.Match{}
	miss := 0
	for _, d := range diffs {
		matches = append(matches, d.Match)
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

func (ss sliceContainedUnorderedMatcher[T]) Write(ww itsio.Writer) error {
	if err := ss.template.Write(ww); err != nil {
		return err
	}
	iw := ww.Indent()
	for _, s := range ss.spec {
		if err := s.Write(iw); err != nil {
			return err
		}
	}
	return nil
}

func (ss sliceContainedUnorderedMatcher[T]) String() string {
	return itskit.MatcherToString[[]T](ss)
}
