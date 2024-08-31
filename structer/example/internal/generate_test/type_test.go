package generatetest_test

import (
	"io"
	"testing"

	"github.com/youta-t/its"
	types "github.com/youta-t/its/structer/example/internal/generate_test"
	"github.com/youta-t/its/structer/example/internal/generate_test/dot"
	"github.com/youta-t/its/structer/example/internal/generate_test/gen_structer"
	"github.com/youta-t/its/structer/example/internal/generate_test/sub1"
	"github.com/youta-t/its/structer/example/internal/generate_test/sub2"
)

func TestT(t *testing.T) {
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

		F6: its.ForItems(its.Slice, its.EqEq[types.U], []types.U{
			{FieldU1: true},
			{FieldU1: false},
			{FieldU1: true},
		}),
		F7: its.ForItems(its.Slice, its.EqEq, []int{3, 2, 1}),

		F8: its.Property(
			"as slice",
			func(got [2]types.U) []types.U { return got[:] },
			its.ForItems(its.Slice, its.EqEq, []types.U{
				{FieldU1: false}, {FieldU1: true},
			}),
		),
		F9: its.Property(
			"as slice",
			func(got [2]int) []int { return got[:] },
			its.ForItems(its.Slice, its.EqEq, []int{4, 5}),
		),

		F10: its.Not(its.Nil[func(int, bool) (string, error)]()),

		F11: its.ForEntries(its.Map, its.EqEq, map[string]int{
			"a": 1, "b": 2, "c": 3,
		}),
		F12: its.ForEntries(its.Map, its.EqEq, map[string]types.U{
			"no":  {FieldU1: false},
			"yes": {FieldU1: true},
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

		U: its.EqEq(types.U{FieldU1: true}),
		X: its.Pointer(its.EqEq[types.X]("X")),
		Sub2: its.EqEq(sub2.Sub2{
			IntField: 123,
		}),
		G: gen_structer.ItsG(gen_structer.GSpec[int]{
			Fx: its.EqEq(987),
		}),

		DS: its.EqEq(dot.DotStruct{
			Value: "E",
		}),

		DI: its.Nil[dot.DotInterface](),

		DF: its.Nil[dot.DotFunc](),

		DN: its.EqEq(dot.DotName(42)),

		DSStar: its.Pointer(its.EqEq(dot.DotStruct{
			Value: "EStar",
		})),

		GDS: gen_structer.ItsG(gen_structer.GSpec[dot.DotStruct]{
			Fx: its.EqEq(dot.DotStruct{
				Value: "GDS",
			}),
		}),

		GDI: gen_structer.ItsG(gen_structer.GSpec[dot.DotInterface]{
			Fx: its.Nil[dot.DotInterface](),
		}),

		GDF: gen_structer.ItsG(gen_structer.GSpec[dot.DotFunc]{
			Fx: its.Nil[dot.DotFunc](),
		}),

		GDN: gen_structer.ItsG(gen_structer.GSpec[dot.DotName]{
			Fx: its.EqEq(dot.DotName(42)),
		}),
	})

	f1 := "field F1"
	x := types.X("X")

	ch := make(chan int, 1)
	close(ch)

	testee.Match(types.T[fs]{
		F0: "field F0",
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

		U:    types.U{FieldU1: true},
		X:    &x,
		Sub2: sub2.Sub2{IntField: 123},
		G:    types.G[int]{Fx: 987},

		DS: dot.DotStruct{
			Value: "E",
		},
		DI: nil,
		DF: nil,
		DN: 42,

		DSStar: &dot.DotStruct{
			Value: "EStar",
		},

		GDS: types.G[dot.DotStruct]{
			Fx: dot.DotStruct{
				Value: "GDS",
			},
		},
		GDI: types.G[dot.DotInterface]{
			Fx: nil,
		},
		GDF: types.G[dot.DotFunc]{
			Fx: nil,
		},
		GDN: types.G[dot.DotName]{
			Fx: 42,
		},
	}).OrError(t)
}

