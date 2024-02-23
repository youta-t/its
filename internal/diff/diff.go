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

type Diff[T any] struct {
	Mode  Mode
	Value T
}

func OkItem[T any](m T) Diff[T] {
	return Diff[T]{Mode: Ok, Value: m}
}

func ExtraItem[T any](value T) Diff[T] {
	return Diff[T]{
		Mode:  Extra,
		Value: value,
	}
}

func MissingItem[T any](m T) Diff[T] {
	return Diff[T]{
		Mode:  Missing,
		Value: m,
	}
}

func ExtraMatch[T any](value T) itskit.Match {
	return itskit.NG(
		itskit.NewLabel("+ %+v", itskit.Got).Fill(value),
	)
}

func MissingMatch[T any](m itskit.Matcher[T]) itskit.Match {
	return itskit.NG(
		itskit.NewLabel("- %s", m.String()).Fill(itskit.Missing),
	)
}
