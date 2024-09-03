package its

import (
	"fmt"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type eqeqPtrMatcher[T comparable] struct {
	label itskit.Label
	want  *T
}

func ptrLabel[T any](v *T) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("&(%+v)", *v)
}

func (epm eqeqPtrMatcher[T]) Match(got *T) itskit.Match {
	ok := false

	if got == nil || epm.want == nil {
		ok = got == nil && epm.want == nil
	} else {
		ok = *got == *epm.want
	}

	return itskit.NewMatch(
		ok,
		epm.label.Fill(ptrLabel(got)),
	)
}

func (epm eqeqPtrMatcher[T]) Write(w itsio.Writer) error {
	return epm.label.Write(w)
}

func (epm eqeqPtrMatcher[T]) String() string {
	return itskit.MatcherToString[*T](epm)
}

// EqEqPtr tests of pointer for comparable with
//
//	(want == got) || (*want == *got)
//
// Deprecated: Use Pointer(EqEq(...)) .
func EqEqPtr[T comparable](want *T) Matcher[*T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return eqeqPtrMatcher[T]{
		label: itskit.NewLabelWithLocation(
			"%+v == %+v",
			itskit.Got, itskit.Want(ptrLabel(want)),
		),
		want: want,
	}
}
