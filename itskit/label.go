package itskit

import (
	"fmt"
	"io"
	"strings"

	"github.com/youta-t/its/itskit/itsio"
)

// Label is label of Matcher.
//
// This message is printed out when Matcher detect assertion error.
type Label struct {
	template string
	params   []any
}

// NewLabel create new Label for Matcher.
//
// # Args
//
// - template: template of label message
// You can use standard format syntax, for examle, %s, %v or so.
//
// - params...: parameters fill template.
//
// # Params and Special Value
//
// Params supports some special placeholding values.
//
// You can them in params, and Fill() later.
//
// - Placeholder: Placeholder is a placeholder for a value not determined yet.
//
// - Got: Got is a placeholder filled by got (actual) value.
// If Got presents whereever, it will be filled by got value.
// On printing, value will be prefixed with "/* got */"
//
// And, where are value decorator.
//
// - Want(value): return wrapped value.
// When Want-wrapped value is printed, it will be prefixed with "/* want */"
//
// # Example
//
// Creating a label with 1 placeholder and "want" value.
//
//	l := NewLabel(
//		"%+v == %+v",
//		Got, Want(42),
//	)
//
// and l.Stirng() comes
//
//	"/* got */ ?? == /* want */ 42"
//
// Placeholder Got is stubbed by "??" yet.
//
// Then, fill it like l.Fill(99). You get
//
//	"/* got */ 99 == /* want */ 42"
func NewLabel(template string, params ...any) Label {
	return Label{
		template: template,
		params:   params,
	}
}

// NewLabelWithLocation like NewLabel, but appends where this function is called.
func NewLabelWithLocation(template string, params ...any) Label {
	cancel := SkipStack()
	defer cancel()

	p := make([]any, len(params))
	copy(p, params)

	from := InvokedFrom()
	p = append(p, from)
	return NewLabel(
		template+"\t\t--- @ %s",
		p...,
	)
}

func (c Label) Write(w itsio.Writer) error {
	if err := w.WriteStringln(fmt.Sprintf(c.template, c.params...)); err != nil {
		return err
	}
	return nil
}

func (c Label) String() string {
	sb := new(strings.Builder)
	c.Write(itsio.Wrap(sb))

	s := sb.String()
	if last := len(s) - 1; s[last] == '\n' {
		s = s[:last]
	}
	return s
}

// Fill fills placeholders in the label and build into string value.
//
// # Args
//
// - got: got value in test assertion
// If there are no values are meaningful as actual, you can pass Missing.
//
// - params...: other parameters for the label.
// They fills placeholders in order (except Got placeholder).
func (c Label) Fill(got any, params ...any) string {
	p := []any{}
	for _, item := range c.params {
		switch x := item.(type) {
		case interface{ FillByGot(any) any }:
			p = append(p, x.FillByGot(got))
		case interface{ Fill(any) any }:
			if 0 < len(params) {
				p = append(p, x.Fill(params[0]))
				params = params[1:]
			}
		default:
			p = append(p, item)
		}
	}

	return fmt.Sprintf(c.template, p...)
}

type placeholder struct{}

var Placeholder = placeholder{}

func (placeholder) Format(f fmt.State, verb rune) {
	io.WriteString(f, "??")
}

func (placeholder) Fill(v any) any {
	return v
}

type want struct {
	Value any
}

// Want marks the value is the test expectation.
func Want(value any) want {
	return want{value}
}

func (a want) Format(s fmt.State, verb rune) {
	io.WriteString(s, "/* want */ ")
	fmt.Fprintf(s, fmt.FormatString(s, verb), a.Value)
}

type emptyGot struct{}

func (emptyGot) Format(f fmt.State, verb rune) {
	io.WriteString(f, "/* got */ ??")
}

func (emptyGot) FillByGot(v any) any {
	return filledGot{Value: v}
}

// Got is a placeholder will be filled with the actual value.
var Got = emptyGot{}

type filledGot struct {
	Value any
}

func (a filledGot) Format(f fmt.State, verb rune) {
	io.WriteString(f, "/* got */ ")
	fmt.Fprintf(f, fmt.FormatString(f, verb), a.Value)
}

type missing struct{}

// Missing represents "some value is expected, but not given".
var Missing = missing{}

func (missing) Format(s fmt.State, verb rune) {
	fmt.Fprintf(s, "??")
}
