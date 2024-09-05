//go:generate go run github.com/youta-t/its/structer
//go:generate gofmt -w ./gen_structer/type.go
package generatetest

import (
	"io"

	. "github.com/youta-t/its/structer/example/internal/generate_test/dot" // dot import to test
	"github.com/youta-t/its/structer/example/internal/generate_test/sub1"
	renamed "github.com/youta-t/its/structer/example/internal/generate_test/sub2"
)

// perfect example
type T[P interface {
	float32 | ~float64
	M() string
}] struct {
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
		I1
		I2
	}

	F20 private
	F21 G[private]
	F22 G[G[private]]
	f23 string

	// embedded
	U
	*X
	renamed.Sub2
	G[int]

	DS     DotStruct
	DSStar *DotStruct

	DI DotInterface

	DF DotFunc

	DN DotName

	DSlice DotSlice

	DMap DotMap

	DGene DotG[string]

	GDS G[DotStruct]
	GDI G[DotInterface]
	GDF G[DotFunc]
	GDN G[DotName]
}

func F() any {
	type Inner struct{}
	return Inner{}
}

type G[H any] struct {
	Fx H
}

type H[T, U any] struct {
	Fx T
	Fy U
}

type X string

type I1 interface {
	String() string
	Int() int
}

type private struct{}

func init() {
	_ = T[fs]{
		f23: "",
	}
}

type fs float64

func (fs) M() string {
	return ""
}

type Map map[string]int

type MapSV map[string]struct {
	Field1 int
}

type MapGV[T any] map[string]T

type MapGGV[T any] map[string]G[T]

type MapGK[K comparable] map[K]string

type MapSK map[struct {
	Field int
}]string

type Slice []string

type SliceS []struct {
	Field1 int
}

type SliceG[T any] []T
type SliceGG[T any] []G[T]

type S1 struct {
	Field string
}

type S2_2 S2
type S2 S1

type S3 struct {
	Map    Map
	MapSV  MapSV
	MapSK  MapSK
	Slice  Slice
	SliceS SliceS
}
