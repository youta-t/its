// Code generated -- DO NOT EDIT

package gen_structer
import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal/type_test/sub2"
	
)


type Sub2Spec struct {
	IntField its.Matcher[int]
	
}

type _Sub2Matcher struct {
	label  itskit.Label
	fields []its.Matcher[testee.Sub2]
}

func ItsSub2(want Sub2Spec) its.Matcher[testee.Sub2] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.Sub2]{}
	
	{
		matcher := want.IntField
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[int]()
			} else {
				matcher = its.Always[int]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.Sub2, int](
				".IntField",
				func(got testee.Sub2) int { return got.IntField },
				matcher,
			),
		)
	}
	

	return _Sub2Matcher{
		label: itskit.NewLabelWithLocation("type Sub2:"),
		fields: sub,
	}
}

func (m _Sub2Matcher) Match(got testee.Sub2) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		m.label.Fill(got),
		sub...,
	)
}

func (m _Sub2Matcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type Sub2:", m.fields)
}

func (m _Sub2Matcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

