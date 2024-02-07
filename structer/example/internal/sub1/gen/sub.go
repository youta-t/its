// Code generated -- DO NOT EDIT

package gen

import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal/sub1"
)

type Sub1Spec struct {
	StringField its.Matcher[string]
}

type _Sub1Matcher struct {
	fields []its.Matcher[testee.Sub1]
}

func ItsSub1(want Sub1Spec) its.Matcher[testee.Sub1] {
	sub := []its.Matcher[testee.Sub1]{}

	{
		matcher := want.StringField
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.Sub1, string](
				".StringField",
				func(got testee.Sub1) string { return got.StringField },
				matcher,
			),
		)
	}

	return _Sub1Matcher{fields: sub}
}

func (m _Sub1Matcher) Match(got testee.Sub1) itskit.Match {
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
		itskit.NewLabel("type Sub1:").Fill(struct{}{}),
		sub...,
	)
}

func (m _Sub1Matcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type Sub1:", m.fields)
}

func (m _Sub1Matcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
