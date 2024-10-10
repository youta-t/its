package its

import (
	"context"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type contextWithValueMatcher[T any] struct {
	label    itskit.Label
	nilLabel string
	key      any
	base     Matcher[T]
}

func (m *contextWithValueMatcher[T]) Match(ctx context.Context) itskit.Match {
	if ctx == nil {
		return itskit.NG(m.nilLabel)
	}

	got := ctx.Value(m.key)
	switch gott := got.(type) {
	case T:
		match := m.base.Match(gott)
		return itskit.NewMatch(
			match.Ok(),
			m.label.Fill(gott),
			match,
		)
	case any: // type mismatch
		return itskit.NG(
			m.label.Fill(got),
			itskit.NG(
				itskit.NewLabel("%+v is not %T", itskit.Got, *new(T)).Fill(got),
			),
		)
	default: // untyped!
		if mm, ok := m.base.(Matcher[any]); ok {
			match := mm.Match(gott)
			return itskit.NewMatch(
				match.Ok(),
				m.label.Fill(gott),
				match,
			)
		}
		return itskit.NG(
			m.label.Fill(got),
			itskit.NG(
				itskit.NewLabel("%+v is not %T", itskit.Got, *new(T)).Fill(got),
			),
		)
	}

}

func (m *contextWithValueMatcher[T]) Write(ww itsio.Writer) error {
	if err := m.label.Write(ww); err != nil {
		return err
	}
	in := ww.Indent()
	if err := m.base.Write(in); err != nil {
		return err
	}
	return nil
}

func (m *contextWithValueMatcher[T]) String() string {
	return itskit.MatcherToString(m)
}

// ContextWithValue tests with
//
//	value.Match(ctx.Value(key))
//
// If the value is not type T, the match will fail.
//
// # Args
//
// - key: key associated with Value
//
// - value: the matcher for Value typed T
func ContextWithValue[T any](key any, value Matcher[T]) Matcher[context.Context] {
	cancel := itskit.SkipStack()
	defer cancel()

	return &contextWithValueMatcher[T]{
		label: itskit.NewLabelWithLocation(
			"// got = ctx.Value(%+v)",
			key,
		),
		nilLabel: itskit.NewLabelWithLocation("// given context.Context is nil").Fill(nil),
		key:      key,
		base:     value,
	}
}
