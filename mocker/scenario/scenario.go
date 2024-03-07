package scenario

import (
	"fmt"
	"reflect"

	"github.com/youta-t/its/itskit"
)

type Scenario interface {
	t() TestLike
	addWantCall(Token)
	call(*call)

	End()
}

type TestLike interface {
	Fatal(...any)
	Fatalf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
	Helper()
}

// Begin create new empty scenatio.
//
// Scenario tests that all planned functions are called in an order as planned.
//
// You can plan functions with Next().
func Begin(t TestLike) Scenario {
	return &scenario{t_: t}
}

type scenario struct {
	t_            TestLike
	wantCallOrder []Token
	next          int
}

func (s *scenario) t() TestLike {
	return s.t_
}

func (s *scenario) call(c *call) {
	cancel := itskit.SkipStack()
	defer cancel()

	if len(s.wantCallOrder) <= s.next {
		itskit.NG(
			itskit.NewLabelWithLocation(
				"// scenario error: no more planed calls",
			).String(),
		).OrError(s.t_)
		return
	}

	next := s.wantCallOrder[s.next]
	if next == c.tok {
		s.next += 1
		return
	}

	itskit.NG(
		itskit.NewLabelWithLocation(
			"// scenario error: call order is out of plan",
		).String(),
		itskit.NG(
			fmt.Sprintf("wanted to be called: %s", next.location()),
		),
	).OrError(s.t_)
}

func (s *scenario) addWantCall(tok Token) {
	s.wantCallOrder = append(s.wantCallOrder, tok)
}

// End verify that all planned functions are called.
// If not, it reports test error.
func (s *scenario) End() {
	if s.next == len(s.wantCallOrder) {
		return
	}
	cancel := itskit.SkipStack()
	defer cancel()

	reminders := []itskit.Match{}
	for _, rem := range s.wantCallOrder[s.next:] {
		reminders = append(reminders, itskit.NG(
			rem.location().String(),
		))
	}
	itskit.NG(
		itskit.NewLabelWithLocation("// there are functions planned but not called").String(),
		reminders...,
	).OrError(s.t_)
}

type token struct {
	loc itskit.Location
}

func (t *token) location() itskit.Location {
	return t.loc
}

// Token tracks a function planned in the scenario.
type Token interface {
	// where the func is planned
	location() itskit.Location
}

type call struct {
	tok *token

	// where the call is invoked
	loc     itskit.Location
	args    []any
	returns []any
}

// Next plan a function to be called in the scenario test.
//
// If the (returned) function is called, but the call-order is wrong,
// it reports test error.
func Next[T any](s Scenario, fn T) T {
	cancel := itskit.SkipStack()
	defer cancel()

	rfn := reflect.ValueOf(fn)

	switch rfn.Kind() {
	case reflect.Func:
		// ok!
	default:
		s.t().Fatalf("[USAGE ERROR] mocker.Next: fn should be a function, but: %+v", fn)
	}

	tok := &token{
		loc: itskit.InvokedFrom(),
	}

	rwfn := reflect.MakeFunc(rfn.Type(), func(args []reflect.Value) (results []reflect.Value) {
		s.t().Helper()
		cancel := itskit.SkipStack()
		defer cancel()

		a := make([]any, len(args))
		for i, arg := range args {
			a[i] = arg.Interface()
		}

		c := &call{
			args: a, tok: tok, loc: itskit.InvokedFrom(),
		}

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

		s.call(c)

		return outs
	})

	wfn := rwfn.Interface()
	w, ok := wfn.(T)
	if !ok {
		s.t().Fatal("[BUG]: mock cast failed")
	}

	s.addWantCall(tok)
	return w
}
