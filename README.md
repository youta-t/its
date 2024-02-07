its  --  A Matcher Library
================================

What it is? -- yes, it's its.
------------------------------

`its` provides value matchers. For example, its has `its.EqEq` for `comparable`.

```go
import "testing"

import "github.com/youta-t/its"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    actual := Add(3, 7)
    its.EqEq(10).Match(actual).OrError(t)
}
```

It passes becauase `10 == 3 + 7` is true, as you see.

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
    .../example_test.go:33: ✘ /* got */ 11 == /* want */ 10
```

Error message is tailored for each matchers. See examples to meet more matchers!

### Composeable

Mathers of its can be composed. For example,

```go
package example_test

import (
	"testing"

	"github.com/youta-t/its"
)

// ...

func TestBetween(t *testing.T) {
	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(7).OrError(t)

	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(8).OrError(t)

	its.All(
		its.GreaterThan(3),
		its.LesserEq(8),
	).Match(9).OrError(t)
}
```

provides,

```
--- FAIL: TestBetween (0.00s)
    .../example_test.go:20:
        ✘ // all: (1 ok / 2 matchers)
            ✔ /* want */ 3 < /* got */ 9
            ✘ /* want */ 8 > /* got */ 9
```

"A between B" means "greater than A, and lesser than B", as you know.

Not only `All`, there are `Some` (require match at least one) and `Not` (invert match).

### Generate Struct Matcher: structer

its has a tool for `//go:generate` to generate matchers of struct, `github.com/youta-t/its/structer`.

For example, given a snippet below,

```go
//go:generate go run github.com/youta-t/its/structer -dest gen

package ...

type MyStruct struct {
	Name   string
	Value int
}
```

Structer generates a file at `./gen/${the filename with go:generate}.go`.
And the content will be:

Notice of `MyStructSpec` and `ItsMyStruct` type, or more generally, `T_Spec` and `Its_T`.

- `Its_T` type is the matcher for struct `T`.
    - It is compliant with `itskit.Matcher[T]`.
- `T_Spec` type is container of matchers for each field of `T`

Use them as `gen.Its_T(T_Spec{ ... })`, like:

```go
	gen.ItsMyStruct1(gen.MyStructSpec{
		Name: its.StringHavingPrefix[string]("its"),
		Values: its.GreaterThan(3),
	}).
		Match(types.MyStruct{
			Name: "its a matching library",
			Values: 300,
		}).
		OrError(t)
```

### Easy DIY

Matcher developmenet kit, `itskit`, is bundled.

#### Simple matcher

Using `itskit.SimpleMatcher`, you can jot down your matcher.

```go
package example_test

import (
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/itskit"
)

// ...

func TestEven(t *testing.T) {
	itsEven := itskit.SimpleMatcher(
        // when got is want?
        func(got int) bool { return got%2 == 0 },

        // format of error message
        "%d is even",
        itskit.Got, // ...and placeholder
	)

	itsEven.Match(7).OrError(t)
	itsEven.Match(8).OrError(t)
}
```


The example above provides log like below:

```
--- FAIL: TestEven (0.00s)
    .../example_test.go:43: 
        ✘ /* got */ 7 is even
```

#### More matcher

When you need more complex matchers, implement your Matcher.

Only 3 methods are needed.

- `Match[T](got T) itskit.Match` : determine whether the got value matchs or not.
- `Write(itsio.Writer) error` : string expression of the matcher itself.
- `String() string` : same above, but return string directly.

For example: this is full implementation of `EqEqPtr` matcher.

```go
type eqeqPtrMatcher[T comparable] struct {
	label itskit.Label
	want  *T
}

func ptrLabel[T any](v *T) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("&(%+v)", *v)
}

func (epm eqeqPtrMatcher[T]) Match(got *T) itskit.Match {
	ok := false

	if got == nil || epm.want == nil {
		ok = got == nil && epm.want == nil
	} else {
		ok = *got == *epm.want
	}

	return itskit.NewMatch(
		ok,
		epm.label.Fill(ptrLabel(got)),
	)
}

func (epm eqeqPtrMatcher[T]) Write(w itsio.Writer) error {
	return epm.label.Write(w)
}

func (epm eqeqPtrMatcher[T]) String() string {
	return itskit.MatcherToString[*T](epm)
}

// EqEqPtr tests of pointer for comparable with
//
//	(want == got) || (*want == *got)
func EqEqPtr[T comparable](want *T) its.Matcher[*T] {
	return eqeqPtrMatcher[T]{
		label: itskit.NewLabel(
			"%+v == %+v",
			itskit.Got, itskit.Want(ptrLabel(want)),
		),
		want: want,
	}
}
```
