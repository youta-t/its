package its

import (
	"strings"

	"github.com/youta-t/its/internal/diff"
	editorialgraph "github.com/youta-t/its/internal/editorial_graph"
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type textMatcher struct {
	label itskit.Label
	want  string
}

func (tm textMatcher) Match(got string) itskit.Match {
	gs := strings.SplitAfter(got, "\n")
	ws := strings.SplitAfter(tm.want, "\n")

	diffs := editorialgraph.New(
		gs, ws,
		func(s1, s2 string) (string, bool) {
			return s1, s1 == s2
		},
		func(s string) string { return s },
		func(s string) string { return s },
	)

	message := new(strings.Builder)
	message.WriteString(tm.label.String())
	message.WriteString("\n")

	unmatch := 0
	for _, d := range diffs {
		header := "      | "
		switch d.Mode {
		case diff.Extra:
			unmatch += 1
			header = "    + | "
		case diff.Missing:
			unmatch += 1
			header = "    - | "
		default:
		}

		val := d.Value
		if len(val) == 0 {
			message.WriteString(header[:len(header)-1])
			message.WriteString("\n")
			continue
		} else if val == "\n" {
			message.WriteString(header[:len(header)-1])
		} else {
			message.WriteString(header)
		}

		message.WriteString(d.Value)
	}

	return itskit.NewMatch(
		unmatch == 0,
		message.String(),
	)
}

func (tm textMatcher) Write(w itsio.Writer) error {
	if err := tm.label.Write(w); err != nil {
		return err
	}
	in := w.Indent()
	if err := in.WriteStringln(tm.want); err != nil {
		return err
	}
	return nil
}

func (tm textMatcher) String() string {
	return itskit.MatcherToString(tm)
}

// Text returns a matcher for a long text.
//
// When it get unmatch, it shows diff of text.
func Text(want string) Matcher[string] {
	cancel := itskit.SkipStack()
	defer cancel()

	return textMatcher{
		label: itskit.NewLabelWithLocation("(+ = got, - = want)"),
		want:  want,
	}
}
