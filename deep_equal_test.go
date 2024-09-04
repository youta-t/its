package its_test

import "github.com/youta-t/its"

func ExampleDeepEqual() {
	type MyType struct {
		Foo []int
	}

	its.DeepEqual(MyType{Foo: []int{42}}).Match(MyType{Foo: []int{42}}).OrError(t)
	its.DeepEqual(MyType{Foo: []int{42}}).Match(MyType{Foo: []int{24}}).OrError(t)

	// Output:
	// ✘ reflect.DeepEqual(/* got */ {Foo:[24]}, /* want */ {Foo:[42]})		--- @ ./deep_equal_test.go:11
}
