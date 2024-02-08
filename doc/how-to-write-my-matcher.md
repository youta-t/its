How to write my mathcer?
==============================

In the case of needs of a custom matcher, its has matcher development kit, `github.com/youta-t/its/itskit`.

Simple Matcher
---------------

For just simple matcher, use `itskit.SimpleMatcher`.

In this context, "simple" means...

- matching is determined by short `func(got)bool` function, and
- to build message, only got value and want value are needed.

Now, showing "ApproxEq" matcher:

```go
func ApproxEq(want float64, tolerance float64) its.Matcher[float64] {
	return itskit.SimpleMatcher(
		func(got float64) bool {
			d := want - got
			return -tolerance < d && d < tolerance
		},
		"%f in %f±%f",
		itskit.Got, itskit.Want(want), tolerance,
	)
}
```

It tests that the got value is in the want value ± torerance.

The first argument of `itskit.SimpleMatcher` determines when the match is passed.
It is written `func(got type)bool`.

The second and the third argument is message template and parameters.
The second is just a format string. Nothing special.

Specials are found in the third. Values `its.Got` and `its.Want` appear.
`its.Got` is a placeholder. It will be filled with the got value when matching and prefixed with `/* got */`.
`its.Want` is a decorator. In message, it prefixes `/* want */`.
Others are passed to formatter as it is.

Now, let us try `ApproxEq`.

```go
func TestApproxEq(t *testing.T) {
	ApproxEq(10.0, 3.0).Match(12).OrError(t)    // pass
	ApproxEq(0.12, 0.01).Match(0.1).OrError(t)  // fail
}
```

goes...

```
--- FAIL: TestClose (0.00s)
    /path/to/example_test.go:82: 
        ✘ /* got */ 0.100000 in /* want */ 0.120000±0.010000
```

It works!

### Matcher from Scratch

To write matcher from scratch, you need a type compliant with `its.Matcher[T]`.

`its.Matcher[T]` requires 3 methods

```go
type Matcher[T] interface {
    // Match test value and return result.
    Match(T) itskit.Match

    // String expression of this matcher.
    //
    // You can implement this method as
    //
    // 	func (m youtMatcher) String() string {
    // 		return itskit.MatcherToString(m)
    // 	}
    //
    String() string

    // Write writes its string expression into itsio.Writer
    Write(itsio.Writer) error
}
```

- `Match(T)` returns `itskit.Match` object as result of matching with got value typed `T`.
- `String()` returns string expression of this matcher.
- `Write(itsio.Writer)` writes string expression into `itsio.Writer` (likes `io`).

Needless to say, the most impotant thing is `Match`.

Typically, `Match` is implemented like below:

```go
func (y YourMatcher) Match(got SomeType) itskit.Match {
    var ok bool
    //
    // some routine to determine pass or not. If it pass, set true to ok.
    //
    ok = ...

    return itskit.NewMatch(
        ok,
        itskit.NewLabel(
            "...format string...",
            params...,
        ).Fill(got),
        submatch...,  // if any.
    )
}
```

`itskit.NewLabel` is same as `itskit.SimpleMatcher`'s formatting arguments.
`.Fill` fills values known on matching timing, like the got value.
In `Match`, you can build result messages in a dinamic way.

`itskit.NewMatch` can accept "submatch", corresponds with "submathcer".
Passing submatch, error message of submatches will be concateneted with indented automatically (as you see `its.Slice`, for example).

The method `Write` writes string expression of the matcher itself.
If your matcher has `itskit.Label`, created by `itskit.NewLabel` as a field value, it is easy.

```go
func (y YoutMatcher) Write(w itsio.Writer) error {
    return y.label.Write(w)
}
```

To do that, move `NewLabel` from `Match` method to the factory function.

```go
type YourMatcher struct {
    label itskit.Label
}

func ItsYourMatcher(...) its.Matcher[SomeType] {
    return YourMatcher{
        label: itskit.NewLabel(
            "...format string...",
            params...,
        ),
    }
}

func (y YourMatcher) Match(got SomeType) itskit.Match {
    var ok bool
    //
    // some routine to determine pass or not. If it pass, set true to ok.
    //
    ok = ...

    return itskit.NewMatch(
        ok,
        y.label.Fill(got),
        submatch...,  // if any.
    )
}

func (y YourMatcher) Write(w itsio.Writer) error {
    return y.label.Write(w)
}
```

At the last, `String` method can be implement with utility function.
Do that

```go
func (y YourMatcher) String() string {
	return itskit.MatcherToString(y)
}

```

Now we have walked through of implementation a matcher from scratch.

### Bonus

When creating `itskit.Label`, a plain placeholder, `itskit.Placeholder`, can be used.
This means "need some value, but it will determined at matching".

In contrast of `itskit.Got`, `itskit.Placehodler` has no extra prefix.

Once create `itskit.Label` with `itskit.Placeholder`, like

```go
itskit.newLabel(
    "%d %s %d",
    itskit.Got, itskit.Placeholder, itskit.Want(want)
)
```

Parameter can fill `itskit.Placeholder` by `Label.Fill`'s second or more arguments.

Conclusion
-----------

- For simple condition and simple label, `itskit.SimpleMatcher` is useful.
- For complex matcher, write `its.Matcher[T]` from scratch.
