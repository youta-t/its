//go:generate go run github.com/youta-t/its/structer -s MyStruct -s MyStruct1 -s T -dest gen
//go:generate gofmt -w ./gen/type.go
package types

import (
	"io"
	"time"

	"github.com/youta-t/its/structer/example/internal/sub1"
	"github.com/youta-t/its/structer/example/internal/sub2"
)

type MyStruct struct {
	Name      string
	Value     []int
	Timestamp time.Time
}

type MyStruct1 struct {
	Name   string
	Values []int
	Sub1   sub1.Sub1
}

type MyStruct2 struct {
	Name   string
	Values []int
	Sub1   sub2.Sub2
}

// perfect example
type T[P any] struct {
	// builtin
	F0 string
	F1 *string

	// imported from other package
	F2 sub1.Sub1
	F3 *sub1.Sub1

	// type parametered
	F4 G[int]
	F5 H[int, bool]

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
