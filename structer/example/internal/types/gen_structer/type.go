// Code generated -- DO NOT EDIT

package gen_structer

import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal/types"
	u_sub1 "github.com/youta-t/its/structer/example/internal/types/sub1"
	u_sub2 "github.com/youta-t/its/structer/example/internal/types/sub2"
	u_io "io"
)

type TSpec[P interface {
	M() string
	float32 | ~float64
}] struct {
	F0   its.Matcher[string]
	F1   its.Matcher[*string]
	F2   its.Matcher[u_sub1.Sub1]
	F3   its.Matcher[*u_sub1.Sub1]
	F4   its.Matcher[testee.G[int]]
	F5   its.Matcher[testee.H[int, bool]]
	F5_5 its.Matcher[testee.G[testee.G[int]]]
	F6   its.Matcher[[]testee.U]
	F7   its.Matcher[[]int]
	F8   its.Matcher[[2]testee.U]
	F9   its.Matcher[[2]int]
	F10  its.Matcher[func(arg0 int, arg1 bool) (string, error)]
	F11  its.Matcher[map[string]int]
	F12  its.Matcher[map[string]testee.U]
	F13  its.Matcher[map[testee.U]int]
	F14  its.Matcher[chan int]
	F15  its.Matcher[<-chan int]
	F16  its.Matcher[chan<- int]
	F17  its.Matcher[struct {
		Inline string
	}]
	F18 its.Matcher[interface {
		M(arg0 string, arg1 testee.X, vararg ...int) (int, error)
		u_io.Writer
		testee.I1
		testee.I2
	}]
	U    its.Matcher[testee.U]
	X    its.Matcher[*testee.X]
	Sub2 its.Matcher[u_sub2.Sub2]
	G    its.Matcher[testee.G[int]]
}

type _TMatcher[P interface {
	M() string
	float32 | ~float64
}] struct {
	label  itskit.Label
	fields []its.Matcher[testee.T[P]]
}

