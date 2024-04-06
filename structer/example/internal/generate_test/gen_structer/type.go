// Code generated -- DO NOT EDIT

package gen_structer

import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"

	pkg1 "github.com/youta-t/its/structer/example/internal/generate_test"
	pkg5 "github.com/youta-t/its/structer/example/internal/generate_test/dot"
	pkg2 "github.com/youta-t/its/structer/example/internal/generate_test/sub1"
	pkg4 "github.com/youta-t/its/structer/example/internal/generate_test/sub2"
	pkg3 "io"
)

type TSpec[P interface {
	M() string
	float32 | ~float64
}] struct {
	F0 its.Matcher[string]

	F1 its.Matcher[*string]

	F2 its.Matcher[pkg2.Sub1]

	F3 its.Matcher[*pkg2.Sub1]

	F4 its.Matcher[pkg1.G[int]]

	F5 its.Matcher[pkg1.H[int, bool]]

	F5_5 its.Matcher[pkg1.G[pkg1.G[int]]]

	F6 its.Matcher[[]pkg1.U]

	F7 its.Matcher[[]int]

	F8 its.Matcher[[2]pkg1.U]

	F9 its.Matcher[[2]int]

	F10 its.Matcher[func(arg0 int, arg1 bool) (string, error)]

	F11 its.Matcher[map[string]int]

	F12 its.Matcher[map[string]pkg1.U]

	F13 its.Matcher[map[pkg1.U]int]

	F14 its.Matcher[chan int]

	F15 its.Matcher[<-chan int]

	F16 its.Matcher[chan<- int]

	F17 its.Matcher[struct {
		Inline string
	}]

	F18 its.Matcher[interface {
		M(arg0 string, arg1 pkg1.X, vararg ...int) (int, error)
		pkg3.Writer
		pkg1.I1
		pkg1.I2
	}]

	U its.Matcher[pkg1.U]

	X its.Matcher[*pkg1.X]

	Sub2 its.Matcher[pkg4.Sub2]

	G its.Matcher[pkg1.G[int]]

	DS its.Matcher[pkg5.DotStruct]

	DSStar its.Matcher[*pkg5.DotStruct]

	DI its.Matcher[pkg5.DotInterface]

	DF its.Matcher[pkg5.DotFunc]

	DN its.Matcher[pkg5.DotName]

	GDS its.Matcher[pkg1.G[pkg5.DotStruct]]

	GDI its.Matcher[pkg1.G[pkg5.DotInterface]]

	GDF its.Matcher[pkg1.G[pkg5.DotFunc]]

	GDN its.Matcher[pkg1.G[pkg5.DotName]]
}

type _TMatcher[P interface {
	M() string
	float32 | ~float64
}] struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.T[P]]
}

