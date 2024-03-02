package editorialgraph_test

import (
	"strings"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
	"github.com/youta-t/its/itskit"
)

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

func TestEditorialGraph(t *testing.T) {
	type When struct {
		values []string
		specs  []itskit.Matcher[string]
	}
	type Then struct {
		matches []itskit.Matcher[diff.Diff[itskit.Match]]
	}

	theory := func(when When, then Then) func(t *testing.T) {
		return func(t *testing.T) {
			actual := editorialgraph.NewWithMatcher(when.values, when.specs)

			if a, x := len(actual), len(then.matches); a != x {
				t.Fatalf("len: %d, (want: %d)", a, x)
			}

			for i := range actual {
				m := then.matches[i].Match(actual[i])
				if !m.Ok() {
					t.Errorf("#%d: not match: %+v", i, m)
				}
			}
		}
	}

	t.Run("empty", theory(
		When{
			values: []string{},
			specs:  []itskit.Matcher[string]{},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{},
		},
	))

	t.Run("empty values", theory(
		When{
			values: []string{},
			specs: []itskit.Matcher[string]{
				its.EqEq("a"),
			},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsMissingDiff(its.EqEq("a")),
			},
		},
	))

	t.Run("empty matchers", theory(
		When{
			values: []string{"a"},
			specs:  []itskit.Matcher[string]{},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsExtraDiff("a"),
			},
		},
	))

	t.Run("all match", theory(
		When{
			values: strings.Split("abcdefg", ""),
			specs: []itskit.Matcher[string]{
				its.EqEq("a"),
				its.EqEq("b"),
				its.EqEq("c"),
				its.EqEq("d"),
				its.EqEq("e"),
				its.EqEq("f"),
				its.EqEq("g"),
			},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsOkDiff(its.EqEq("a").Match("a")),
				ItsOkDiff(its.EqEq("b").Match("b")),
				ItsOkDiff(its.EqEq("c").Match("c")),
				ItsOkDiff(its.EqEq("d").Match("d")),
				ItsOkDiff(its.EqEq("e").Match("e")),
				ItsOkDiff(its.EqEq("f").Match("f")),
				ItsOkDiff(its.EqEq("g").Match("g")),
			},
		},
	))

	t.Run("extra token", theory(
		When{
			values: strings.Split("abc!defg", ""),
			specs: []itskit.Matcher[string]{
				its.EqEq("a"),
				its.EqEq("b"),
				its.EqEq("c"),
				its.EqEq("d"),
				its.EqEq("e"),
				its.EqEq("f"),
				its.EqEq("g"),
			},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsOkDiff(its.EqEq("a").Match("a")),
				ItsOkDiff(its.EqEq("b").Match("b")),
				ItsOkDiff(its.EqEq("c").Match("c")),
				ItsExtraDiff("!"),
				ItsOkDiff(its.EqEq("d").Match("d")),
				ItsOkDiff(its.EqEq("e").Match("e")),
				ItsOkDiff(its.EqEq("f").Match("f")),
				ItsOkDiff(its.EqEq("g").Match("g")),
			},
		},
	))

	t.Run("missing token", theory(
		When{
			values: strings.Split("abcefg", ""),
			specs: []itskit.Matcher[string]{
				its.EqEq("a"),
				its.EqEq("b"),
				its.EqEq("c"),
				its.EqEq("d"),
				its.EqEq("e"),
				its.EqEq("f"),
				its.EqEq("g"),
			},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsOkDiff(its.EqEq("a").Match("a")),
				ItsOkDiff(its.EqEq("b").Match("b")),
				ItsOkDiff(its.EqEq("c").Match("c")),
				ItsMissingDiff(its.EqEq("d")),
				ItsOkDiff(its.EqEq("e").Match("e")),
				ItsOkDiff(its.EqEq("f").Match("f")),
				ItsOkDiff(its.EqEq("g").Match("g")),
			},
		},
	))

	t.Run("missing & extra token", theory(
		When{
			values: strings.Split("aabcxef", ""),
			specs: []itskit.Matcher[string]{
				its.EqEq("a"),
				its.EqEq("b"),
				its.EqEq("c"),
				its.EqEq("d"),
				its.EqEq("e"),
				its.EqEq("f"),
				its.EqEq("g"),
			},
		},
		Then{
			matches: []itskit.Matcher[diff.Diff[itskit.Match]]{
				ItsOkDiff(its.EqEq("a").Match("a")),
				ItsExtraDiff("a"),
				ItsOkDiff(its.EqEq("b").Match("b")),
				ItsOkDiff(its.EqEq("c").Match("c")),
				ItsMissingDiff(its.EqEq("d")),
				ItsExtraDiff("x"),
				ItsOkDiff(its.EqEq("e").Match("e")),
				ItsOkDiff(its.EqEq("f").Match("f")),
				ItsMissingDiff(its.EqEq("g")),
			},
		},
	))
}
