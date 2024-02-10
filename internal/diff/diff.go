package diff

import (
	"strings"

	"github.com/youta-t/its/itskit"
)

type Mode int

const (
	Missing = -1
	Ok      = 0
	Extra   = 1
)

type Diff struct {
	Mode  Mode
	Match itskit.Match
}

func (d Diff) String() string {
	return d.Match.String()
}

func OkItem(m itskit.Match) Diff {
	return Diff{Mode: Ok, Match: m}
}

func ExtraItem[T any](value T) Diff {
	return Diff{
		Mode: Extra,
		Match: itskit.NG(
			itskit.NewLabel("+ %+v", itskit.Got).Fill(value),
		),
	}
}

func MissingItem[T any](m itskit.Matcher[T]) Diff {
	return Diff{
		Mode: Missing,
		Match: itskit.NG(
			itskit.NewLabel("- %s", m.String()).Fill(itskit.Missing),
		),
	}
}

func IsExtra[T any](value T) itskit.Matcher[Diff] {
	want := ExtraItem(value)
	return itskit.SimpleMatcher(
		func(s Diff) bool {
			if s.Mode != Extra {
				return false
			}
			// ignore file:line
			dd, _, _ := strings.Cut(s.Match.String(), "--- @")
			ww, _, _ := strings.Cut(want.Match.String(), "--- @")
			return dd == ww
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(want),
	)
}

func IsMissing[T any](s itskit.Matcher[T]) itskit.Matcher[Diff] {
	want := MissingItem[T](s)
	return itskit.SimpleMatcher(
		func(s Diff) bool {
			if s.Mode != Missing {
				return false
			}
			// ignore file:line
			dd, _, _ := strings.Cut(s.Match.String(), "--- @")
			ww, _, _ := strings.Cut(want.Match.String(), "--- @")
			return dd == ww
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(want),
	)
}

func IsOk(m itskit.Match) itskit.Matcher[Diff] {
	want := OkItem(m)
	return itskit.SimpleMatcher(
		func(d Diff) bool {
			if d.Mode != Ok {
				return false
			}

			// ignore file:line
			dd, _, _ := strings.Cut(d.Match.String(), "--- @")
			ww, _, _ := strings.Cut(want.Match.String(), "--- @")
			return dd == ww
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(want),
	)
}