func ItsT[P interface {
	M() string
	float32 | ~float64
}](want TSpec[P]) its.Matcher[testee.T[P]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.T[P]]{}

	{
		matcher := want.F0
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], string](
				".F0",
				func(got testee.T[P]) string { return got.F0 },
				matcher,
			),
		)
	}

	{
		matcher := want.F1
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*string]()
			} else {
				matcher = its.Always[*string]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], *string](
				".F1",
				func(got testee.T[P]) *string { return got.F1 },
				matcher,
			),
		)
	}

	{
		matcher := want.F2
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[u_sub1.Sub1]()
			} else {
				matcher = its.Always[u_sub1.Sub1]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], u_sub1.Sub1](
				".F2",
				func(got testee.T[P]) u_sub1.Sub1 { return got.F2 },
				matcher,
			),
		)
	}

	{
		matcher := want.F3
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*u_sub1.Sub1]()
			} else {
				matcher = its.Always[*u_sub1.Sub1]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], *u_sub1.Sub1](
				".F3",
				func(got testee.T[P]) *u_sub1.Sub1 { return got.F3 },
				matcher,
			),
		)
	}

	{
		matcher := want.F4
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[testee.G[int]]()
			} else {
				matcher = its.Always[testee.G[int]]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], testee.G[int]](
				".F4",
				func(got testee.T[P]) testee.G[int] { return got.F4 },
				matcher,
			),
		)
	}

	{
		matcher := want.F5
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[testee.H[int, bool]]()
			} else {
				matcher = its.Always[testee.H[int, bool]]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], testee.H[int, bool]](
				".F5",
				func(got testee.T[P]) testee.H[int, bool] { return got.F5 },
				matcher,
			),
		)
	}

	{
		matcher := want.F5_5
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[testee.G[testee.G[int]]]()
			} else {
				matcher = its.Always[testee.G[testee.G[int]]]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], testee.G[testee.G[int]]](
				".F5_5",
				func(got testee.T[P]) testee.G[testee.G[int]] { return got.F5_5 },
				matcher,
			),
		)
	}

	{
		matcher := want.F6
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]testee.U]()
			} else {
				matcher = its.Always[[]testee.U]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], []testee.U](
				".F6",
				func(got testee.T[P]) []testee.U { return got.F6 },
				matcher,
			),
		)
	}

	{
		matcher := want.F7
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]int]()
			} else {
				matcher = its.Always[[]int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], []int](
				".F7",
				func(got testee.T[P]) []int { return got.F7 },
				matcher,
			),
		)
	}

	{
		matcher := want.F8
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[2]testee.U]()
			} else {
				matcher = its.Always[[2]testee.U]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], [2]testee.U](
				".F8",
				func(got testee.T[P]) [2]testee.U { return got.F8 },
				matcher,
			),
		)
	}

	{
		matcher := want.F9
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[2]int]()
			} else {
				matcher = its.Always[[2]int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], [2]int](
				".F9",
				func(got testee.T[P]) [2]int { return got.F9 },
				matcher,
			),
		)
	}

	{
		matcher := want.F10
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[func(arg0 int, arg1 bool) (string, error)]()
			} else {
				matcher = its.Always[func(arg0 int, arg1 bool) (string, error)]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], func(arg0 int, arg1 bool) (string, error)](
				".F10",
				func(got testee.T[P]) func(arg0 int, arg1 bool) (string, error) { return got.F10 },
				matcher,
			),
		)
	}

	{
		matcher := want.F11
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[map[string]int]()
			} else {
				matcher = its.Always[map[string]int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], map[string]int](
				".F11",
				func(got testee.T[P]) map[string]int { return got.F11 },
				matcher,
			),
		)
	}

	{
		matcher := want.F12
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[map[string]testee.U]()
			} else {
				matcher = its.Always[map[string]testee.U]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], map[string]testee.U](
				".F12",
				func(got testee.T[P]) map[string]testee.U { return got.F12 },
				matcher,
			),
		)
	}

	{
		matcher := want.F13
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[map[testee.U]int]()
			} else {
				matcher = its.Always[map[testee.U]int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], map[testee.U]int](
				".F13",
				func(got testee.T[P]) map[testee.U]int { return got.F13 },
				matcher,
			),
		)
	}

	{
		matcher := want.F14
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[chan int]()
			} else {
				matcher = its.Always[chan int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], chan int](
				".F14",
				func(got testee.T[P]) chan int { return got.F14 },
				matcher,
			),
		)
	}

	{
		matcher := want.F15
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[<-chan int]()
			} else {
				matcher = its.Always[<-chan int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], <-chan int](
				".F15",
				func(got testee.T[P]) <-chan int { return got.F15 },
				matcher,
			),
		)
	}

	{
		matcher := want.F16
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[chan<- int]()
			} else {
				matcher = its.Always[chan<- int]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], chan<- int](
				".F16",
				func(got testee.T[P]) chan<- int { return got.F16 },
				matcher,
			),
		)
	}

	{
		matcher := want.F17
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[struct {
					Inline string
				}]()
			} else {
				matcher = its.Always[struct {
					Inline string
				}]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], struct {
				Inline string
			}](
				".F17",
				func(got testee.T[P]) struct {
					Inline string
				} {
					return got.F17
				},
				matcher,
			),
		)
	}

	{
		matcher := want.F18
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[interface {
					M(arg0 string, arg1 testee.X, vararg ...int) (int, error)
					u_io.Writer
					testee.I1
					testee.I2
				}]()
			} else {
				matcher = its.Always[interface {
					M(arg0 string, arg1 testee.X, vararg ...int) (int, error)
					u_io.Writer
					testee.I1
					testee.I2
				}]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], interface {
				M(arg0 string, arg1 testee.X, vararg ...int) (int, error)
				u_io.Writer
				testee.I1
				testee.I2
			}](
				".F18",
				func(got testee.T[P]) interface {
					M(arg0 string, arg1 testee.X, vararg ...int) (int, error)
					u_io.Writer
					testee.I1
					testee.I2
				} {
					return got.F18
				},
				matcher,
			),
		)
	}

	{
		matcher := want.U
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[testee.U]()
			} else {
				matcher = its.Always[testee.U]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], testee.U](
				".U",
				func(got testee.T[P]) testee.U { return got.U },
				matcher,
			),
		)
	}

	{
		matcher := want.X
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*testee.X]()
			} else {
				matcher = its.Always[*testee.X]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], *testee.X](
				".X",
				func(got testee.T[P]) *testee.X { return got.X },
				matcher,
			),
		)
	}

	{
		matcher := want.Sub2
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[u_sub2.Sub2]()
			} else {
				matcher = its.Always[u_sub2.Sub2]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], u_sub2.Sub2](
				".Sub2",
				func(got testee.T[P]) u_sub2.Sub2 { return got.Sub2 },
				matcher,
			),
		)
	}

	{
		matcher := want.G
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[testee.G[int]]()
			} else {
				matcher = its.Always[testee.G[int]]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.T[P], testee.G[int]](
				".G",
				func(got testee.T[P]) testee.G[int] { return got.G },
				matcher,
			),
		)
	}

	return _TMatcher[P]{
		label:  itskit.NewLabelWithLocation("type T:"),
		fields: sub,
	}
}

