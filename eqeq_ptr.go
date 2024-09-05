package its

import (
	"github.com/youta-t/its/itskit"
)

// EqEqPtr tests of pointer for comparable with
//
//	(want == got) || (*want == *got)
//
// Deprecated: Use Pointer(EqEq(...)) or Nil[T]() .
func EqEqPtr[T comparable](want *T) Matcher[*T] {
	cancel := itskit.SkipStack()
	defer cancel()
	if want == nil {
		return Nil[*T]()
	}
	return Pointer(EqEq(*want))
}
