package its_test

import "github.com/youta-t/its"

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
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 4; +1, -0)		--- @ ./slice_unordered_test.go:31
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_unordered_test.go:31
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_unordered_test.go:31
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_unordered_test.go:31
	//     ✘ + /* got */ 3
	//
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 2; +0, -1)		--- @ ./slice_unordered_test.go:36
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_unordered_test.go:36
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_unordered_test.go:36
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_unordered_test.go:36
}

func ExampleSliceUnordered_empty() {
	// pass. its empty.
	its.SliceUnordered[int]().Match([]int{}).OrError(t)

	// fail
	its.SliceUnordered(its.EqEq(1)).Match([]int{}).OrError(t)

	// fail. its empty.
	its.SliceUnordered[int]().Match([]int{1}).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered; len: /* want */ 1, /* got */ 0; +0, -1)		--- @ ./slice_unordered_test.go:58
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_unordered_test.go:58
	//
	// ✘ []int{ ... (unordered; len: /* want */ 0, /* got */ 1; +1, -0)		--- @ ./slice_unordered_test.go:61
	//     ✘ + /* got */ 1
}
