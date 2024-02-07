// Code generated -- DO NOT EDIT

package gen

import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal"
	u_sub1 "github.com/youta-t/its/structer/example/internal/sub1"
	u_sub2 "github.com/youta-t/its/structer/example/internal/sub2"
	u_time "time"
)

type MyStructSpec struct {
	Name      its.Matcher[string]
	Value     its.Matcher[[]int]
	Timestamp its.Matcher[u_time.Time]
}

type _MyStructMatcher struct {
	fields []its.Matcher[testee.MyStruct]
}

func ItsMyStruct(want MyStructSpec) its.Matcher[testee.MyStruct] {
	sub := []its.Matcher[testee.MyStruct]{}

	{
		matcher := want.Name
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct, string](
				".Name",
				func(got testee.MyStruct) string { return got.Name },
				matcher,
			),
		)
	}

	{
		matcher := want.Value
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]int]()
			} else {
				matcher = its.Always[[]int]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct, []int](
				".Value",
				func(got testee.MyStruct) []int { return got.Value },
				matcher,
			),
		)
	}

	{
		matcher := want.Timestamp
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[u_time.Time]()
			} else {
				matcher = its.Always[u_time.Time]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct, u_time.Time](
				".Timestamp",
				func(got testee.MyStruct) u_time.Time { return got.Timestamp },
				matcher,
			),
		)
	}

	return _MyStructMatcher{fields: sub}
}

func (m _MyStructMatcher) Match(got testee.MyStruct) itskit.Match {
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
		itskit.NewLabel("type MyStruct:").Fill(struct{}{}),
		sub...,
	)
}

func (m _MyStructMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MyStruct:", m.fields)
}

func (m _MyStructMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type MyStruct1Spec struct {
	Name   its.Matcher[string]
	Values its.Matcher[[]int]
	Sub1   its.Matcher[u_sub1.Sub1]
}

type _MyStruct1Matcher struct {
	fields []its.Matcher[testee.MyStruct1]
}

func ItsMyStruct1(want MyStruct1Spec) its.Matcher[testee.MyStruct1] {
	sub := []its.Matcher[testee.MyStruct1]{}

	{
		matcher := want.Name
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct1, string](
				".Name",
				func(got testee.MyStruct1) string { return got.Name },
				matcher,
			),
		)
	}

	{
		matcher := want.Values
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]int]()
			} else {
				matcher = its.Always[[]int]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct1, []int](
				".Values",
				func(got testee.MyStruct1) []int { return got.Values },
				matcher,
			),
		)
	}

	{
		matcher := want.Sub1
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[u_sub1.Sub1]()
			} else {
				matcher = its.Always[u_sub1.Sub1]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct1, u_sub1.Sub1](
				".Sub1",
				func(got testee.MyStruct1) u_sub1.Sub1 { return got.Sub1 },
				matcher,
			),
		)
	}

	return _MyStruct1Matcher{fields: sub}
}

func (m _MyStruct1Matcher) Match(got testee.MyStruct1) itskit.Match {
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
		itskit.NewLabel("type MyStruct1:").Fill(struct{}{}),
		sub...,
	)
}

func (m _MyStruct1Matcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MyStruct1:", m.fields)
}

func (m _MyStruct1Matcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

type TSpec[P any] struct {
	F0  its.Matcher[string]
	F1  its.Matcher[*string]
	F2  its.Matcher[u_sub1.Sub1]
	F3  its.Matcher[*u_sub1.Sub1]
	F4  its.Matcher[testee.G[int]]
	F5  its.Matcher[testee.H[int, bool]]
	F6  its.Matcher[[]testee.U]
	F7  its.Matcher[[]int]
	F8  its.Matcher[[2]testee.U]
	F9  its.Matcher[[2]int]
	F10 its.Matcher[func(int, bool) (string, error)]
	F11 its.Matcher[map[string]int]
	F12 its.Matcher[map[string]testee.U]
	F13 its.Matcher[map[testee.U]int]
	F14 its.Matcher[chan int]
	F15 its.Matcher[<-chan int]
	F16 its.Matcher[chan<- int]
	F17 its.Matcher[struct {
		Inline string
	}]
	F18 its.Matcher[interface {
		M(string, testee.X, ...int) (int, error)
	}]
	U    its.Matcher[testee.U]
	X    its.Matcher[*testee.X]
	Sub2 its.Matcher[u_sub2.Sub2]
	G    its.Matcher[testee.G[int]]
}

type _TMatcher[P any] struct {
	fields []its.Matcher[testee.T[P]]
}

func ItsT[P any](want TSpec[P]) its.Matcher[testee.T[P]] {
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
			itskit.Property[testee.T[P], string](
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
			itskit.Property[testee.T[P], *string](
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
			itskit.Property[testee.T[P], u_sub1.Sub1](
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
			itskit.Property[testee.T[P], *u_sub1.Sub1](
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
			itskit.Property[testee.T[P], testee.G[int]](
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
			itskit.Property[testee.T[P], testee.H[int, bool]](
				".F5",
				func(got testee.T[P]) testee.H[int, bool] { return got.F5 },
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
			itskit.Property[testee.T[P], []testee.U](
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
			itskit.Property[testee.T[P], []int](
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
			itskit.Property[testee.T[P], [2]testee.U](
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
			itskit.Property[testee.T[P], [2]int](
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
				matcher = its.Never[func(int, bool) (string, error)]()
			} else {
				matcher = its.Always[func(int, bool) (string, error)]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.T[P], func(int, bool) (string, error)](
				".F10",
				func(got testee.T[P]) func(int, bool) (string, error) { return got.F10 },
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
			itskit.Property[testee.T[P], map[string]int](
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
			itskit.Property[testee.T[P], map[string]testee.U](
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
			itskit.Property[testee.T[P], map[testee.U]int](
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
			itskit.Property[testee.T[P], chan int](
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
			itskit.Property[testee.T[P], <-chan int](
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
			itskit.Property[testee.T[P], chan<- int](
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
			itskit.Property[testee.T[P], struct {
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
					M(string, testee.X, ...int) (int, error)
				}]()
			} else {
				matcher = its.Always[interface {
					M(string, testee.X, ...int) (int, error)
				}]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.T[P], interface {
				M(string, testee.X, ...int) (int, error)
			}](
				".F18",
				func(got testee.T[P]) interface {
					M(string, testee.X, ...int) (int, error)
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
			itskit.Property[testee.T[P], testee.U](
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
			itskit.Property[testee.T[P], *testee.X](
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
			itskit.Property[testee.T[P], u_sub2.Sub2](
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
			itskit.Property[testee.T[P], testee.G[int]](
				".G",
				func(got testee.T[P]) testee.G[int] { return got.G },
				matcher,
			),
		)
	}

	return _TMatcher[P]{fields: sub}
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
		itskit.NewLabel("type T:").Fill(struct{}{}),
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
