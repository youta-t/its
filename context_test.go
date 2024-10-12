package its_test

import (
	"context"
	"time"

	"github.com/youta-t/its"
)

func ExampleContext() {
	ctx := context.Background()
	todo := context.TODO()

	withCancel, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	withoutCancel := context.WithoutCancel(withCancel)

	withDeadline, cancel2 := context.WithDeadline(ctx, time.Now().Add(-1*time.Second))
	defer cancel2()

	withTimeout, cancel3 := context.WithTimeout(ctx, 0*time.Second)
	defer cancel3()

	type key string
	var k key = "key"
	withValue := context.WithValue(ctx, k, "value")

	var nilContext context.Context = nil

	its.Context().Match(ctx).OrError(t)
	its.Context().Match(todo).OrError(t)
	its.Context().Match(withCancel).OrError(t)
	its.Context().Match(withoutCancel).OrError(t)
	its.Context().Match(withDeadline).OrError(t)
	its.Context().Match(withTimeout).OrError(t)
	its.Context().Match(withValue).OrError(t)
	its.Context().Match(nilContext).OrError(t)

	// Output:
}
