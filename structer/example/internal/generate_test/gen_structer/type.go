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

	DSlice its.Matcher[pkg5.DotSlice]

	DMap its.Matcher[pkg5.DotMap]

	DGene its.Matcher[pkg5.DotG[string]]

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
		matcher := want.DSlice
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotSlice]()
			} else {
				matcher = its.Always[pkg5.DotSlice]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotSlice](
				".DSlice",
				func(got pkg1.T[P]) pkg5.DotSlice { return got.DSlice },
				matcher,
			),
		)
	}

	{
		matcher := want.DMap
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotMap]()
			} else {
				matcher = its.Always[pkg5.DotMap]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotMap](
				".DMap",
				func(got pkg1.T[P]) pkg5.DotMap { return got.DMap },
				matcher,
			),
		)
	}

	{
		matcher := want.DGene
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg5.DotG[string]]()
			} else {
				matcher = its.Always[pkg5.DotG[string]]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.T[P], pkg5.DotG[string]](
				".DGene",
				func(got pkg1.T[P]) pkg5.DotG[string] { return got.DGene },
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

type S1Spec struct {
	Field its.Matcher[string]
}

type _S1Matcher struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.S1]
}

func ItsS1(want S1Spec) its.Matcher[pkg1.S1] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.S1]{}

	{
		matcher := want.Field
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S1, string](
				".Field",
				func(got pkg1.S1) string { return got.Field },
				matcher,
			),
		)
	}

	return _S1Matcher{
		label:  itskit.NewLabelWithLocation("type S1:"),
		fields: sub,
	}
}

func (m _S1Matcher) Match(got pkg1.S1) itskit.Match {
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

func (m _S1Matcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type S1:", m.fields)
}

func (m _S1Matcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type S3Spec struct {
	Map its.Matcher[pkg1.Map]

	MapSV its.Matcher[pkg1.MapSV]

	MapSK its.Matcher[pkg1.MapSK]

	Slice its.Matcher[pkg1.Slice]

	SliceS its.Matcher[pkg1.SliceS]
}

type _S3Matcher struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.S3]
}

func ItsS3(want S3Spec) its.Matcher[pkg1.S3] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.S3]{}

	{
		matcher := want.Map
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.Map]()
			} else {
				matcher = its.Always[pkg1.Map]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S3, pkg1.Map](
				".Map",
				func(got pkg1.S3) pkg1.Map { return got.Map },
				matcher,
			),
		)
	}

	{
		matcher := want.MapSV
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.MapSV]()
			} else {
				matcher = its.Always[pkg1.MapSV]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S3, pkg1.MapSV](
				".MapSV",
				func(got pkg1.S3) pkg1.MapSV { return got.MapSV },
				matcher,
			),
		)
	}

	{
		matcher := want.MapSK
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.MapSK]()
			} else {
				matcher = its.Always[pkg1.MapSK]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S3, pkg1.MapSK](
				".MapSK",
				func(got pkg1.S3) pkg1.MapSK { return got.MapSK },
				matcher,
			),
		)
	}

	{
		matcher := want.Slice
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.Slice]()
			} else {
				matcher = its.Always[pkg1.Slice]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S3, pkg1.Slice](
				".Slice",
				func(got pkg1.S3) pkg1.Slice { return got.Slice },
				matcher,
			),
		)
	}

	{
		matcher := want.SliceS
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[pkg1.SliceS]()
			} else {
				matcher = its.Always[pkg1.SliceS]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.S3, pkg1.SliceS](
				".SliceS",
				func(got pkg1.S3) pkg1.SliceS { return got.SliceS },
				matcher,
			),
		)
	}

	return _S3Matcher{
		label:  itskit.NewLabelWithLocation("type S3:"),
		fields: sub,
	}
}

func (m _S3Matcher) Match(got pkg1.S3) itskit.Match {
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

func (m _S3Matcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type S3:", m.fields)
}

func (m _S3Matcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMap(want its.Matcher[map[string]int]) its.Matcher[pkg1.Map] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapMatcher{matchers: want}
}

type _MapMatcher struct {
	matchers its.Matcher[map[string]int]
}

func (m _MapMatcher) Match(got pkg1.Map) itskit.Match {
	gotm := map[string]int(got)
	return m.matchers.Match(gotm)
}

func (m _MapMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type Map:", []its.Matcher[map[string]int]{m.matchers})
}

func (m _MapMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMapSV(want its.Matcher[map[string]struct {
	Field1 int
}]) its.Matcher[pkg1.MapSV] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapSVMatcher{matchers: want}
}

type _MapSVMatcher struct {
	matchers its.Matcher[map[string]struct {
		Field1 int
	}]
}

func (m _MapSVMatcher) Match(got pkg1.MapSV) itskit.Match {
	gotm := map[string]struct {
		Field1 int
	}(got)
	return m.matchers.Match(gotm)
}

func (m _MapSVMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MapSV:", []its.Matcher[map[string]struct {
		Field1 int
	}]{m.matchers})
}

