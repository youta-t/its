// Code generated -- DO NOT EDIT

package gen_structer

import (
	"strings"

	its "github.com/youta-t/its"
	config "github.com/youta-t/its/config"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	testee "github.com/youta-t/its/structer/example/internal"
	u_time "time"
)

type MyStructSpec struct {
	Name      its.Matcher[string]
	Value     its.Matcher[[]int]
	Timestamp its.Matcher[u_time.Time]
}

type _MyStructMatcher struct {
	label  itskit.Label
	fields []its.Matcher[testee.MyStruct]
}

func ItsMyStruct(want MyStructSpec) its.Matcher[testee.MyStruct] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[testee.MyStruct]{}

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
			itskit.Property[testee.MyStruct, string](
				".Name",
				func(got testee.MyStruct) string { return got.Name },
				matcher,
			),
		)
	}

	{
		matcher := want.Value
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[[]int]()
			} else {
				matcher = its.Always[[]int]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct, []int](
				".Value",
				func(got testee.MyStruct) []int { return got.Value },
				matcher,
			),
		)
	}

	{
		matcher := want.Timestamp
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[u_time.Time]()
			} else {
				matcher = its.Always[u_time.Time]()
			}
		}
		sub = append(
			sub,
			itskit.Property[testee.MyStruct, u_time.Time](
				".Timestamp",
				func(got testee.MyStruct) u_time.Time { return got.Timestamp },
				matcher,
			),
		)
	}

	return _MyStructMatcher{
		label:  itskit.NewLabelWithLocation("type MyStruct:"),
		fields: sub,
	}
}

func (m _MyStructMatcher) Match(got testee.MyStruct) itskit.Match {
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

func (m _MyStructMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type MyStruct:", m.fields)
}

func (m _MyStructMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
