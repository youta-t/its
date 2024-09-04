package its

import (
	"reflect"

	"github.com/youta-t/its/itskit"
)

// DeepEqual tests with
//
//	reflect.DeepEqual(want, got)
func DeepEqual[T any](want T) Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got T) bool { return reflect.DeepEqual(got, want) },
		"reflect.DeepEqual(%+v, %+v)",
		itskit.Got, itskit.Want(want),
	)
}
