package types_test

import (
	"io"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/structer/example/internal/types"
	"github.com/youta-t/its/structer/example/internal/types/gen_structer"
	"github.com/youta-t/its/structer/example/internal/types/sub1"
	"github.com/youta-t/its/structer/example/internal/types/sub2"
)

func TestT(T *testing.T) {
	testee := gen_structer.ItsT(gen_structer.TSpec[fs]{
		F0: its.EqEq("field F0"),
		F1: its.Pointer(its.EqEq("field F1")),

		F2: its.EqEq(sub1.Sub1{StringField: "F2.Sub1"}),
		F3: its.Pointer(its.EqEq(sub1.Sub1{StringField: "F3.Sub1"})),
		F4: gen_structer.ItsG(gen_structer.GSpec[int]{
			Fx: its.EqEq(42),
		}),
		F5: gen_structer.ItsH(gen_structer.HSpec[int, bool]{
			Fx: its.EqEq(42),
			Fy: its.EqEq(true),
		}),
		F5_5: gen_structer.ItsG(gen_structer.GSpec[types.G[int]]{
			Fx: gen_structer.ItsG(gen_structer.GSpec[int]{
				Fx: its.EqEq(42),
			}),
		}),

		F6: its.ForItems(its.Slice, gen_structer.ItsU, []gen_structer.USpec{
			{FieldU1: its.EqEq(true)},
			{FieldU1: its.EqEq(false)},
			{FieldU1: its.EqEq(true)},
		}),
		F7: its.ForItems(its.Slice, its.EqEq, []int{3, 2, 1}),

		F8: itskit.Property(
			"as slice",
			func(got [2]types.U) []types.U { return got[:] },
			its.ForItems(its.Slice, its.EqEq, []types.U{
				{FieldU1: false}, {FieldU1: true},
			}),
		),
		F9: itskit.Property(
			"as slice",
			func(got [2]int) []int { return got[:] },
			its.ForItems(its.Slice, its.EqEq, []int{4, 5}),
		),

		F10: its.Not(its.Nil[func(int, bool) (string, error)]()),

		F11: its.ForEntries(its.Map, its.EqEq, map[string]int{
			"a": 1, "b": 2, "c": 3,
		}),
		F12: its.ForEntries(its.Map, gen_structer.ItsU, map[string]gen_structer.USpec{
			"no":  {FieldU1: its.EqEq(false)},
			"yes": {FieldU1: its.EqEq(true)},
		}),
		F13: its.ForEntries(its.Map, its.EqEq, map[types.U]int{
			{FieldU1: false}: 0,
			{FieldU1: true}:  1,
		}),

		F14: its.ClosedChan[chan int](),
		F15: its.ClosedChan[<-chan int](),
		F16: its.Not(its.Nil[chan<- int]()),

		F17: its.EqEq(struct{ Inline string }{Inline: "F17.Inline"}),
		F18: its.Not(its.Nil[interface {
			types.I2
			types.I1
			io.Writer
			M(string, types.X, ...int) (int, error)
		}]()),

		U: gen_structer.ItsU(gen_structer.USpec{
			FieldU1: its.EqEq(true),
		}),
		X: its.Pointer(its.EqEq[types.X]("X")),
		Sub2: its.EqEq(sub2.Sub2{
			IntField: 123,
		}),
		G: gen_structer.ItsG(gen_structer.GSpec[int]{
			Fx: its.EqEq(987),
		}),
	})

	f1 := "field F1"
	x := types.X("x")

	ch := make(chan int, 1)
	close(ch)

	testee.Match(types.T[fs]{
		F0: "field 0",
		F1: &f1,

		F2: sub1.Sub1{StringField: "F2.Sub1"},
		F3: &sub1.Sub1{StringField: "F3.Sub1"},
		F4: types.G[int]{Fx: 42},
		F5: types.H[int, bool]{Fx: 42, Fy: true},
		F5_5: types.G[types.G[int]]{
			Fx: types.G[int]{
				Fx: 42,
			},
		},

		F6: []types.U{
			{FieldU1: true}, {FieldU1: false}, {FieldU1: true},
		},
		F7: []int{3, 2, 1},

		F8: [2]types.U{
			{FieldU1: false}, {FieldU1: true},
		},
		F9:  [2]int{4, 5},
		F10: func(i int, b bool) (string, error) { return "", nil },

		F11: map[string]int{
			"a": 1, "b": 2, "c": 3,
		},
		F12: map[string]types.U{
			"no":  {FieldU1: false},
			"yes": {FieldU1: true},
		},
		F13: map[types.U]int{
			{FieldU1: false}: 0,
			{FieldU1: true}:  1,
		},
		F14: ch,
		F15: recvOnly(ch),
		F16: sendOnly(ch),

		F17: struct{ Inline string }{Inline: "F17.Inline"},
		F18: F18{},

		U:    types.U{FieldU1: false},
		X:    &x,
		Sub2: sub2.Sub2{IntField: 123},
		G:    types.G[int]{Fx: 987},
	})
}

type fs float64

func (fs) M() string {
	return ""
}

func sendOnly[T any](ch chan T) chan<- T {
	return ch
}

func recvOnly[T any](ch chan T) <-chan T {
	return ch
}

type F18 struct{}

func (F18) Write(b []byte) (int, error) {
	return 0, io.EOF
}

func (F18) M(string, types.X, ...int) (int, error) {
	return 0, nil
}

func (F18) String() string {
	return ""
}

func (F18) Int() int {
	return 0
}