func (m _MapSVMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMapGV[T any](want its.Matcher[map[string]T]) its.Matcher[pkg1.MapGV[T]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapGVMatcher[T]{matchers: want}
}

type _MapGVMatcher[T any] struct {
	matchers its.Matcher[map[string]T]
}

func (m _MapGVMatcher[T]) Match(got pkg1.MapGV[T]) itskit.Match {
	gotm := map[string]T(got)
	return m.matchers.Match(gotm)
}

func (m _MapGVMatcher[T]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MapGV:", []its.Matcher[map[string]T]{m.matchers})
}

func (m _MapGVMatcher[T]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMapGGV[T any](want its.Matcher[map[string]pkg1.G[T]]) its.Matcher[pkg1.MapGGV[T]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapGGVMatcher[T]{matchers: want}
}

type _MapGGVMatcher[T any] struct {
	matchers its.Matcher[map[string]pkg1.G[T]]
}

func (m _MapGGVMatcher[T]) Match(got pkg1.MapGGV[T]) itskit.Match {
	gotm := map[string]pkg1.G[T](got)
	return m.matchers.Match(gotm)
}

func (m _MapGGVMatcher[T]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MapGGV:", []its.Matcher[map[string]pkg1.G[T]]{m.matchers})
}

func (m _MapGGVMatcher[T]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMapGK[K comparable](want its.Matcher[map[K]string]) its.Matcher[pkg1.MapGK[K]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapGKMatcher[K]{matchers: want}
}

type _MapGKMatcher[K comparable] struct {
	matchers its.Matcher[map[K]string]
}

func (m _MapGKMatcher[K]) Match(got pkg1.MapGK[K]) itskit.Match {
	gotm := map[K]string(got)
	return m.matchers.Match(gotm)
}

func (m _MapGKMatcher[K]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MapGK:", []its.Matcher[map[K]string]{m.matchers})
}

func (m _MapGKMatcher[K]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsMapSK(want its.Matcher[map[struct {
	Field int
}]string]) its.Matcher[pkg1.MapSK] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _MapSKMatcher{matchers: want}
}

type _MapSKMatcher struct {
	matchers its.Matcher[map[struct {
		Field int
	}]string]
}

func (m _MapSKMatcher) Match(got pkg1.MapSK) itskit.Match {
	gotm := map[struct {
		Field int
	}]string(got)
	return m.matchers.Match(gotm)
}

func (m _MapSKMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MapSK:", []its.Matcher[map[struct {
		Field int
	}]string]{m.matchers})
}

func (m _MapSKMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsSlice(want its.Matcher[[]string]) its.Matcher[pkg1.Slice] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _SliceMatcher{matcher: want}
}

type _SliceMatcher struct {
	matcher its.Matcher[[]string]
}

func (m _SliceMatcher) Match(got pkg1.Slice) itskit.Match {
	gots := []string(got)
	return m.matcher.Match(gots)
}

func (m _SliceMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type Slice:", []its.Matcher[[]string]{m.matcher})
}

func (m _SliceMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsSliceS(want its.Matcher[[]struct {
	Field1 int
}]) its.Matcher[pkg1.SliceS] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _SliceSMatcher{matcher: want}
}

type _SliceSMatcher struct {
	matcher its.Matcher[[]struct {
		Field1 int
	}]
}

func (m _SliceSMatcher) Match(got pkg1.SliceS) itskit.Match {
	gots := []struct {
		Field1 int
	}(got)
	return m.matcher.Match(gots)
}

func (m _SliceSMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type SliceS:", []its.Matcher[[]struct {
		Field1 int
	}]{m.matcher})
}

func (m _SliceSMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsSliceG[T any](want its.Matcher[[]T]) its.Matcher[pkg1.SliceG[T]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _SliceGMatcher[T]{matcher: want}
}

type _SliceGMatcher[T any] struct {
	matcher its.Matcher[[]T]
}

func (m _SliceGMatcher[T]) Match(got pkg1.SliceG[T]) itskit.Match {
	gots := []T(got)
	return m.matcher.Match(gots)
}

func (m _SliceGMatcher[T]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type SliceG:", []its.Matcher[[]T]{m.matcher})
}

func (m _SliceGMatcher[T]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

func ItsSliceGG[T any](want its.Matcher[[]pkg1.G[T]]) its.Matcher[pkg1.SliceGG[T]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _SliceGGMatcher[T]{matcher: want}
}

type _SliceGGMatcher[T any] struct {
	matcher its.Matcher[[]pkg1.G[T]]
}

func (m _SliceGGMatcher[T]) Match(got pkg1.SliceGG[T]) itskit.Match {
	gots := []pkg1.G[T](got)
	return m.matcher.Match(gots)
}

func (m _SliceGGMatcher[T]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type SliceGG:", []its.Matcher[[]pkg1.G[T]]{m.matcher})
}

func (m _SliceGGMatcher[T]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
