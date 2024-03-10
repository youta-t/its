package location

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/youta-t/its/config"
)

var skipstack = map[uintptr]struct{}{}
var mustack sync.Mutex

func SkipStackSkipping(depth int) (cancel func()) {
	mustack.Lock()
	defer mustack.Unlock()
	if skipstack == nil {
		skipstack = map[uintptr]struct{}{}
	}

	pcs := []uintptr{}

	for d := 0; d <= depth; d += 1 {
		pc, file, line, ok := runtime.Caller(1 + d)
		if !ok {
			return func() {}
		}
		_, _ = file, line
		fnpc := runtime.FuncForPC(pc).Entry()

		skipstack[fnpc] = struct{}{}
		pcs = append(pcs, fnpc)
	}

	return func() {
		mustack.Lock()
		defer mustack.Unlock()
		for _, fnpc := range pcs {
			delete(skipstack, fnpc)
		}
	}
}

var reflectPackage string

func init() {
	goroot := runtime.GOROOT()
	reflectPackage = filepath.Join(goroot, "src", "reflect") + string(os.PathSeparator)

	// when calling func as reflect.Value, like
	//
	//    reflectFn := reflect.ValueOf(someFunc)
	//    reflectFn.Call(...)
	//
	// creates extra call stacks in reflect package.
	//
	// So, to detect invocation locations, it is needed to ignore reflect.
}

func InvokedFrom() Location {
	skip := 2
	for {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			return Location{File: "(unknwon)", Line: -1}
		}
		fn := runtime.FuncForPC(pc)
		fnpc := fn.Entry()
		if strings.HasPrefix(file, reflectPackage) {
			skip += 1
			continue
		}
		if _, ok := skipstack[fnpc]; ok {
			skip += 1
			continue
		}

		file = config.ReplaceFilepath(file)
		return Location{File: file, Line: line}
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
