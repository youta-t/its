package its_test

import "github.com/youta-t/its"

func ExampleEqEq_ok() {
	its.EqEq(42).Match(42).OrError(t) // pass
	// Output:
}

func ExampleEqEq_ng() {
	its.EqEq(42).Match(49).OrError(t) // fail!
	// Output:
	// ✘ /* got */ 49 == /* want */ 42		--- @ ./eqeq_test.go:11
}

func ExampleEqEq_ng_non_primitive_type() {
	type MyType struct {
		Foo int
	}

	its.EqEq(MyType{Foo: 42}).Match(MyType{Foo: 24}).OrError(t) // also fail!

	// Output:
	// ✘ /* got */ {Foo:24} == /* want */ {Foo:42}		--- @ ./eqeq_test.go:21
}