func (m _TMatcher[P]) Match(got testee.T[P]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		m.label.Fill(got),
		sub...,
	)
}

func (m _TMatcher[P]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type T:", m.fields)
}

func (m _TMatcher[P]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type USpec struct {
	FieldU1 its.Matcher[bool]
}

type _UMatcher struct {
	label  itskit.Label
	fields []its.Matcher[testee.U]
}

func ItsU(want USpec) its.Matcher[testee.U] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.U]{}

	{
		matcher := want.FieldU1
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[bool]()
			} else {
				matcher = its.Always[bool]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.U, bool](
				".FieldU1",
				func(got testee.U) bool { return got.FieldU1 },
				matcher,
			),
		)
	}

	return _UMatcher{
		label:  itskit.NewLabelWithLocation("type U:"),
		fields: sub,
	}
}

func (m _UMatcher) Match(got testee.U) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		m.label.Fill(got),
		sub...,
	)
}

func (m _UMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type U:", m.fields)
}

func (m _UMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type GSpec[H any] struct {
	Fx its.Matcher[H]
}

type _GMatcher[H any] struct {
	label  itskit.Label
	fields []its.Matcher[testee.G[H]]
}

func ItsG[H any](want GSpec[H]) its.Matcher[testee.G[H]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.G[H]]{}

	{
		matcher := want.Fx
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[H]()
			} else {
				matcher = its.Always[H]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.G[H], H](
				".Fx",
				func(got testee.G[H]) H { return got.Fx },
				matcher,
			),
		)
	}

	return _GMatcher[H]{
		label:  itskit.NewLabelWithLocation("type G:"),
		fields: sub,
	}
}

func (m _GMatcher[H]) Match(got testee.G[H]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		m.label.Fill(got),
		sub...,
	)
}

func (m _GMatcher[H]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type G:", m.fields)
}

func (m _GMatcher[H]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type HSpec[T any, U any] struct {
	Fx its.Matcher[T]
	Fy its.Matcher[U]
}

type _HMatcher[T any, U any] struct {
	label  itskit.Label
	fields []its.Matcher[testee.H[T, U]]
}

func ItsH[T any, U any](want HSpec[T, U]) its.Matcher[testee.H[T, U]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.H[T, U]]{}

	{
		matcher := want.Fx
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[T]()
			} else {
				matcher = its.Always[T]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.H[T, U], T](
				".Fx",
				func(got testee.H[T, U]) T { return got.Fx },
				matcher,
			),
		)
	}

	{
		matcher := want.Fy
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[U]()
			} else {
				matcher = its.Always[U]()
			}
		}
		sub = append(
			sub,
			its.Property[testee.H[T, U], U](
				".Fy",
				func(got testee.H[T, U]) U { return got.Fy },
				matcher,
			),
		)
	}

	return _HMatcher[T, U]{
		label:  itskit.NewLabelWithLocation("type H:"),
		fields: sub,
	}
}

func (m _HMatcher[T, U]) Match(got testee.H[T, U]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		m.label.Fill(got),
		sub...,
	)
}

func (m _HMatcher[T, U]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type H:", m.fields)
}

func (m _HMatcher[T, U]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
