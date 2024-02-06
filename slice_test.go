package its_test

import (
	"github.com/youta-t/its"
)

func ExampleSlice() {
	// pass
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Order is matter.
	its.Slice(
		its.EqEq(2), its.EqEq(1), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Actual has not enough 3.
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Actual has too much 3.
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 3},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
	//     ✘ - /* got */ ?? == /* want */ 2
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3
	//
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 4; +0, -1)
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✔ /* got */ 2 == /* want */ 2
	//     ✔ /* got */ 3 == /* want */ 3
	//     ✘ - /* got */ ?? == /* want */ 3
	//
	// ✘ []int{ ... (len: /* got */ 4, /* want */ 3; +1, -0)
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✔ /* got */ 2 == /* want */ 2
	//     ✔ /* got */ 3 == /* want */ 3
	//     ✘ + /* got */ 3
}

func ExampleSliceUnordered() {
	// pass
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// pass. order is not matter.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{3, 1, 2},
	).OrError(t)

	// fail. there is an extra item 42.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 42},
	).OrError(t)

	// fail. 3 is missing.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 4; +1, -0)
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✔ /* got */ 2 == /* want */ 2
	//     ✔ /* got */ 3 == /* want */ 3
	//     ✘ + /* got */ 42
	//
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 2; +0, -1)
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✔ /* got */ 2 == /* want */ 2
	//     ✘ - /* got */ ?? == /* want */ 3
}

func ExampleSliceUnorderedContaining() {
	// pass
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// pass. order is not matter.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{3, 1, 2},
	).OrError(t)

	// pass. extra item is okay.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 42},
	).OrError(t)

	// fail. 3 is missing.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered, contain; len: /* got */ 2, /* want */ 3; -1)
	//     ✔ /* got */ 1 == /* want */ 1
	//     ✔ /* got */ 2 == /* want */ 2
	//     ✘ - /* got */ ?? == /* want */ 3
}
