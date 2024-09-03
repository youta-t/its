package its_test

import "github.com/youta-t/its"

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
	// ✘ []int{ ... (unordered, contain; len: /* got */ 2, /* want */ 3; -1)		--- @ ./slice_unordered_containing_test.go:41
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_unordered_containing_test.go:41
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_unordered_containing_test.go:41
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_unordered_containing_test.go:41
}

func ExampleSliceUnorderedContaining_empty() {
	// pass. its empty.
	its.SliceUnorderedContaining[int]().Match([]int{}).OrError(t)

	// fail
	its.SliceUnorderedContaining(its.EqEq(1)).Match([]int{}).OrError(t)

	// pass.
	its.SliceUnorderedContaining[int]().Match([]int{1}).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered, contain; len: /* got */ 0, /* want */ 1; -1)		--- @ ./slice_unordered_containing_test.go:57
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_unordered_containing_test.go:57
}
