package mockkit

import (
	"fmt"
	"reflect"

	"github.com/youta-t/its"
	"github.com/youta-t/its/itskit"
)

// TestLike is a type like *testing.T
type TestLike interface {
	Error(...any)
	Fatal(...any)
}

// FuncBehavior holds the expected behaveor of a function.
type FuncBehavior[F any] interface {
	// Fn build a function from the FuncBehaviour.
	Fn(TestLike) F
}

type wrappedBehavior[T any] struct {
	fn T
}

func (w wrappedBehavior[T]) Fn(TestLike) T {
	return w.fn
}

// Effect wraps a function as FuncBehavior
//
// The FuncBehaviour returned from this does not perform any match. (always pass)
func Effect[F any](func_ F) FuncBehavior[F] {
	rfn := reflect.ValueOf(func_)
	if rfn.Kind() != reflect.Func {
		panic("[USAGE ERROR] func should be passed")
	}
	return wrappedBehavior[F]{fn: func_}
}

// SequentialBehabiour is a container of FuncBehaviors
//
// For each invocation of its Mock(), it calles Mock() of registered FuncBehaviors in a order.
type SequentialBehavior[F any] interface {
	FuncBehavior[F]

	// Append registers FuncBehaviours to this SequentialBehaviour.
	Append(...FuncBehavior[F]) SequentialBehavior[F]
}

// Sequential builds SequentialBehavior.
//
// All passed FuncBehaviours are registered to the returned SequentialBehavior
// as the order in parameters.
func Sequential[F any](fb ...FuncBehavior[F]) SequentialBehavior[F] {
	rt := reflect.TypeOf(*new(F))
	switch rt.Kind() {
	case reflect.Func:
		// ok
	default:
		panic(fmt.Errorf("func is required, but %T", *new(F)))
	}

	b := &sequentialBehavior[F]{}
	b.Append(fb...)
	return b
}

type sequentialBehavior[T any] struct {
	b []FuncBehavior[T]
}

func (sb *sequentialBehavior[T]) Append(b ...FuncBehavior[T]) SequentialBehavior[T] {
	sb.b = append(sb.b, b...)

	return sb
}

func (sb *sequentialBehavior[T]) Fn(t TestLike) T {
	cancel := itskit.SkipStack()
	defer cancel()

	maximumInvoke := its.LesserEq(len(sb.b))

	nth := 1
	rt := reflect.TypeOf(*new(T))
	fn := reflect.MakeFunc(rt, func(args []reflect.Value) (results []reflect.Value) {
		cancel := itskit.SkipStack()
		defer cancel()
		defer func() { nth += 1 }()

		itskit.Named(
			itskit.NewLabelWithLocation("// invoke count :"),
			maximumInvoke,
		).Match(nth).OrFatal(t)

		b := sb.b[nth-1]

		fn := b.Fn(t)
		rfn := reflect.ValueOf(fn)

		if rfn.Type().IsVariadic() {
			args_ := args[:len(args)-1]

			vararg := args[len(args)-1]
			for i := 0; i < vararg.Len(); i += 1 {
				args_ = append(args_, vararg.Index(i))
			}

			args = args_
		}

		return rfn.Call(args)
	})

	return fn.Interface().(T)
}
