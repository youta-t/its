package itskit

import (
	"strings"

	"github.com/youta-t/its/config"
	"github.com/youta-t/its/itskit/itsio"
)

// Match is result of test assertion of Matcher.
type Match struct {
	ok       bool
	message  string
	submatch []Match
}

// create new Match.
//
// # Args
//
// - ok: passing or not.
//
// - msg: message for this match.
//
// - submatch...: submatch of this match.
func NewMatch(ok bool, f string, submatch ...Match) Match {
	return Match{
		ok:       ok,
		message:  f,
		submatch: submatch,
	}
}

// create new Match as ok (passing test).
//
// # Args
//
// - msg: message for this match
//
// - submatch...: submatch of this match
func OK(msg string, submatch ...Match) Match {
	return NewMatch(true, msg, submatch...)
}

// create new Match as not ok (not passed).
//
// # Args
//
// - msg: message for this match
//
// - submatch...: submatch of this match
func NG(msg string, submatch ...Match) Match {
	return NewMatch(false, msg, submatch...)
}

// return true if this match is ok.
func (m Match) Ok() bool {
	return m.ok
}

// Write writes test report into w
func (m Match) Write(w itsio.Writer) error {
	if err := w.WriteStringln(""); err != nil {
		return err
	}
	return m.write(w, !m.ok)
}

func (m Match) write(w itsio.Writer, fail bool) error {

	formatter := config.Pass
	if !m.ok {
		formatter = config.Failed
		if !fail {
			formatter = config.FailedSuppressed
		}
	}

	if err := w.Writeln([]byte(formatter(m.message))); err != nil {
		return err
	}

	if len(m.submatch) <= 0 {
		return nil
	}
	ww := w.Indent()
	for _, s := range m.submatch {
		if err := s.write(ww, !m.ok); err != nil {
			return err
		}
	}

	return nil
}

func (m Match) String() string {
	w := new(strings.Builder)
	m.Write(itsio.Wrap(w))
	s := w.String()
	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return s
}

// call Error only if this Match is not ok.
func (m Match) OrError(t interface{ Error(...any) }) {
	if m.Ok() {
		return
	}
	maybeHelper(t).Helper()
	t.Error(m.String())
}

// call Fatal only if this Match is not ok.
func (m Match) OrFatal(t interface{ Fatal(...any) }) {
	if m.Ok() {
		return
	}
	maybeHelper(t).Helper()
	t.Fatal(m.String())
}

type helper interface {
	Helper()
}

type _fakehelper struct{}

func (fh _fakehelper) Helper() {}

var fakehelper = _fakehelper{}

func maybeHelper(v any) helper {
	if h, ok := v.(helper); ok {
		return h
	}
	return fakehelper
}
