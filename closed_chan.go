package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type chanMatcher[T any, C chan T | <-chan T] struct {
	label itskit.Label
}

// ClosedChan tests wheather channel is closed or not.
//
// This matcher tries to receive from channel, it may cause sideeffect.
func ClosedChan[C chan T | <-chan T, T any]() Matcher[C] {
	cancel := itskit.SkipStack()
	defer cancel()
	return chanMatcher[T, C]{
		label: itskit.NewLabelWithLocation("chan %T is %s.", *new(T), itskit.Placeholder),
	}
}

func (c chanMatcher[T, C]) Match(ch C) itskit.Match {
	var closed bool
	select {
	case _, ok := <-ch:
		closed = !ok
	default:
		closed = false
	}

	message := "not closed"
	if closed {
		message = "closed"
	}
	return itskit.NewMatch(
		closed,
		c.label.Fill(itskit.Missing, message),
	)
}

func (c chanMatcher[T, C]) Write(ww itsio.Writer) error {
	return c.label.Write(ww)
}

func (c chanMatcher[T, C]) String() string {
	return itskit.MatcherToString(c)
}
