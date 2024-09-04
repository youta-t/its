package its

import "github.com/youta-t/its/itskit"

// always fail.
func Never[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()
	return itskit.SimpleMatcher(
		func(T) bool { return false },
		"(never pass)",
	)
}