func TestMap(t *testing.T) {
	{
		testee := gen_structer.ItsMap(its.Map(map[string]its.Matcher[int]{
			"key1": its.EqEq(1),
			"key2": its.EqEq(2),
		}))

		testee.Match(types.Map(map[string]int{
			"key1": 1,
			"key2": 2,
		})).OrError(t)
	}

	{
		testee := gen_structer.ItsMapSV(its.Map(map[string]its.Matcher[struct{ Field1 int }]{
			"key1": its.EqEq(struct{ Field1 int }{Field1: 1}),
			"key2": its.EqEq(struct{ Field1 int }{Field1: 2}),
		}))
		testee.Match(types.MapSV{
			"key1": {Field1: 1},
			"key2": {Field1: 2},
		}).OrError(t)
	}

	{
		testee := gen_structer.ItsMapGV(its.Map(map[string]its.Matcher[string]{
			"key1": its.EqEq("value1"),
		}))
		testee.Match(types.MapGV[string]{
			"key1": "value1",
		}).OrError(t)
	}

	{
		testee := gen_structer.ItsMapGGV(its.Map(map[string]its.Matcher[types.G[int]]{
			"key1": gen_structer.ItsG(gen_structer.GSpec[int]{
				Fx: its.EqEq(42),
			}),
		}))
		testee.Match(types.MapGGV[int]{
			"key1": types.G[int]{Fx: 42},
		}).OrError(t)
	}

	{
		testee := gen_structer.ItsMapGK(its.Map(map[string]its.Matcher[string]{
			"key1": its.EqEq("value1"),
		}))
		testee.Match(types.MapGK[string]{
			"key1": "value1",
		}).OrError(t)
	}

	{
		testee := gen_structer.ItsMapSK(its.Map(map[struct{ Field int }]its.Matcher[string]{
			{Field: 1}: its.EqEq("value1"),
			{Field: 2}: its.EqEq("value2"),
		}))
		testee.Match(types.MapSK{
			{Field: 1}: "value1",
			{Field: 2}: "value2",
		}).OrError(t)
	}
}

func TestSlice(t *testing.T) {
	{
		testee := gen_structer.ItsSlice(its.Slice(
			its.EqEq("a"),
			its.EqEq("b"),
			its.EqEq("c"),
		))
		testee.Match(types.Slice{"a", "b", "c"}).OrError(t)
	}

	{
		testee := gen_structer.ItsSliceS(its.Slice(
			its.EqEq(struct{ Field1 int }{Field1: 1}),
			its.EqEq(struct{ Field1 int }{Field1: 2}),
			its.EqEq(struct{ Field1 int }{Field1: 3}),
		))

		testee.Match(types.SliceS{{Field1: 1}, {Field1: 2}, {Field1: 3}}).OrError(t)
	}

	{
		testee := gen_structer.ItsSliceG(its.Slice(
			its.EqEq(1),
			its.EqEq(2),
			its.EqEq(3),
		))
		testee.Match(types.SliceG[int]{1, 2, 3}).OrError(t)
	}

	{
		testee := gen_structer.ItsSliceGG(its.Slice(
			gen_structer.ItsG[int](gen_structer.GSpec[int]{Fx: its.EqEq(1)}),
			gen_structer.ItsG[int](gen_structer.GSpec[int]{Fx: its.EqEq(2)}),
			gen_structer.ItsG[int](gen_structer.GSpec[int]{Fx: its.EqEq(3)}),
		))
		testee.Match([]types.G[int]{
			{Fx: 1}, {Fx: 2}, {Fx: 3},
		}).OrError(t)
	}
}

func TestS1(t *testing.T) {
	gen_structer.
		ItsS1(gen_structer.S1Spec{
			Field: its.EqEq("value"),
		}).
		Match(types.S1(types.S2{Field: "value"})).
		OrError(t)

	gen_structer.
		ItsS1(gen_structer.S1Spec{
			Field: its.EqEq("value"),
		}).
		Match(types.S1(types.S2_2{Field: "value"})).
		OrError(t)
}

func TestS3(t *testing.T) {
	testee := gen_structer.ItsS3(gen_structer.S3Spec{
		Map: gen_structer.ItsMap(its.Map(map[string]its.Matcher[int]{
			"key1": its.EqEq(1),
		})),
		MapSV: gen_structer.ItsMapSV(its.Map(map[string]its.Matcher[struct{ Field1 int }]{
			"key2": its.EqEq(struct{ Field1 int }{2}),
		})),
		MapSK: gen_structer.ItsMapSK(its.Map(map[struct{ Field int }]its.Matcher[string]{
			{Field: 3}: its.EqEq("value3"),
		})),
	})
	testee.Match(types.S3{
		Map: types.Map{
			"key1": 1,
		},
		MapSV: types.MapSV{
			"key2": struct{ Field1 int }{2},
		},
		MapSK: types.MapSK{
			{Field: 3}: "value3",
		},
	}).OrError(t)
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
