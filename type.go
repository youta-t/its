package its

import "github.com/youta-t/its/itskit"

// Type tests got value is a type.
func Type[T any]() Matcher[any] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(got any) bool {
			_, ok := got.(T)
			return ok
		},
		"%+v is a %T",
		itskit.Got, *new(T),
	)
}
