package itsio

import (
	"fmt"
	"io"

	"github.com/youta-t/its/config"
)

type Writer interface {
	Writeln([]byte) error
	WriteStringln(string) error
	Indent() Writer
}

type writer struct {
	w io.Writer
}

func Wrap(w io.Writer) Writer {
	return &writer{w: w}
}

func (ww *writer) Writeln(p []byte) error {
	if _, err := ww.w.Write(p); err != nil {
		return err
	}
	if _, err := fmt.Fprint(ww.w, "\n"); err != nil {
		return err
	}
	return nil
}

func (ww *writer) WriteStringln(s string) error {
	if _, err := fmt.Fprintln(ww.w, s); err != nil {
		return err
	}
	return nil
}

func (ww *writer) Indent() Writer {
	return newIndentWriter(ww.w, config.INDENT)
}
