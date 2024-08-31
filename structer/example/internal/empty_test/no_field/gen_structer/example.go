// Code generated -- DO NOT EDIT

package gen_structer
import (
	"strings"

	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	

	pkg1 "github.com/youta-t/its/structer/example/internal/empty_test/no_field"
	
)


type EmptySpec struct {
	
}

type _EmptyMatcher struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.Empty]
}

func ItsEmpty(want EmptySpec) its.Matcher[pkg1.Empty] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.Empty]{}
	

	return _EmptyMatcher{
		label: itskit.NewLabelWithLocation("type Empty:"),
		fields: sub,
	}
}

func (m _EmptyMatcher) Match(got pkg1.Empty) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(len(sub) == ok, m.label.Fill(got), sub...)
}

func (m _EmptyMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type Empty:", m.fields)
}

func (m _EmptyMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}






