package itsio

import (
	"bytes"
	"io"

	"github.com/youta-t/its/config"
)

type indentWriter struct {
	indent []byte
	base   io.Writer
}

func newIndentWriter(w io.Writer, indent string) Writer {
	return &indentWriter{indent: []byte(indent), base: w}
}

func (iw *indentWriter) Writeln(p []byte) error {
	ps := bytes.SplitAfter(p, []byte("\n"))
	for _, p := range ps {
		if _, err := iw.base.Write(iw.indent); err != nil {
			return err
		}
		if _, err := iw.base.Write(p); err != nil {
			return err
		}
		if _, err := iw.base.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}

func (iw *indentWriter) WriteStringln(text string) error {
	return iw.Writeln([]byte(text))
}

func (iw *indentWriter) Indent() Writer {
	newIdnent := []byte{}
	newIdnent = append(newIdnent, iw.indent...)
	newIdnent = append(newIdnent, []byte(config.INDENT)...)
	return newIndentWriter(iw.base, string(newIdnent))
}

func WriteBlock[T interface{ Write(Writer) error }](
	w Writer, head string, body []T,
) error {
	if err := w.WriteStringln(head); err != nil {
		return err
	}
	iw := w.Indent()
	for _, b := range body {
		if err := b.Write(iw); err != nil {
			return err
		}
	}
	return nil
}
