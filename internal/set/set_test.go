package set_test

import (
	"strings"
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

func cutLast(s string, sep string) string {
	i := strings.LastIndex(s, sep)
	if i < 0 {
		return s
	}
	return s[:i]
}

// ItsDiff checks equality between Matches, roughly.
func ItsDiff(want diff.Diff[itskit.Match]) itskit.Matcher[diff.Diff[itskit.Match]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return itskit.SimpleMatcher(
		func(got diff.Diff[itskit.Match]) bool {
			if want.Mode != got.Mode {
				return false
			}

			w := cutLast(want.Value.String(), "@")
			a := cutLast(got.Value.String(), "@")
			return w == a
		},
		"%s likes %s",
		itskit.Got, itskit.Want(want),
	)
}

func ItsMissingDiff[T any](matcher itskit.Matcher[T]) itskit.Matcher[diff.Diff[itskit.Match]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return ItsDiff(diff.MissingItem(diff.MissingMatch(matcher)))
}

func ItsExtraDiff[T any](v T) itskit.Matcher[diff.Diff[itskit.Match]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return ItsDiff(diff.ExtraItem(diff.ExtraMatch(v)))
}

func ItsOkDiff(match itskit.Match) itskit.Matcher[diff.Diff[itskit.Match]] {
	cancel := itskit.SkipStack()
	defer cancel()

	return ItsDiff(diff.OkItem(match))
}

func TestCompair(t *testing.T) {
	type When struct {
		values []int
		specs  []itskit.Matcher[int]
	}
	type Then []itskit.Matcher[diff.Diff[itskit.Match]]
	theory := func(when When, then Then) func(*testing.T) {
		return func(t *testing.T) {
			diffs := set.CompareWithMatcher[int](when.values, when.specs)
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
				t.Errorf("missing spec for diff: %+v", d)
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
			ItsOkDiff(IsOdd().Match(1)),
			ItsOkDiff(IsPrime().Match(2)),
			ItsOkDiff(IsPrime().Match(3)),
			ItsOkDiff(IsEven().Match(4)),
			ItsOkDiff(IsPrime().Match(5)),
			ItsOkDiff(IsEven().Match(6)),
			ItsOkDiff(IsPrime().Match(7)),
			ItsOkDiff(IsEven().Match(8)),
			ItsOkDiff(IsOdd().Match(9)),
			ItsOkDiff(IsEven().Match(10)),
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
			ItsOkDiff(IsOdd().Match(1)),
			ItsOkDiff(IsPrime().Match(2)),
			ItsOkDiff(IsPrime().Match(3)),
			ItsOkDiff(IsEven().Match(4)),
			ItsOkDiff(IsPrime().Match(5)),
			ItsOkDiff(IsEven().Match(6)),
			ItsOkDiff(IsPrime().Match(7)),
			ItsOkDiff(IsEven().Match(8)),
			ItsOkDiff(IsOdd().Match(9)),
			ItsOkDiff(IsEven().Match(10)),
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
			ItsExtraDiff(1),
			ItsOkDiff(IsPrime().Match(2)),
			ItsOkDiff(IsPrime().Match(3)),
			ItsOkDiff(IsEven().Match(4)),
			ItsOkDiff(IsPrime().Match(5)),
			ItsOkDiff(IsEven().Match(6)),
			ItsOkDiff(IsPrime().Match(7)),
			ItsOkDiff(IsEven().Match(8)),
			ItsExtraDiff(9),
			ItsOkDiff(IsEven().Match(10)),
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
			ItsOkDiff(IsOdd().Match(1)),
			ItsOkDiff(IsPrime().Match(2)),
			ItsOkDiff(IsPrime().Match(3)),
			ItsOkDiff(IsEven().Match(4)),
			ItsOkDiff(IsPrime().Match(5)),
			ItsOkDiff(IsEven().Match(6)),
			ItsOkDiff(IsPrime().Match(7)),
			ItsOkDiff(IsEven().Match(8)),
			ItsOkDiff(IsOdd().Match(9)),
			ItsOkDiff(IsEven().Match(10)),
			ItsMissingDiff(IsPrime()),
			ItsMissingDiff(IsPrime()),
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
			ItsExtraDiff(1),
			ItsOkDiff(IsPrime().Match(2)),
			ItsOkDiff(IsPrime().Match(3)),
			ItsOkDiff(IsEven().Match(4)),
			ItsOkDiff(IsPrime().Match(5)),
			ItsOkDiff(IsEven().Match(6)),
			ItsOkDiff(IsPrime().Match(7)),
			ItsOkDiff(IsEven().Match(8)),
			ItsExtraDiff(9),
			ItsOkDiff(IsEven().Match(10)),
			ItsMissingDiff(IsPrime()),
			ItsMissingDiff(IsPrime()),
		},
	))
}
