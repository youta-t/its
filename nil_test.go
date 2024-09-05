package its_test

import (
	"errors"

	"github.com/youta-t/its"
)

func ExampleNil_ok() {
	var none *int
	its.Nil[*int]().Match(none).OrError(t)

	var chNil chan int = nil
	its.Nil[chan int]().Match(chNil).OrError(t)

	var fnNil func()
	its.Nil[func()]().Match(fnNil).OrError(t)

	its.Nil[any]().Match(nil).OrError(t)
	its.Nil[error]().Match(nil).OrError(t)

	var mapNil map[string]int
	its.Nil[map[string]int]().Match(mapNil).OrError(t)

	var sliceNil []int
	its.Nil[[]int]().Match(sliceNil).OrError(t)

	// Output:
}

func ExampleNil_ng_value_ptr() {
	three := 3
	its.Nil[*int]().Match(&three).OrError(t)

	// Output:
	// ✘ (/* got */ 3) is nil		--- @ ./nil_test.go:33
}

func ExampleNil_ng_chan() {
	chNonNil := make(chan int)
	its.Nil[chan int]().Match(chNonNil).OrError(t)
	// Output:
	// ✘ (/* got */ chan int) is nil		--- @ ./nil_test.go:41
}

func ExampleNil_ng_func() {
	funNonNil := func() {}
	its.Nil[func()]().Match(funNonNil).OrError(t)
	// Output:
	//
	// ✘ (/* got */ func()) is nil		--- @ ./nil_test.go:48
}

func ExampleNil_ng_nonpointer_value() {
	its.Nil[any]().Match(struct{}{}).OrError(t)
	its.Nil[int]().Match(3).OrError(t)
	// Output:
	// ✘ (/* got */ {}) is nil		--- @ ./nil_test.go:55
	//
	// ✘ (/* got */ 3) is nil		--- @ ./nil_test.go:56
}

func ExampleNil_ng_interface() {
	its.Nil[error]().Match(errors.New("error")).OrError(t)
	// Output:
	// ✘ (/* got */ error) is nil		--- @ ./nil_test.go:64
}

func ExampleNil_ng_map() {
	mapNonNil := map[string]int{}
	its.Nil[map[string]int]().Match(mapNonNil).OrError(t)
	// Output:
	//✘ (/* got */ map[]) is nil		--- @ ./nil_test.go:71
}

func ExampleNil_ng_slice() {
	sliceNonNil := []int{}
	its.Nil[[]int]().Match(sliceNonNil).OrError(t)

	// Output:
	//
	// ✘ (/* got */ []) is nil		--- @ ./nil_test.go:78
}
