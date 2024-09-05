package its_test

import "github.com/youta-t/its"

func ExampleEqEqPtr_ok() {
	type S struct {
		Field int
	}

	got := &S{Field: 42}

	its.EqEqPtr(&S{Field: 42}).
		Match(got).
		OrError(t)
	// Output:
}

func ExampleEqEqPtr_ng_nonnil() {
	type S struct {
		Field int
	}

	got := &S{Field: 42}

	its.EqEqPtr(&S{Field: 24}).
		Match(got).
		OrError(t)
	// Output:
	// ✘ /* got */ *its_test.S is not nil,		--- @ ./eqeq_ptr_test.go:25
	//     ✘ /* got */ {Field:42} == /* want */ {Field:24}		--- @ ./eqeq_ptr_test.go:25
}

func ExampleEqEqPtr_ok_nil() {
	type S struct {
		Field int
	}

	var got *S = nil

	its.EqEqPtr[S](nil).
		Match(got).
		OrError(t)
	// Output:
}

func ExampleEqEqPtr_ng_nil() {
	type S struct {
		Field int
	}

	its.EqEqPtr(&S{Field: 42}).
		Match(nil).
		OrError(t)

	its.EqEqPtr[S](nil).
		Match(&S{Field: 42}).
		OrError(t)

	// Output:
	// ✘ /* got */ nil is not nil,		--- @ ./eqeq_ptr_test.go:51
	//     ✘ /* got */ ?? == /* want */ {Field:42}		--- @ ./eqeq_ptr_test.go:51
	//
	// ✘ (/* got */ {Field:42}) is nil		--- @ ./eqeq_ptr_test.go:55
}
