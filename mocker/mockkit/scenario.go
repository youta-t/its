package mockkit

import (
	"fmt"
	"reflect"

	"github.com/youta-t/its"
	"github.com/youta-t/its/itskit"
)

// Scenario describes the calling order of functions in the test target.
type Scenario interface {
	t() TestLike
	addWantCall(*token)
	call(*call)

	// End checks if all functions are called in the expected order.
	//
	// If it isn't, it reports a test error.
	End()
}

// BeginScenario create new empty scenatio.
//
// Scenario tests that all planned functions are called in an order as planned.
//
// You can plan functions with Next().
func BeginScenario(t TestLike) Scenario {
	return &scenario{t_: t}
}

type scenario struct {
	t_            TestLike
	wantCallOrder []*token
	gotCallOrder  []*call
}

func (s *scenario) t() TestLike {
	return s.t_
}

func (s *scenario) call(c *call) {
	s.gotCallOrder = append(s.gotCallOrder, c)
}

func (s *scenario) addWantCall(tok *token) {
	s.wantCallOrder = append(s.wantCallOrder, tok)
}

// End verify that all planned functions are called.
// If not, it reports test error.
func (s *scenario) End() {
	cancel := itskit.SkipStack()
	defer cancel()

	itskit.Named("// call order",
		its.ForItems(its.Slice, its.Equal, s.wantCallOrder),
	).
		Match(s.gotCallOrder).
		OrError(s.t())
}

type token struct {
	loc itskit.Location
}

func (t *token) String() string {
	return t.loc.String()
}

func (t *token) Equal(c *call) bool {
	return t == c.tok
}

type call struct {
	tok *token

	// where the call is invoked
	loc     itskit.Location
	args    []any
	returns []any
}

func (c *call) String() string {
	return fmt.Sprintf("%s (invoked at %s)", c.tok.loc, c.loc)
}

type tracedBehavior[T any] struct {
	sc    Scenario
	token *token
	base  FuncBehavior[T]
}

func (tr *tracedBehavior[T]) Fn(t TestLike) T {
	func_ := tr.base.Fn(t)

	rfn := reflect.ValueOf(func_)

	rwfn := reflect.MakeFunc(rfn.Type(), func(args []reflect.Value) (results []reflect.Value) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}

		cancel := itskit.SkipStack()
		defer cancel()

		a := make([]any, len(args))
		for i, arg := range args {
			a[i] = arg.Interface()
		}

		c := &call{args: a, tok: tr.token, loc: itskit.InvokedFrom()}

		if 0 < len(args) && rfn.Type().IsVariadic() {
			_args := args[:len(args)-1]
			varg := args[len(args)-1]
			if varg.Kind() != reflect.Slice {
				panic("non-slice variadic arg!")
			}
			l := varg.Len()
			for i := 0; i < l; i += 1 {
				_args = append(_args, varg.Index(i))
			}

			args = _args
		}

		outs := rfn.Call(args)

		o := make([]any, len(outs))
		for i, out := range outs {
			o[i] = out.Interface()
		}
		c.returns = o

		tr.sc.call(c)

		return outs
	})

	return rwfn.Interface().(T)
}

// Next register a FuncBehavior to the scenario.
//
// # Args
//
// - s Scenario: Testing Scenario
//
// - funcBehaviour mocker.FuncBehaviour[T]: a behaviour to be called next.
func Next[T any](s Scenario, funcBehaviour FuncBehavior[T]) FuncBehavior[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	rt := reflect.TypeOf(*new(T))
	if rt.Kind() != reflect.Func {
		panic(fmt.Errorf("[USAGE ERROR] func is requried, but %T", *new(T)))
	}

	tok := &token{loc: itskit.InvokedFrom()}
	s.addWantCall(tok)
	beh := &tracedBehavior[T]{
		sc:    s,
		token: tok,
		base:  funcBehaviour,
	}

	return beh
}
