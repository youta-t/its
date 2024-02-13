// Code generated -- DO NOT EDIT

package gen
import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/mocker/internal/example"
	
)


type UserSpec struct {
	Id its.Matcher[string]
	Name its.Matcher[string]
	
}

type _UserMatcher struct {
	label  itskit.Label
	fields []its.Matcher[testee.User]
}

func ItsUser(want UserSpec) its.Matcher[testee.User] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.User]{}
	
	{
		matcher := want.Id
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.User, string](
				".Id",
				func(got testee.User) string { return got.Id },
				matcher,
			),
		)
	}
	
	{
		matcher := want.Name
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.User, string](
				".Name",
				func(got testee.User) string { return got.Name },
				matcher,
			),
		)
	}
	

	return _UserMatcher{
		label: itskit.NewLabelWithLocation("type User:"),
		fields: sub,
	}
}

func (m _UserMatcher) Match(got testee.User) itskit.Match {
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

func (m _UserMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type User:", m.fields)
}

func (m _UserMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}

