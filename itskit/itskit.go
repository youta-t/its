package itskit

import (
	"fmt"
	"runtime"
	"sync"
)

var skipstack = map[uintptr]struct{}{}
var mustack sync.Mutex

// SkipStack marks callstack not to show in error messge.
//
// # Returns
//
// cancel function. It deletes a mark created by SkipStack.
func SkipStack() (cancel func()) {
	mustack.Lock()
	defer mustack.Unlock()
	if skipstack == nil {
		skipstack = map[uintptr]struct{}{}
	}

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return func() {}
	}

	fnpc := runtime.FuncForPC(pc).Entry()

	skipstack[fnpc] = struct{}{}

	return func() {
		mustack.Lock()
		defer mustack.Unlock()
		delete(skipstack, fnpc)
	}
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
	skip := 2
	for {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			return Location{File: "(unknwon)", Line: -1}
		}
		fn := runtime.FuncForPC(pc)
		fnpc := fn.Entry()
		if _, ok := skipstack[fnpc]; ok {
			skip += 1
			continue
		}

		return Location{File: file, Line: line}
	}
}