func ItsT[P interface {
	M() string
	float32 | ~float64
}](want TSpec[P]) its.Matcher[pkg1.T[P]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.T[P]]{}

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
			its.Property[pkg1.T[P], string](
				".F0",
				func(got pkg1.T[P]) string { return got.F0 },
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
			its.Property[pkg1.T[P], *string](
				".F1",
				func(got pkg1.T[P]) *string { return got.F1 },
				matcher,
			),
		)
	}

	{
		matcher := want.F2
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg2.Sub1]()
			} else {
				matcher = its.Always[pkg2.Sub1]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg2.Sub1](
				".F2",
				func(got pkg1.T[P]) pkg2.Sub1 { return got.F2 },
				matcher,
			),
		)
	}

	{
		matcher := want.F3
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*pkg2.Sub1]()
			} else {
				matcher = its.Always[*pkg2.Sub1]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], *pkg2.Sub1](
				".F3",
				func(got pkg1.T[P]) *pkg2.Sub1 { return got.F3 },
				matcher,
			),
		)
	}

	{
		matcher := want.F4
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[int]]()
			} else {
				matcher = its.Always[pkg1.G[int]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[int]](
				".F4",
				func(got pkg1.T[P]) pkg1.G[int] { return got.F4 },
				matcher,
			),
		)
	}

	{
		matcher := want.F5
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.H[int, bool]]()
			} else {
				matcher = its.Always[pkg1.H[int, bool]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.H[int, bool]](
				".F5",
				func(got pkg1.T[P]) pkg1.H[int, bool] { return got.F5 },
				matcher,
			),
		)
	}

	{
		matcher := want.F5_5
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[pkg1.G[int]]]()
			} else {
				matcher = its.Always[pkg1.G[pkg1.G[int]]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[pkg1.G[int]]](
				".F5_5",
				func(got pkg1.T[P]) pkg1.G[pkg1.G[int]] { return got.F5_5 },
				matcher,
			),
		)
	}

	{
		matcher := want.F6
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]pkg1.U]()
			} else {
				matcher = its.Always[[]pkg1.U]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], []pkg1.U](
				".F6",
				func(got pkg1.T[P]) []pkg1.U { return got.F6 },
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
			its.Property[pkg1.T[P], []int](
				".F7",
				func(got pkg1.T[P]) []int { return got.F7 },
				matcher,
			),
		)
	}

	{
		matcher := want.F8
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[2]pkg1.U]()
			} else {
				matcher = its.Always[[2]pkg1.U]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], [2]pkg1.U](
				".F8",
				func(got pkg1.T[P]) [2]pkg1.U { return got.F8 },
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
			its.Property[pkg1.T[P], [2]int](
				".F9",
				func(got pkg1.T[P]) [2]int { return got.F9 },
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
			its.Property[pkg1.T[P], func(arg0 int, arg1 bool) (string, error)](
				".F10",
				func(got pkg1.T[P]) func(arg0 int, arg1 bool) (string, error) { return got.F10 },
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
			its.Property[pkg1.T[P], map[string]int](
				".F11",
				func(got pkg1.T[P]) map[string]int { return got.F11 },
				matcher,
			),
		)
	}

	{
		matcher := want.F12
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[map[string]pkg1.U]()
			} else {
				matcher = its.Always[map[string]pkg1.U]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], map[string]pkg1.U](
				".F12",
				func(got pkg1.T[P]) map[string]pkg1.U { return got.F12 },
				matcher,
			),
		)
	}

	{
		matcher := want.F13
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[map[pkg1.U]int]()
			} else {
				matcher = its.Always[map[pkg1.U]int]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], map[pkg1.U]int](
				".F13",
				func(got pkg1.T[P]) map[pkg1.U]int { return got.F13 },
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
			its.Property[pkg1.T[P], chan int](
				".F14",
				func(got pkg1.T[P]) chan int { return got.F14 },
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
			its.Property[pkg1.T[P], <-chan int](
				".F15",
				func(got pkg1.T[P]) <-chan int { return got.F15 },
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
			its.Property[pkg1.T[P], chan<- int](
				".F16",
				func(got pkg1.T[P]) chan<- int { return got.F16 },
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
			its.Property[pkg1.T[P], struct {
				Inline string
			}](
				".F17",
				func(got pkg1.T[P]) struct {
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
					M(arg0 string, arg1 pkg1.X, vararg ...int) (int, error)
					pkg3.Writer
					pkg1.I1
					pkg1.I2
				}]()
			} else {
				matcher = its.Always[interface {
					M(arg0 string, arg1 pkg1.X, vararg ...int) (int, error)
					pkg3.Writer
					pkg1.I1
					pkg1.I2
				}]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], interface {
				M(arg0 string, arg1 pkg1.X, vararg ...int) (int, error)
				pkg3.Writer
				pkg1.I1
				pkg1.I2
			}](
				".F18",
				func(got pkg1.T[P]) interface {
					M(arg0 string, arg1 pkg1.X, vararg ...int) (int, error)
					pkg3.Writer
					pkg1.I1
					pkg1.I2
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
				matcher = its.Never[pkg1.U]()
			} else {
				matcher = its.Always[pkg1.U]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.U](
				".U",
				func(got pkg1.T[P]) pkg1.U { return got.U },
				matcher,
			),
		)
	}

	{
		matcher := want.X
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*pkg1.X]()
			} else {
				matcher = its.Always[*pkg1.X]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], *pkg1.X](
				".X",
				func(got pkg1.T[P]) *pkg1.X { return got.X },
				matcher,
			),
		)
	}

	{
		matcher := want.Sub2
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg4.Sub2]()
			} else {
				matcher = its.Always[pkg4.Sub2]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg4.Sub2](
				".Sub2",
				func(got pkg1.T[P]) pkg4.Sub2 { return got.Sub2 },
				matcher,
			),
		)
	}

	{
		matcher := want.G
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[int]]()
			} else {
				matcher = its.Always[pkg1.G[int]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[int]](
				".G",
				func(got pkg1.T[P]) pkg1.G[int] { return got.G },
				matcher,
			),
		)
	}

	{
		matcher := want.DS
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotStruct]()
			} else {
				matcher = its.Always[pkg5.DotStruct]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotStruct](
				".DS",
				func(got pkg1.T[P]) pkg5.DotStruct { return got.DS },
				matcher,
			),
		)
	}

	{
		matcher := want.DSStar
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[*pkg5.DotStruct]()
			} else {
				matcher = its.Always[*pkg5.DotStruct]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], *pkg5.DotStruct](
				".DSStar",
				func(got pkg1.T[P]) *pkg5.DotStruct { return got.DSStar },
				matcher,
			),
		)
	}

	{
		matcher := want.DI
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotInterface]()
			} else {
				matcher = its.Always[pkg5.DotInterface]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotInterface](
				".DI",
				func(got pkg1.T[P]) pkg5.DotInterface { return got.DI },
				matcher,
			),
		)
	}

	{
		matcher := want.DF
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotFunc]()
			} else {
				matcher = its.Always[pkg5.DotFunc]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotFunc](
				".DF",
				func(got pkg1.T[P]) pkg5.DotFunc { return got.DF },
				matcher,
			),
		)
	}

	{
		matcher := want.DN
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotName]()
			} else {
				matcher = its.Always[pkg5.DotName]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotName](
				".DN",
				func(got pkg1.T[P]) pkg5.DotName { return got.DN },
				matcher,
			),
		)
	}

	{
		matcher := want.GDS
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[pkg5.DotStruct]]()
			} else {
				matcher = its.Always[pkg1.G[pkg5.DotStruct]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[pkg5.DotStruct]](
				".GDS",
				func(got pkg1.T[P]) pkg1.G[pkg5.DotStruct] { return got.GDS },
				matcher,
			),
		)
	}

	{
		matcher := want.GDI
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[pkg5.DotInterface]]()
			} else {
				matcher = its.Always[pkg1.G[pkg5.DotInterface]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[pkg5.DotInterface]](
				".GDI",
				func(got pkg1.T[P]) pkg1.G[pkg5.DotInterface] { return got.GDI },
				matcher,
			),
		)
	}

	{
		matcher := want.GDF
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[pkg5.DotFunc]]()
			} else {
				matcher = its.Always[pkg1.G[pkg5.DotFunc]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[pkg5.DotFunc]](
				".GDF",
				func(got pkg1.T[P]) pkg1.G[pkg5.DotFunc] { return got.GDF },
				matcher,
			),
		)
	}

	{
		matcher := want.GDN
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.G[pkg5.DotName]]()
			} else {
				matcher = its.Always[pkg1.G[pkg5.DotName]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg1.G[pkg5.DotName]](
				".GDN",
				func(got pkg1.T[P]) pkg1.G[pkg5.DotName] { return got.GDN },
				matcher,
			),
		)
	}

	return _TMatcher[P]{
		label:  itskit.NewLabelWithLocation("type T:"),
		fields: sub,
	}
}

