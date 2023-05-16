// Code generated -- DO NOT EDIT

package gen

import (
	"strings"

	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal/sub2"
)

type Sub2Spec struct {
	IntField itskit.Matcher[int]
}

type _Sub2Matcher struct {
	fields []itskit.Matcher[testee.Sub2]
}

func ItsSub2(want Sub2Spec) itskit.Matcher[testee.Sub2] {
	sub := []itskit.Matcher[testee.Sub2]{}

	sub = append(
		sub,
		itskit.Property(
			".IntField",
			func(got testee.Sub2) int { return got.IntField },
			want.IntField,
		),
	)

	return _Sub2Matcher{fields: sub}
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
		itskit.NewLabel("type Sub2:").Fill(struct{}{}),
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
