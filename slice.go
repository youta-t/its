package its

import (
	"strings"

	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
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

	diffs := editorialgraph.NewWithMatcher(actual, m)
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
