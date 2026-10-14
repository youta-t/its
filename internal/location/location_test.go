package location

import (
	"runtime"
	"sync"
	"testing"
)

func skipstackLen() int {
	mustack.RLock()
	defer mustack.RUnlock()

	return len(skipstack)
}

func isSkipping(fn uintptr) (int, bool) {
	mustack.RLock()
	defer mustack.RUnlock()

	n, ok := skipstack[fn]

	return n, ok
}

func detectEntryPC(t *testing.T) uintptr {
	t.Helper()

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("runtime.Caller could not read the calling frame")
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		t.Fatal("runtime.FuncForPC returned nil for the calling frame")
	}

	return fn.Entry()
}

func TestSkipStackSkippingCancelReleasesOnlyItsOwnFrames(t *testing.T) {
	entry := detectEntryPC(t)

	first := SkipStackSkipping(0)
	second := SkipStackSkipping(0)

	if got, ok := isSkipping(entry); !ok || got != 2 {
		t.Fatalf("skipstack holds %d frames for the calling function, want 2", got)
	}

	first()

	if got, ok := isSkipping(entry); !ok || got != 1 {
		t.Errorf("skipstack holds %d frames for the calling function, want 1", got)
	}

	second()

	if _, ok := isSkipping(entry); ok {
		t.Errorf("skipstack still holds frames for the calling function after both cancels ran")
	}
	if got := skipstackLen(); got != 0 {
		t.Errorf("skipstack holds %d frames after every cancel ran", got)
	}
}

func TestSkipStackSkippingCancelTwiceIsNoOp(t *testing.T) {
	entry := detectEntryPC(t)

	other := SkipStackSkipping(0)
	mine := SkipStackSkipping(0)

	if got, ok := isSkipping(entry); !ok || got != 2 {
		t.Fatalf("skipstack holds %d frames for the calling function, want 2", got)
	}

	mine()
	mine()

	if got, ok := isSkipping(entry); !ok || got != 1 {
		t.Errorf("skipstack holds %d frames for the calling function, want 1", got)
	}

	other()

	if _, ok := isSkipping(entry); ok {
		t.Errorf("skipstack still holds frames for the calling function after both cancels ran")
	}
	if got := skipstackLen(); got != 0 {
		t.Errorf("skipstack holds %d frames after every cancel ran", got)
	}
}

// No real call stack can satisfy this depth, so the registration loop stops partway.
const depthBeyondCallStack = 1 << 16

func TestSkipStackSkippingReleasesPartialRegistration(t *testing.T) {
	cancel := SkipStackSkipping(depthBeyondCallStack)

	if skipstackLen() == 0 {
		t.Fatal("no frame was recorded, so this run does not exercise a partial registration")
	}

	cancel()

	if got := skipstackLen(); got != 0 {
		t.Errorf("skipstack holds %d frames after a partial registration was cancelled", got)
	}
}

func TestSkipStackSkippingCancelConcurrently(t *testing.T) {
	entry := detectEntryPC(t)

	other := SkipStackSkipping(0)
	mine := SkipStackSkipping(0)

	if got, ok := isSkipping(entry); !ok || got != 2 {
		t.Fatalf("skipstack holds %d frames for the calling function, want 2", got)
	}

	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mine()
		}()
	}
	wg.Wait()

	if got, ok := isSkipping(entry); !ok || got != 1 {
		t.Errorf("skipstack holds %d frames for the calling function, want 1", got)
	}

	other()

	if got := skipstackLen(); got != 0 {
		t.Errorf("skipstack holds %d frames after every cancel ran", got)
	}
}

func TestInvokedFromWhileFramesAreSkippedConcurrently(t *testing.T) {
	var wg sync.WaitGroup
	for range 4 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range 500 {
				cancel := SkipStackSkipping(0)
				cancel()
			}
		}()
	}

	for range 4 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range 500 {
				InvokedFrom()
			}
		}()
	}

	wg.Wait()
}

// InvokedFrom's walk starts two frames up, which assumes one wrapper frame between itself and the code being located,
// so the tests call it the way itskit does.
func invokedFrom() Location { return InvokedFrom() }

func TestInvokedFromSkipsOnlyMarkedFrames(t *testing.T) {
	t.Run("marked", func(t *testing.T) {
		cancel := SkipStackSkipping(0)
		defer cancel()

		got := invokedFrom()
		_, file, _, _ := runtime.Caller(0)
		if got.File == file {
			t.Errorf("InvokedFrom did not skip the marked frame:\ngot: %s\nwant: not %s\n", got.File, file)
		}
	})

	t.Run("released", func(t *testing.T) {
		cancel := SkipStackSkipping(0)
		cancel()

		_, file, line, _ := runtime.Caller(0)
		got := invokedFrom()

		if got.File != file || got.Line != line+1 {
			t.Errorf("\ngot: %s:%d\nwant: %s:%d\n", got.File, got.Line, file, line)
		}
	})
}
