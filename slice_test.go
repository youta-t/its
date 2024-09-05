package its_test

import "github.com/youta-t/its"

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
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./slice_test.go:21
	//     ✘ - /* got */ ?? == /* want */ 2		--- @ ./slice_test.go:21
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:21
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:21
	//
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./slice_test.go:26
	//     ✘ - /* got */ ?? == /* want */ 2		--- @ ./slice_test.go:26
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:26
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:26
}

func ExampleSlice_dfferent_length() {
	// fail. Actual has not enough 3.
	its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3, 3}).
		Match([]int{1, 2, 3}).
		OrError(t)

	// fail. Actual has too much 3.
	its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).
		Match([]int{1, 2, 3, 3}).
		OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 4; +0, -1)		--- @ ./slice_test.go:46
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:46
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:46
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:46
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:46
	//
	// ✘ []int{ ... (len: /* got */ 4, /* want */ 3; +1, -0)		--- @ ./slice_test.go:51
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:51
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:51
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:51
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
	// ✘ []int{ ... (len: /* got */ 0, /* want */ 1; +0, -1)		--- @ ./slice_test.go:74
	//     ✘ - /* got */ ?? == /* want */ 1		--- @ ./slice_test.go:74
	//
	// ✘ []int{ ... (len: /* got */ 1, /* want */ 0; +1, -0)		--- @ ./slice_test.go:77
	//     ✘ + /* got */ 1
}
