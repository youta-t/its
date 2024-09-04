package its_test

import "github.com/youta-t/its"

func ExampleAlways() {
	its.Always[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
}
