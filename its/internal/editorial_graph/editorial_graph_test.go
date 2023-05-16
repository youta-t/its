package editorialgraph_test

import (
	"strings"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
	"github.com/youta-t/its/itskit"
)

func TestEditorialGraph(t *testing.T) {
	type When struct {
		values []string
		specs  []itskit.Matcher[string]
	}
	type Then struct {
		matches []itskit.Matcher[diff.Diff]
	}

	theory := func(when When, then Then) func(t *testing.T) {
		return func(t *testing.T) {
			actual := editorialgraph.New(when.specs, when.values)

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
			matches: []itskit.Matcher[diff.Diff]{
				diff.IsOk(its.EqEq("a").Match("a")),
				diff.IsOk(its.EqEq("b").Match("b")),
				diff.IsOk(its.EqEq("c").Match("c")),
				diff.IsOk(its.EqEq("d").Match("d")),
				diff.IsOk(its.EqEq("e").Match("e")),
				diff.IsOk(its.EqEq("f").Match("f")),
				diff.IsOk(its.EqEq("g").Match("g")),
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
			matches: []itskit.Matcher[diff.Diff]{
				diff.IsOk(its.EqEq("a").Match("a")),
				diff.IsOk(its.EqEq("b").Match("b")),
				diff.IsOk(its.EqEq("c").Match("c")),
				diff.IsExtra("!"),
				diff.IsOk(its.EqEq("d").Match("d")),
				diff.IsOk(its.EqEq("e").Match("e")),
				diff.IsOk(its.EqEq("f").Match("f")),
				diff.IsOk(its.EqEq("g").Match("g")),
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
			matches: []itskit.Matcher[diff.Diff]{
				diff.IsOk(its.EqEq("a").Match("a")),
				diff.IsOk(its.EqEq("b").Match("b")),
				diff.IsOk(its.EqEq("c").Match("c")),
				diff.IsMissing(its.EqEq("d")),
				diff.IsOk(its.EqEq("e").Match("e")),
				diff.IsOk(its.EqEq("f").Match("f")),
				diff.IsOk(its.EqEq("g").Match("g")),
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
			matches: []itskit.Matcher[diff.Diff]{
				diff.IsOk(its.EqEq("a").Match("a")),
				diff.IsExtra("a"),
				diff.IsOk(its.EqEq("b").Match("b")),
				diff.IsOk(its.EqEq("c").Match("c")),
				diff.IsMissing(its.EqEq("d")),
				diff.IsExtra("x"),
				diff.IsOk(its.EqEq("e").Match("e")),
				diff.IsOk(its.EqEq("f").Match("f")),
				diff.IsMissing(its.EqEq("g")),
			},
		},
	))
}
