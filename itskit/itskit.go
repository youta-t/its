package itskit

import (
	"fmt"

	"github.com/youta-t/its/internal/location"
)

// SkipStack marks callstack not to show in error messge.
//
// # Returns
//
// cancel function. It deletes a mark created by SkipStack.
func SkipStack() (cancel func()) {
	return location.SkipStackSkipping(1)
}

// Location in sourcecode.
type Location struct {
	// Filepath to sourcecode.
	//
	// If it is not detected, "(unknown)" is set.
	File string

	// Line number in File.
	//
	// If it is not detected, -1 is set.
	Line int
}

func (l Location) String() string {
	return fmt.Sprintf("%s:%d", l.File, l.Line)
}

// InvokedFrom detect invoked Location.
//
// InvokedFrom considers skipped call stack by SkipStack().
func InvokedFrom() Location {
	loc := location.InvokedFrom()

	return Location(loc)
}
