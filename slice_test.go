package its_test

import (
	"github.com/youta-t/its"
)

func ExampleSlice_ok() {
	// pass
	its.Slice(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{1, 2, 3}).
		OrError(t)

	// it is same thing as above
	its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// Output:
}

func ExampleSlice_different_order() {
	// fail. Order is matter.
	its.Slice(its.EqEq(2), its.EqEq(1), its.EqEq(3)).
		Match([]int{1, 2, 3}).
		OrError(t)

	// same as above
	its.ForItems(its.Slice, its.EqEq, []int{2, 1, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./slice_test.go:23
	//     ✘ - /* got */ ?? == /* want */ 2		--- @ ./slice_test.go:23
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:23
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:23
	//
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./slice_test.go:28
	//     ✘ - /* got */ ?? == /* want */ 2		--- @ ./slice_test.go:28
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:28
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:28
}

func ExampleSlice_different_length() {
	// fail. Actual has not enough 3.
	its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// fail. Actual has too much 3.
	its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3, 3}).
		OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 4; +0, -1)		--- @ ./slice_test.go:48
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:48
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:48
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:48
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:48
	//
	// ✘ []int{ ... (len: /* got */ 4, /* want */ 3; +1, -0)		--- @ ./slice_test.go:53
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:53
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:53
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:53
	//     ✘ + /* got */ 3
}

func ExampleSlice_empty() {
	// pass. its empty.
	its.Slice[int]().Match([]int{}).OrError(t)

	// fail
	its.Slice(its.EqEq(1)).Match([]int{}).OrError(t)

	// fail. its empty.
	its.Slice[int]().Match([]int{1}).OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 0, /* want */ 1; +0, -1)		--- @ ./slice_test.go:76
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_test.go:76
	//
	// ✘ []int{ ... (len: /* got */ 1, /* want */ 0; +1, -0)		--- @ ./slice_test.go:79
	//     ✘ + /* got */ 1
}

func ExampleSliceUnordered_ok() {
	// pass
	its.SliceUnordered(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{1, 2, 3}).
		OrError(t)

	// same as above
	its.ForItems(its.SliceUnordered, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// pass. order is not matter.
	its.SliceUnordered(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{3, 1, 2}).
		OrError(t)

	// same as above
	its.ForItems(its.SliceUnordered, its.EqEq, []int{1, 2, 3}).
		Match([]int{3, 1, 2}).
		OrError(t)

	// Output:
}

func ExampleSliceUnordered_different_length() {
	// fail. there is an extra item 42.
	its.ForItems(its.SliceUnordered, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3, 3}).
		OrError(t)

	// fail. 3 is missing.
	its.ForItems(its.SliceUnordered, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2}).
		OrError(t)

	// Output:
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 4; +1, -0)		--- @ ./slice_test.go:115
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:115
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:115
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:115
	//     ✘ + /* got */ 3
	//
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 2; +0, -1)		--- @ ./slice_test.go:120
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:120
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:120
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:120
}

func ExampleSliceUnordered_empty() {
	// pass. its empty.
	its.SliceUnordered[int]().Match([]int{}).OrError(t)

	// fail
	its.SliceUnordered(its.EqEq(1)).Match([]int{}).OrError(t)

	// fail. its empty.
	its.SliceUnordered[int]().Match([]int{1}).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered; len: /* want */ 1, /* got */ 0; +0, -1)		--- @ ./slice_test.go:142
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_test.go:142
	//
	// ✘ []int{ ... (unordered; len: /* want */ 0, /* got */ 1; +1, -0)		--- @ ./slice_test.go:145
	//     ✘ + /* got */ 1
}

func ExampleSliceUnorderedContaining_ok() {
	// pass
	its.SliceUnorderedContaining(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{1, 2, 3}).
		OrError(t)

	// same as above
	its.ForItems(its.SliceUnorderedContaining, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// pass. order is not matter.
	its.SliceUnorderedContaining(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{3, 1, 2}).
		OrError(t)

	// same as above
	its.ForItems(its.SliceUnorderedContaining, its.EqEq, []int{1, 2, 3}).
		Match([]int{3, 1, 2}).
		OrError(t)

	// pass. extra item is okay.
	its.SliceUnorderedContaining(its.EqEq(1), its.EqEq(2), its.EqEq(3)).
		Match([]int{1, 2, 3, 3}).
		OrError(t)

	// same as above
	its.ForItems(its.SliceUnorderedContaining, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3, 3}).
		OrError(t)

	// Ouptut:
}

func ExampleSliceUnorderedContaining_missing_item() {
	// fail. 3 is missing.
	its.ForItems(its.SliceUnorderedContaining, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2}).
		OrError(t)

	// Output:
	// ✘ []int{ ... (unordered, contain; len: /* got */ 2, /* want */ 3; -1)		--- @ ./slice_test.go:191
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:191
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:191
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:191
}

func ExampleSliceUnorderedContaining_empty() {
	// pass. its empty.
	its.SliceUnorderedContaining[int]().Match([]int{}).OrError(t)

	// fail
	its.SliceUnorderedContaining(its.EqEq(1)).Match([]int{}).OrError(t)

	// pass.
	its.SliceUnorderedContaining[int]().Match([]int{1}).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered, contain; len: /* got */ 0, /* want */ 1; -1)		--- @ ./slice_test.go:207
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_test.go:207
}
