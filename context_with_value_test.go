package its_test

import (
	"context"

	"github.com/youta-t/its"
)

func ExampleContextWithValue_ok() {
	ctx := context.Background()

	type key string
	var k1 key = "key 1"
	var v1 string = "value 1"

	var k2 key = "key 2"
	var v2 int = 42

	var k3 key = "key 3"

	ctx = context.WithValue(ctx, k1, v1)
	ctx = context.WithValue(ctx, k2, v2)

	its.ContextWithValue(k1, its.EqEq("value 1")).Match(ctx).OrError(t)
	its.ContextWithValue(k2, its.EqEq(42)).Match(ctx).OrError(t)
	its.ContextWithValue(k3, its.Nil[any]()).Match(ctx).OrError(t)

	// Output:
}

func ExampleContextWithValue_ng_by_value() {
	ctx := context.Background()

	type key string
	var k1 key = "key 1"
	var v1 string = "value 1"

	var k2 key = "key 2"
	var v2 int = 42

	ctx = context.WithValue(ctx, k1, v1)
	ctx = context.WithValue(ctx, k2, v2)

	its.ContextWithValue(k1, its.EqEq("value 2")).Match(ctx).OrError(t)
	its.ContextWithValue(k2, its.EqEq(43)).Match(ctx).OrError(t)

	// Output:
	// ✘ // got = ctx.Value(key 1)		--- @ ./context_with_value_test.go:44
	//     ✘ /* got */ value 1 == /* want */ value 2		--- @ ./context_with_value_test.go:44
	//
	// ✘ // got = ctx.Value(key 2)		--- @ ./context_with_value_test.go:45
	//     ✘ /* got */ 42 == /* want */ 43		--- @ ./context_with_value_test.go:45
}

func ExampleContextWithValue_ng_by_type() {
	ctx := context.Background()

	type key string
	var k1 key = "key 1"
	var v1 string = "value 1"

	var k2 key = "key 2"
	var v2 int = 42

	ctx = context.WithValue(ctx, k1, v1)
	ctx = context.WithValue(ctx, k2, v2)

	its.ContextWithValue(k2, its.EqEq("value 1")).Match(ctx).OrError(t)
	its.ContextWithValue(k1, its.EqEq(42)).Match(ctx).OrError(t)

	// Output:
	// ✘ // got = ctx.Value(key 2)		--- @ ./context_with_value_test.go:68
	//     ✘ /* got */ 42 is not string
	//
	// ✘ // got = ctx.Value(key 1)		--- @ ./context_with_value_test.go:69
	//     ✘ /* got */ value 1 is not int
}

func ExampleContextWithValue_ng_by_nil_context() {
	var ctx context.Context = nil

	type key string
	var k1 key = "key 1"
	its.ContextWithValue(k1, its.Always[any]()).Match(ctx).OrError(t)

	// Output:
	// ✘ // given context.Context is nil		--- @ ./context_with_value_test.go:84
}
