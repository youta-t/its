its  --  A Matcher Library
================================

What it is? -- yes, it's its.
------------------------------

`its` provides value matchers, matcher generator and mock generator.

Install
---------

```
go get github.com/youta-t/its
```

Its requires go1.21+.

Core package: `its`
------------------

`github.com/youta-t/its` package is core package.
Built-in mathcers are here.

For example, `its.EqEq` matcher is for `comparable`.

```go
import (
	"testing"

	"github.com/youta-t/its"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(3, 7)
	its.EqEq(10).   // want is 10
		Match(got). // got
		OrError(t)  // test
}
```

It passes becauase `10 == 3 + 7` is true.

All matchers have example. See them in pkg.go.dev:
https://pkg.go.dev/github.com/youta-t/its and [doc/getting-started](./doc/getting-started.md).

If unmatch is fatal, you can do so.

```go
import (
	"testing"
	"errors"

	"github.com/youta-t/its"
)

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("zero div!")
	}
	return a + b
}

func TestAdd(t *testing.T) {
	_, err := Div(3, 0)
	its.Nil[error](). // want is nil.
		Match(err).   // got
		OrFatal(t)    // test
}
```

### Nice message

If it does not match, it leave nice message.

```go
import "testing"

import "github.com/youta-t/its"

// ...

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(3, 7)
	its.EqEq(got).Match(10).OrError(t)
	its.EqEq(got).Match(11).OrError(t)
}
```

provides,

```
--- FAIL: TestAdd (0.00s)
✘ /* got */ 11 == /* want */ 10		--- @.../example_test.go:33:
```

Error message is tailored for each matchers.

Generally, each messages start with...

- `✔ `: passed matcher
- `✘ `: failed matcher
- `~ `: failed matcher, but not matter

### Composeable

Mathers of its can be composed. For example,

```go
package example_test

import (
	"testing"

	"github.com/youta-t/its"
)

func TestBetween(t *testing.T) {
	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(7).OrError(t)  // pass. 3 < 7 <= 8

	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(8).OrError(t)  // pass. 3 < 8 <= 8

	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(9).OrError(t) // fail! 3 < 9 <= 8
}
```

provides,

```
--- FAIL: TestBetween (0.00s)
✘ // all: (1 ok / 2 matchers)		--- @.../example_test.go:19
    ✔ /* want */ 3 < /* got */ 9		--- @ .../example_test.go:20
    ✘ /* want */ 8 > /* got */ 9		--- @ .../example_test.go:21
```

"A between B" means "greater than A, and lesser than B".

there are also `Some` (require match at least one), `Not` (invert match) and `None` (= `Not(Some(...))` + better message).

Generate Struct Matcher: structer
---------------------------------

`its` has a tool for `//go:generate` to generate matchers of struct, `github.com/youta-t/its/structer`.

You can get matchers for structs, `type T []E` and `type T map[K]E` (alias of slice/map) in and not in your package.

See [structer/README.md](./structer/README.md) for more details.

Generate Mocks: mocker
----------------------

There are another `//go:generate` feature to generate mocks of functions and interfaces, `github.com/youta-t/its/mocker`.

You can mock builders for each `type ... interface` and `type ... func`.

And more, there are "scenario" test feature to check injected functions are called in exact order.

See [mocker/README.md](./mocker/README.md) for more details.

DIY kit included
-----------------

Matcher developmenet kit, `itskit`, is included.

You can create your matcher from scratch in 50 lines or so.
Or, in the simplest case, you need just 10 lines per one matcher.

See [doc/how-to-write-my-matcher.md](./doc/how-to-write-my-matcher.md) to know how to.
