package its

import (
	"fmt"
	"sort"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type keyedMatch struct {
	strkey string
	match  itskit.Match
}

func (a keyedMatch) less(b keyedMatch) bool {
	return a.strkey < b.strkey
}

type mapSpec[K comparable, V any] map[K]itskit.Matcher[V]

type mapMatcher[K comparable, V any] struct {
	header itskit.Label
	spec   mapSpec[K, V]
}

type MapSpec[K comparable, V any] map[K]itskit.Matcher[V]

func Map[K comparable, V any](spec map[K]itskit.Matcher[V]) itskit.Matcher[map[K]V] {
	return mapMatcher[K, V]{
		header: itskit.NewLabel(
			"map[%T]%T{... ( keys: %d, %d; +%d, -%d )",
			*new(K), *new(V), itskit.Got, itskit.Want(len(spec)),
			itskit.Placeholder, itskit.Placeholder,
		),
		spec: spec,
	}
}

func (mm mapMatcher[K, V]) Match(actual map[K]V) itskit.Match {
	allkeys := map[K]struct{}{}
	for k := range actual {
		allkeys[k] = struct{}{}
	}
	for k := range mm.spec {
		allkeys[k] = struct{}{}
	}

	matches := []keyedMatch{}
	extra := 0
	miss := 0
	ng := 0

	for k := range allkeys {
		strkey := fmt.Sprintf("%+v", k)
		x, xok := mm.spec[k]
		if !xok {
			extra += 1
			matches = append(
				matches,
				keyedMatch{
					strkey: strkey,
					match: itskit.NG(
						fmt.Sprintf("%+v: (not in want)", k),
						itskit.NG(
							itskit.NewLabel(
								"%v, %v",
								itskit.Got, itskit.Want("??"),
							).Fill(actual[k]),
						),
					),
				},
			)
			continue
		}
		a, aok := actual[k]
		if !aok {
			miss += 1
			matches = append(
				matches,
				keyedMatch{
					strkey: strkey,
					match: itskit.NG(
						fmt.Sprintf("%+v: (not in got)", k),
						itskit.NG(mm.spec[k].String()),
					),
				},
			)
			continue
		}

		match := x.Match(a)
		matches = append(
			matches,
			keyedMatch{
				strkey: strkey,
				match: itskit.NewMatch(
					match.Ok(),
					fmt.Sprintf("%+v:", k),
					match,
				),
			},
		)
		if !match.Ok() {
			ng += 1
		}
	}

	submatches := make([]itskit.Match, 0, len(matches))

	sort.Slice(
		matches,
		func(a, b int) bool { return matches[a].less(matches[b]) },
	)
	for _, m := range matches {
		submatches = append(submatches, m.match)
	}

	return itskit.NewMatch(
		ng+extra+miss == 0,
		mm.header.Fill(len(actual), ng+extra, ng+miss),
		submatches...,
	)
}

func (mm mapMatcher[K, V]) Write(w itsio.Writer) error {
	if err := w.WriteStringln(mm.header.String()); err != nil {
		return err
	}
	iw := w.Indent()
	for k, m := range mm.spec {
		ks := fmt.Sprintf("%+v:", k)
		if err := iw.WriteStringln(ks); err != nil {
			return err
		}
		iiw := iw.Indent()
		if err := m.Write(iiw); err != nil {
			return err
		}
	}
	return nil
}

func (mm mapMatcher[K, V]) String() string {
	return itskit.MatcherToString[map[K]V](mm)
}

type mapContainingMatcher[K comparable, V any] struct {
	header itskit.Label
	spec   mapSpec[K, V]
}

func MapContaining[K comparable, V any](spec map[K]itskit.Matcher[V]) itskit.Matcher[map[K]V] {
	return mapContainingMatcher[K, V]{
		header: itskit.NewLabel(
			"map[%T]%T{ ... (contain; keys %d, %d; -%d)",
			*new(K), *new(V),
			itskit.Got, itskit.Want(len(spec)), itskit.Placeholder,
		),
		spec: spec,
	}
}

func (mcm mapContainingMatcher[K, V]) Match(actual map[K]V) itskit.Match {

	allkeys := map[K]struct{}{}
	for k := range actual {
		allkeys[k] = struct{}{}
	}
	for k := range mcm.spec {
		allkeys[k] = struct{}{}
	}
	matches := []keyedMatch{}
	extra := 0
	miss := 0
	ng := 0

	for k := range allkeys {
		strkey := fmt.Sprintf("%+v", k)
		x, xok := mcm.spec[k]
		if !xok {
			extra += 1
			matches = append(
				matches,
				keyedMatch{
					strkey: strkey,
					match: itskit.NG(
						fmt.Sprintf("%+v: (not in want)", k),
						itskit.NG(
							itskit.NewLabel(
								"%v, %v",
								itskit.Got, itskit.Want("??"),
							).Fill(actual[k]),
						),
					),
				},
			)
			continue
		}
		a, aok := actual[k]
		if !aok {
			miss += 1
			matches = append(
				matches,
				keyedMatch{
					strkey: strkey,
					match: itskit.NG(
						fmt.Sprintf("%+v: (not in got)", k),
						itskit.NG(mcm.spec[k].String()),
					),
				},
			)
			continue
		}

		match := x.Match(a)
		matches = append(
			matches,
			keyedMatch{
				strkey: strkey,
				match: itskit.NewMatch(
					match.Ok(),
					fmt.Sprintf("%+v:", k),
					match,
				),
			},
		)
		if !match.Ok() {
			ng += 1
		}
	}

	submatches := make([]itskit.Match, 0, len(matches))
	sort.Slice(
		matches,
		func(i, j int) bool { return matches[i].less(matches[j]) },
	)

	for _, m := range matches {
		submatches = append(submatches, m.match)
	}

	return itskit.NewMatch(
		ng+miss == 0,
		mcm.header.Fill(len(actual), ng+miss),
		submatches...,
	)
}

func (mm mapContainingMatcher[K, V]) Write(w itsio.Writer) error {
	if err := w.WriteStringln(mm.header.String()); err != nil {
		return err
	}
	iw := w.Indent()
	for k, m := range mm.spec {
		ks := fmt.Sprintf("%+v:", k)
		if err := iw.WriteStringln(ks); err != nil {
			return err
		}
		iiw := iw.Indent()
		if err := m.Write(iiw); err != nil {
			return err
		}
	}
	return nil
}

func (mm mapContainingMatcher[K, V]) String() string {
	return itskit.MatcherToString[map[K]V](mm)
}
