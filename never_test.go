package its_test

import "github.com/youta-t/its"

func ExampleNever() {
	its.Never[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
	// ✘ (never pass)		--- @ ./never_test.go:6
}