func (m _TMatcher[P]) Match(got pkg1.T[P]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(len(sub) == ok, m.label.Fill(got), sub...)
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

type GSpec[H any] struct {
	Fx its.Matcher[H]
}

type _GMatcher[H any] struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.G[H]]
}

func ItsG[H any](want GSpec[H]) its.Matcher[pkg1.G[H]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.G[H]]{}

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
			its.Property[pkg1.G[H], H](
				".Fx",
				func(got pkg1.G[H]) H { return got.Fx },
				matcher,
			),
		)
	}

	return _GMatcher[H]{
		label:  itskit.NewLabelWithLocation("type G:"),
		fields: sub,
	}
}

func (m _GMatcher[H]) Match(got pkg1.G[H]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(len(sub) == ok, m.label.Fill(got), sub...)
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
	fields []its.Matcher[pkg1.H[T, U]]
}

func ItsH[T any, U any](want HSpec[T, U]) its.Matcher[pkg1.H[T, U]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.H[T, U]]{}

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
			its.Property[pkg1.H[T, U], T](
				".Fx",
				func(got pkg1.H[T, U]) T { return got.Fx },
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
			its.Property[pkg1.H[T, U], U](
				".Fy",
				func(got pkg1.H[T, U]) U { return got.Fy },
				matcher,
			),
		)
	}

	return _HMatcher[T, U]{
		label:  itskit.NewLabelWithLocation("type H:"),
		fields: sub,
	}
}

func (m _HMatcher[T, U]) Match(got pkg1.H[T, U]) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(len(sub) == ok, m.label.Fill(got), sub...)
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
