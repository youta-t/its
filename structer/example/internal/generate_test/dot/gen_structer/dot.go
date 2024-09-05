// Code generated -- DO NOT EDIT

package gen_structer
import (
	"strings"

	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	config "github.com/youta-t/its/config"

	pkg1 "github.com/youta-t/its/structer/example/internal/generate_test/dot"
	
)


type DotStructSpec struct {
	
	Value its.Matcher[string]
	
}

type _DotStructMatcher struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.DotStruct]
}

func ItsDotStruct(want DotStructSpec) its.Matcher[pkg1.DotStruct] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.DotStruct]{}
	
	{
		matcher := want.Value
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[string]()
			} else {
				matcher = its.Always[string]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.DotStruct, string](
				".Value",
				func(got pkg1.DotStruct) string { return got.Value },
				matcher,
			),
		)
	}
	

	return _DotStructMatcher{
		label: itskit.NewLabelWithLocation("type DotStruct:"),
		fields: sub,
	}
}

func (m _DotStructMatcher) Match(got pkg1.DotStruct) itskit.Match {
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

func (m _DotStructMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type DotStruct:", m.fields)
}

func (m _DotStructMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}


type DotGSpec[T any] struct {
	
	Field its.Matcher[T]
	
}

type _DotGMatcher[T any] struct {
	label  itskit.Label
	fields []its.Matcher[pkg1.DotG[T]]
}

func ItsDotG[T any](want DotGSpec[T]) its.Matcher[pkg1.DotG[T]] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[pkg1.DotG[T]]{}
	
	{
		matcher := want.Field
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[T]()
			} else {
				matcher = its.Always[T]()
			}
		}
		sub = append(
			sub,
			its.Property[pkg1.DotG[T], T](
				".Field",
				func(got pkg1.DotG[T]) T { return got.Field },
				matcher,
			),
		)
	}
	

	return _DotGMatcher[T]{
		label: itskit.NewLabelWithLocation("type DotG:"),
		fields: sub,
	}
}

func (m _DotGMatcher[T]) Match(got pkg1.DotG[T]) itskit.Match {
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

func (m _DotGMatcher[T]) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type DotG:", m.fields)
}

func (m _DotGMatcher[T]) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}



func ItsDotMap(want its.Matcher[map[string]string]) its.Matcher[ pkg1.DotMap ] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _DotMapMatcher{matchers: want}
}

type _DotMapMatcher struct {
	matchers its.Matcher[map[string]string]
}

func (m _DotMapMatcher) Match(got pkg1.DotMap) itskit.Match {
	gotm := map[string]string(got)
	return m.matchers.Match(gotm)
}

func (m _DotMapMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type DotMap:", []its.Matcher[map[string]string]{m.matchers})
}

func (m _DotMapMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}



func ItsDotSlice(want its.Matcher[[]string]) its.Matcher[pkg1.DotSlice] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _DotSliceMatcher{ matcher: want }
}

type _DotSliceMatcher struct {
	matcher its.Matcher[[]string]
}

func (m _DotSliceMatcher) Match(got pkg1.DotSlice) itskit.Match {
	gots := []string(got)
	return m.matcher.Match(gots)
}

func (m _DotSliceMatcher) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type DotSlice:", []its.Matcher[[]string]{m.matcher})
}

func (m _DotSliceMatcher) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}


