package its

import (
	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

// All tests actual passes all specs.
//
// If no matchers are given, it always pass.
func All[T any](matchers ...itskit.Matcher[T]) itskit.Matcher[T] {
	return allMatcher[T]{
		header: itskit.NewLabel(
			"// all: (%d ok / %d matchers)",
			itskit.Placeholder, len(matchers),
		),
		specs: matchers,
	}
}

type allMatcher[T any] struct {
	header itskit.Label
	specs  []itskit.Matcher[T]
}

func (as allMatcher[T]) Match(actual T) itskit.Match {
	matches := make([]itskit.Match, len(as.specs))
	nSpec := len(as.specs)
	nOk := 0
	for nth, s := range as.specs {
		m := s.Match(actual)
		matches[nth] = m
		if m.Ok() {
			nOk += 1
		}
	}
	return itskit.NewMatch(
		nOk == nSpec,
		as.header.Fill(actual, nOk),
		matches...,
	)
}

func (as allMatcher[T]) Write(ww itsio.Writer) error {
	if err := ww.WriteStringln(as.header.String()); err != nil {
		return err
	}
	in := ww.Indent()
	for _, m := range as.specs {
		if err := m.Write(in); err != nil {
			return err
		}
	}
	return nil
}

func (as allMatcher[T]) String() string {
	return itskit.MatcherToString[T](as)
}

// All tests actual passes at least one spec.
//
// If no matchers are given, it always fail.
func Some[T any](matchers ...itskit.Matcher[T]) itskit.Matcher[T] {
	return someMatcher[T]{
		header: itskit.NewLabel(
			"// some: (%d ok / %d matchers)",
			itskit.Placeholder, len(matchers),
		),
		submatch: matchers,
	}
}

type someMatcher[T any] struct {
	header   itskit.Label
	submatch []itskit.Matcher[T]
}

func (ss someMatcher[T]) Match(actual T) itskit.Match {
	matches := make([]itskit.Match, len(ss.submatch))
	nOk := 0
	for nth, s := range ss.submatch {
		m := s.Match(actual)
		matches[nth] = m
		if m.Ok() {
			nOk += 1
		}
	}
	return itskit.NewMatch(
		0 < nOk,
		ss.header.Fill(actual, nOk),
		matches...,
	)
}

func (ss someMatcher[T]) Write(ww itsio.Writer) error {
	if err := ww.WriteStringln(ss.header.String()); err != nil {
		return err
	}
	in := ww.Indent()
	for _, s := range ss.submatch {
		if err := s.Write(in); err != nil {
			return err
		}
	}
	return nil
}

func (ss someMatcher[T]) String() string {
	return itskit.MatcherToString[T](ss)
}

type notMatcher[T any] struct {
	hedaer  string
	matcher itskit.Matcher[T]
}

// Inverts matcher.
func Not[T any](matcher itskit.Matcher[T]) itskit.Matcher[T] {
	return notMatcher[T]{
		hedaer:  "// not:",
		matcher: matcher,
	}
}

func (n notMatcher[T]) Match(actual T) itskit.Match {
	m := n.matcher.Match(actual)
	return itskit.NewMatch(!m.Ok(), n.hedaer, m)
}

func (n notMatcher[T]) Write(ww itsio.Writer) error {
	if err := ww.WriteStringln(n.hedaer); err != nil {
		return err
	}
	in := ww.Indent()
	return n.matcher.Write(in)
}

func (n notMatcher[T]) String() string {
	return itskit.MatcherToString[T](n)
}

// None tests got value does NOT match for all given matchers.
func None[T any](matchers ...itskit.Matcher[T]) itskit.Matcher[T] {
	return noneMathcer[T]{
		label:   itskit.NewLabel("// none of:"),
		matcher: matchers,
	}
}

type noneMathcer[T any] struct {
	label   itskit.Label
	matcher []itskit.Matcher[T]
}

func (none noneMathcer[T]) Match(got T) itskit.Match {
	ok := false
	matches := make([]itskit.Match, len(none.matcher))

	for nth, m := range none.matcher {
		match := m.Match(got)
		ok = ok || match.Ok()
		matches[nth] = match
	}

	return itskit.NewMatch(
		!ok,
		none.label.Fill(got),
		matches...,
	)
}

func (none noneMathcer[T]) Write(w itsio.Writer) error {
	return itsio.WriteBlock(w, none.label.String(), none.matcher)
}

func (none noneMathcer[T]) String() string {
	return itskit.MatcherToString[T](none)
}
