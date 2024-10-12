package its

import (
	"context"

	"github.com/youta-t/its/itskit"
)

// Context is alias of `its.Always[context.Context]()`
func Context() Matcher[context.Context] {
	// define as func not to be replaced.
	cancel := itskit.SkipStack()
	defer cancel()

	return Always[context.Context]()
}
