// Code generated -- DO NOT EDIT

package gen

import (
	"strings"

	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal"
	u_sub1 "github.com/youta-t/its/structer/example/internal/sub1"
	u_sub2 "github.com/youta-t/its/structer/example/internal/sub2"
)

type MyStructSpec struct {
	Name  itskit.Matcher[string]
	Value itskit.Matcher[int]
}

type _MyStructMatcher struct {
	fields []itskit.Matcher[testee.MyStruct]
}

func ItsMyStruct(want MyStructSpec) itskit.Matcher[testee.MyStruct] {
	sub := []itskit.Matcher[testee.MyStruct]{}

	sub = append(
		sub,
		itskit.Property(
			".Name",
			func(got testee.MyStruct) string { return got.Name },
			want.Name,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".Value",
			func(got testee.MyStruct) int { return got.Value },
			want.Value,
		),
	)

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
	Name   itskit.Matcher[string]
	Values itskit.Matcher[[]int]
	Sub1   itskit.Matcher[u_sub1.Sub1]
}

type _MyStruct1Matcher struct {
	fields []itskit.Matcher[testee.MyStruct1]
}

func ItsMyStruct1(want MyStruct1Spec) itskit.Matcher[testee.MyStruct1] {
	sub := []itskit.Matcher[testee.MyStruct1]{}

	sub = append(
		sub,
		itskit.Property(
			".Name",
			func(got testee.MyStruct1) string { return got.Name },
			want.Name,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".Values",
			func(got testee.MyStruct1) []int { return got.Values },
			want.Values,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".Sub1",
			func(got testee.MyStruct1) u_sub1.Sub1 { return got.Sub1 },
			want.Sub1,
		),
	)

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
	F0  itskit.Matcher[string]
	F1  itskit.Matcher[*string]
	F2  itskit.Matcher[u_sub1.Sub1]
	F3  itskit.Matcher[*u_sub1.Sub1]
	F4  itskit.Matcher[testee.G[int]]
	F5  itskit.Matcher[testee.H[int, bool]]
	F6  itskit.Matcher[[]testee.U]
	F7  itskit.Matcher[[]int]
	F8  itskit.Matcher[[2]testee.U]
	F9  itskit.Matcher[[2]int]
	F10 itskit.Matcher[func(int, bool) (string, error)]
	F11 itskit.Matcher[map[string]int]
	F12 itskit.Matcher[map[string]testee.U]
	F13 itskit.Matcher[map[testee.U]int]
	F14 itskit.Matcher[chan int]
	F15 itskit.Matcher[<-chan int]
	F16 itskit.Matcher[chan<- int]
	F17 itskit.Matcher[struct {
		Inline string
	}]
	F18 itskit.Matcher[interface {
		M(string, testee.X, ...int) (int, error)
	}]
	U    itskit.Matcher[testee.U]
	X    itskit.Matcher[*testee.X]
	Sub2 itskit.Matcher[u_sub2.Sub2]
	G    itskit.Matcher[testee.G[int]]
}

type _TMatcher[P any] struct {
	fields []itskit.Matcher[testee.T[P]]
}

func ItsT[P any](want TSpec[P]) itskit.Matcher[testee.T[P]] {
	sub := []itskit.Matcher[testee.T[P]]{}

	sub = append(
		sub,
		itskit.Property(
			".F0",
			func(got testee.T[P]) string { return got.F0 },
			want.F0,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F1",
			func(got testee.T[P]) *string { return got.F1 },
			want.F1,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F2",
			func(got testee.T[P]) u_sub1.Sub1 { return got.F2 },
			want.F2,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F3",
			func(got testee.T[P]) *u_sub1.Sub1 { return got.F3 },
			want.F3,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F4",
			func(got testee.T[P]) testee.G[int] { return got.F4 },
			want.F4,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F5",
			func(got testee.T[P]) testee.H[int, bool] { return got.F5 },
			want.F5,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F6",
			func(got testee.T[P]) []testee.U { return got.F6 },
			want.F6,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F7",
			func(got testee.T[P]) []int { return got.F7 },
			want.F7,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F8",
			func(got testee.T[P]) [2]testee.U { return got.F8 },
			want.F8,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F9",
			func(got testee.T[P]) [2]int { return got.F9 },
			want.F9,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F10",
			func(got testee.T[P]) func(int, bool) (string, error) { return got.F10 },
			want.F10,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F11",
			func(got testee.T[P]) map[string]int { return got.F11 },
			want.F11,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F12",
			func(got testee.T[P]) map[string]testee.U { return got.F12 },
			want.F12,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F13",
			func(got testee.T[P]) map[testee.U]int { return got.F13 },
			want.F13,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F14",
			func(got testee.T[P]) chan int { return got.F14 },
			want.F14,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F15",
			func(got testee.T[P]) <-chan int { return got.F15 },
			want.F15,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F16",
			func(got testee.T[P]) chan<- int { return got.F16 },
			want.F16,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F17",
			func(got testee.T[P]) struct {
				Inline string
			} {
				return got.F17
			},
			want.F17,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".F18",
			func(got testee.T[P]) interface {
				M(string, testee.X, ...int) (int, error)
			} {
				return got.F18
			},
			want.F18,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".U",
			func(got testee.T[P]) testee.U { return got.U },
			want.U,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".X",
			func(got testee.T[P]) *testee.X { return got.X },
			want.X,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".Sub2",
			func(got testee.T[P]) u_sub2.Sub2 { return got.Sub2 },
			want.Sub2,
		),
	)

	sub = append(
		sub,
		itskit.Property(
			".G",
			func(got testee.T[P]) testee.G[int] { return got.G },
			want.G,
		),
	)

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
