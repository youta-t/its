package its

import (
	"fmt"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type ptrMatcher[T any] struct {
	label itskit.Label
	m     Matcher[T]
}

func (p ptrMatcher[T]) Match(got *T) itskit.Match {
	if got == nil {
		return itskit.NG(
			p.label.Fill("nil"),
			itskit.NG(p.m.String()),
		)
	}
	match := p.m.Match(*got)

	return itskit.NewMatch(
		match.Ok(), p.label.Fill(fmt.Sprintf("%T", got)), match,
	)
}

func (p ptrMatcher[T]) Write(w itsio.Writer) error {
	if err := p.label.Write(w); err != nil {
		return err
	}
	in := w.Indent()
	if err := p.m.Write(in); err != nil {
		return err
	}
	return nil
}

func (p ptrMatcher[T]) String() string {
	return itskit.MatcherToString(p)
}

// Pointer wraps matcher as matcher to pointer.
//
// It checks got value is not nil and matches with m.
//
// To check got is nil, use Nil[T]().
//
// # Args
//
// - m Matcher[T]: matcher for dereferenced value.
func Pointer[T any](m Matcher[T]) Matcher[*T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return ptrMatcher[T]{
		label: itskit.NewLabelWithLocation(
			"%s is not nil,",
			itskit.Got,
		),
		m: m,
	}
}
