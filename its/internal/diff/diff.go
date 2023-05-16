package diff

import (
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
	expected := ExtraItem(value)
	return itskit.SimpleMatcher(
		func(s Diff) bool {
			if s.Mode != Extra {
				return false
			}
			return s.Match.String() == expected.Match.String()
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(expected),
	)
}

func IsMissing[T any](s itskit.Matcher[T]) itskit.Matcher[Diff] {
	expected := MissingItem[T](s)
	return itskit.SimpleMatcher(
		func(s Diff) bool {
			if s.Mode != Missing {
				return false
			}
			return s.Match.String() == expected.Match.String()
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(expected),
	)
}

func IsOk(m itskit.Match) itskit.Matcher[Diff] {
	expected := OkItem(m)
	return itskit.SimpleMatcher(
		func(d Diff) bool {
			if d.Mode != Ok {
				return false
			}
			return d.Match.String() == expected.Match.String()
		},
		"(%s) matches (%s)",
		itskit.Got, itskit.Want(expected),
	)
}
