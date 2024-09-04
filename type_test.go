package its_test

import "github.com/youta-t/its"

func ExampleType_primitive() {
	its.Type[string]().Match("text value").OrError(t)
	its.Type[string]().Match(42).OrError(t)
	// Output:
	// ✘ /* got */ 42 is a string		--- @ ./type_test.go:7
}

func ExampleType_non_primitive() {
	type T struct {
		Foo int
	}
	type U struct {
		Bar int
	}
	its.Type[T]().Match(T{Foo: 42}).OrError(t)
	its.Type[U]().Match(T{Foo: 42}).OrError(t)
	// Output:
	// ✘ /* got */ {Foo:42} is a its_test.U		--- @ ./type_test.go:20
}
