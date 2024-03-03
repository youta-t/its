//go:generate go run github.com/youta-t/its/structer
//go:generate gofmt -w ./gen_structer/type.go
package types

import (
	"io"

	"github.com/youta-t/its/structer/example/internal/type_test/sub1"
	"github.com/youta-t/its/structer/example/internal/type_test/sub2"
)

// perfect example
type T[P any] struct {
	// builtin
	F0 string
	F1 *string

	// imported from other package
	F2 sub1.Sub1
	F3 *sub1.Sub1

	// type parametered
	F4   G[int]
	F5   H[int, bool]
	F5_5 G[G[int]]

	// slice
	F6 []U
	F7 []int

	// array
	F8 [2]U
	F9 [2]int

	// func
	F10 func(int, bool) (string, error)

	// map
	F11 map[string]int
	F12 map[string]U
	F13 map[U]int

	// chan
	F14 chan int
	F15 <-chan int
	F16 chan<- int

	// struct
	F17 struct {
		Inline string
	}

	// interface
	F18 interface {
		M(string, X, ...int) (int, error)
		io.Writer
	}
	F19 G[G[int]]

	F20 private
	F21 G[private]
	F22 G[G[private]]
	f23 string

	// embedded
	U
	*X
	sub2.Sub2
	G[int]
}

func F() any {
	type Inner struct{}
	return Inner{}
}

type U struct {
	FieldU1 bool
}

type G[H any] struct {
	Fx H
}

type H[T, U any] struct {
	Fx T
	Fy U
}

type X string

type private struct{}

func init() {
	_ = T[string]{
		f23: "",
	}
}
