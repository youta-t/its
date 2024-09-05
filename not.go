package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type notMatcher[T any] struct {
	hedaer  itskit.Label
	matcher Matcher[T]
}

// Inverts matcher.
func Not[T any](matcher Matcher[T]) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return notMatcher[T]{
		hedaer:  itskit.NewLabelWithLocation("// not:"),
		matcher: matcher,
	}
}

func (n notMatcher[T]) Match(actual T) itskit.Match {
	m := n.matcher.Match(actual)
	return itskit.NewMatch(!m.Ok(), n.hedaer.Fill(actual), m)
}

func (n notMatcher[T]) Write(ww itsio.Writer) error {
	if err := n.hedaer.Write(ww); err != nil {
		return err
	}
	in := ww.Indent()
	return n.matcher.Write(in)
}

func (n notMatcher[T]) String() string {
	return itskit.MatcherToString[T](n)
}
