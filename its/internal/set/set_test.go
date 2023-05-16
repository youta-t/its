package set_test

import (
	"testing"

	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/internal/set"
	"github.com/youta-t/its/itskit"
)

func IsOdd() itskit.Matcher[int] {
	return itskit.SimpleMatcher(
		func(value int) bool { return value%2 == 1 },
		"%d is odd", itskit.Got,
	)
}

func IsEven() itskit.Matcher[int] {
	return itskit.SimpleMatcher(
		func(value int) bool { return value%2 == 0 },
		"%d is even", itskit.Got,
	)
}

func IsPrime() itskit.Matcher[int] {
	// for testing: support only primes <10.
	table := map[int]struct{}{
		2: {},
		3: {},
		5: {},
		7: {},
	}
	return itskit.SimpleMatcher(
		func(v int) bool {
			_, ok := table[v]
			return ok
		},
		"%d is prime", itskit.Got,
	)
}

func TestCompair(t *testing.T) {
	type When struct {
		values []int
		specs  []itskit.Matcher[int]
	}
	type Then []itskit.Matcher[diff.Diff]
	theory := func(when When, then Then) func(*testing.T) {
		return func(t *testing.T) {
			diffs := set.Compare(when.values, when.specs)
		SPEC:
			for _, spec := range then {
				for _, d := range diffs {
					m := spec.Match(d)
					if m.Ok() {
						continue SPEC
					}
				}
				t.Errorf("missing diff for spec: %s", spec)
			}

		DIFF:
			for _, d := range diffs {
				for _, spec := range then {
					m := spec.Match(d)
					if m.Ok() {
						continue DIFF
					}
				}
				t.Errorf("missing spec for diff: %s", d)
			}
		}
	}

	t.Run("equal sets match", theory(
		When{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			specs: []itskit.Matcher[int]{
				IsOdd(),
				IsPrime(),
				IsPrime(),
				IsEven(),
				IsPrime(),
				IsEven(),
				IsPrime(),
				IsEven(),
				IsOdd(),
				IsEven(),
			},
		},
		Then{
			diff.IsOk(IsOdd().Match(1)),
			diff.IsOk(IsPrime().Match(2)),
			diff.IsOk(IsPrime().Match(3)),
			diff.IsOk(IsEven().Match(4)),
			diff.IsOk(IsPrime().Match(5)),
			diff.IsOk(IsEven().Match(6)),
			diff.IsOk(IsPrime().Match(7)),
			diff.IsOk(IsEven().Match(8)),
			diff.IsOk(IsOdd().Match(9)),
			diff.IsOk(IsEven().Match(10)),
		},
	))

	t.Run("equal sets match regardless its order", theory(
		When{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			specs: []itskit.Matcher[int]{
				IsOdd(),
				IsOdd(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
			},
		},
		Then{
			diff.IsOk(IsOdd().Match(1)),
			diff.IsOk(IsPrime().Match(2)),
			diff.IsOk(IsPrime().Match(3)),
			diff.IsOk(IsEven().Match(4)),
			diff.IsOk(IsPrime().Match(5)),
			diff.IsOk(IsEven().Match(6)),
			diff.IsOk(IsPrime().Match(7)),
			diff.IsOk(IsEven().Match(8)),
			diff.IsOk(IsOdd().Match(9)),
			diff.IsOk(IsEven().Match(10)),
		},
	))

	t.Run("non-equal sets does not match: spec is too less", theory(
		When{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			specs: []itskit.Matcher[int]{
				// IsOdd(),
				// IsOdd(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
			},
		},
		Then{
			diff.IsExtra(1),
			diff.IsOk(IsPrime().Match(2)),
			diff.IsOk(IsPrime().Match(3)),
			diff.IsOk(IsEven().Match(4)),
			diff.IsOk(IsPrime().Match(5)),
			diff.IsOk(IsEven().Match(6)),
			diff.IsOk(IsPrime().Match(7)),
			diff.IsOk(IsEven().Match(8)),
			diff.IsExtra(9),
			diff.IsOk(IsEven().Match(10)),
		},
	))

	t.Run("non-equal sets does not match: spec is too many", theory(
		When{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			specs: []itskit.Matcher[int]{
				IsOdd(),
				IsOdd(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				// extra items
				IsPrime(),
				IsPrime(),
			},
		},
		Then{
			diff.IsOk(IsOdd().Match(1)),
			diff.IsOk(IsPrime().Match(2)),
			diff.IsOk(IsPrime().Match(3)),
			diff.IsOk(IsEven().Match(4)),
			diff.IsOk(IsPrime().Match(5)),
			diff.IsOk(IsEven().Match(6)),
			diff.IsOk(IsPrime().Match(7)),
			diff.IsOk(IsEven().Match(8)),
			diff.IsOk(IsOdd().Match(9)),
			diff.IsOk(IsEven().Match(10)),
			diff.IsMissing(IsPrime()),
			diff.IsMissing(IsPrime()),
		},
	))

	t.Run("non-equal sets does not match: spec is too many & too less", theory(
		When{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			specs: []itskit.Matcher[int]{
				// IsOdd(),
				// IsOdd(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsEven(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				IsPrime(),
				// extra items
				IsPrime(),
				IsPrime(),
			},
		},
		Then{
			diff.IsExtra(1),
			diff.IsOk(IsPrime().Match(2)),
			diff.IsOk(IsPrime().Match(3)),
			diff.IsOk(IsEven().Match(4)),
			diff.IsOk(IsPrime().Match(5)),
			diff.IsOk(IsEven().Match(6)),
			diff.IsOk(IsPrime().Match(7)),
			diff.IsOk(IsEven().Match(8)),
			diff.IsExtra(9),
			diff.IsOk(IsEven().Match(10)),
			diff.IsMissing(IsPrime()),
			diff.IsMissing(IsPrime()),
		},
	))
}
