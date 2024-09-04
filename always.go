package its

import "github.com/youta-t/its/itskit"

// always pass.
func Always[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return itskit.SimpleMatcher(
		func(T) bool { return true },
		"(always pass)",
	)
}
