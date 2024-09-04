package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type propMatcher[T, U any] struct {
	description itskit.Label
	prop        func(T) U
	m           Matcher[U]
}

// Property creates a matcher for property U calcurated from type T.
//
// # Args
//
// - description: description of property.
// It can be string for a static message, or itskit.Label for a dinamic message.
//
// - prop: calcuration extracting U from T
//
// - m: matcher for U
func Property[T, U any, D string | itskit.Label](
	description D,
	prop func(T) U,
	m Matcher[U],
) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	var label itskit.Label
	switch d := any(description).(type) {
	case string:
		label = itskit.NewLabelWithLocation(d + " :")
	case itskit.Label:
		label = d
	}

	return propMatcher[T, U]{
		description: label, prop: prop, m: m,
	}
}

func (k propMatcher[T, U]) Match(actual T) itskit.Match {
	p := k.prop(actual)
	match := k.m.Match(p)
	return itskit.NewMatch(match.Ok(), k.description.Fill(actual), match)
}

func (k propMatcher[T, U]) Write(w itsio.Writer) error {
	if err := k.description.Write(w); err != nil {
		return err
	}
	return k.m.Write(w.Indent())
}

func (k propMatcher[T, U]) String() string {
	return itskit.MatcherToString(k)
}
