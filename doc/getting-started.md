Getting Started
================

`github.com/youta-t/its` has a suite of matchers.
Each matchers is typed `its.Matcher[T]`, here `T` is type of the value to be tested.

For example, the simplest matcher `its.EqEq[T]` tests "want" value and "got" value with `==` operator.

```go
func TestExample(t *testing.T) {
    its.EqEq(42).Match(12).OrError(t)
}
```

`42` is the want value, and `12` is got value.

Trailing `.OrError(t)` is reporting.
It calles `t.Error`, if and only if the match is failed,

Run the test above, you will show like this:

```
--- FAIL: TestExample (0.00s)
    /path/to/example_test.go:10: 
        ✘ /* got */ 12 == /* want */ 42
```

`✘` tells you "this match is failed!"

There are more matchers.

- `its.Equal`: match with `want.Equal(got)` method, like `time.Time`
- `its.StringHavingPrefix`: match with `strings.HasPrefix(got, want)`
- `its.GreaterEq`: match with `want <= got`
- `its.LesserThan`: match with `want > got`
- `its.ClosedChan`: tests the channel is closed
- `its.Always` and `its.Never`
- ...and more!

For all matchers, see https://pkg.go.dev/github.com/youta-t/its .

Slice and Map
--------------

Of cource its has matchers for slice and map.

### Example: `its.Slice`

`its.Slice` matches with slice and `Macther[slice]`s.

See example:

```go
func TestSlice(t *testing.T){
    its.Slice(
        its.LesserEq(3),
        its.EqEq(4),
        its.GreaterEq(5),
    ).
    Match([]int{1, 4, 5}).
    OrError(t)
}
```

`its.Slice` holds matchers for each elements of got value.

This match will pass, because

- the first element in got is `3 >= 1`
- the second is `4 == 4`, and
- the last is `5 <= 5`

To see error message, let the test be broken.

```go
func TestSlice(t *testing.T) {

	its.Slice(
		its.LesserEq(3),
		its.EqEq(4),
		its.GreaterEq(5),
	).
		Match([]int{1, 4, 4}).
		OrError(t)
}
```

And, you get error message

```
--- FAIL: TestSlice (0.00s)
    /path/to/example_test.go:21: 
        ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
            ✔ /* want */ 3 >= /* got */ 1
            ✔ /* got */ 4 == /* want */ 4
            ✘ - /* want */ 5 <= /* got */ ??
            ✘ + /* got */ 4
```

It shows match report with diff.
Extra items in got values are prefixed with `+`, and extra mathcers are prefixed with `-`.

### Example: `its.Map`

Map mathers are in same vives.

```go
func TestMap(t *testing.T) {
	its.Map(its.MapSpec[string, string]{
		"head": its.EqEq("cap"),
		"body": its.MatchString(regexp.MustCompile(`^(t-shirt|jacket)$`)),
		"leg":  its.EqEq("jeans"),
		"foot": its.EqEq("sneaker"),
	}).
		Match(map[string]string{
			"head": "hat",
			"body": "jacket",
			"leg":  "slacks",
			"hand": "white glove",
		}).
		OrError(t)
}

```

will shows you

```
--- FAIL: TestMap (0.00s)
    /path/to/example_test.go:38: 
        ✘ map[string]string{... ( keys: /* got */ 4, /* want */ 4; +3, -3 )
            ✔ body:
                ✔ (/* want */ ^(t-shirt|jacket)$).MatchString(/* got */ "jacket")
            ✘ foot: (not in got)
                ✘ /* got */ ?? == /* want */ sneaker
            ✘ hand: (not in want)
                ✘ /* got */ white glove, /* want */ ??
            ✘ head:
                ✘ /* got */ hat == /* want */ cap
            ✘ leg:
                ✘ /* got */ slacks == /* want */ jeans
```

Matcher Combinator
-------------------

Its has logical macther combinator, `All`, `Some`, `Not` and `None`.
They are mathcers using matchers, like slice and map mathcers.

- `All`: mathes if all submatchers match.
- `Some`: matches if at lease one of submatchers matches.
- `Not`: matches if submatcher does not match.
- `None`: mathes if any submatchers do not match.

### Example: `its.All`

```go
func TestAll(t *testing.T) {
	its.All(
		its.StringContaining("Dragon"),
		its.StringContaining("Dungeon"),
	).
		Match("Dungeons and Dragons").
		OrError(t)
}
```

This is okay. But...

```go
func TestAll(t *testing.T) {
	its.All(
		its.StringContaining("Dragon"),
		its.StringContaining("Dungeon"),
	).
		Match("Elmer and the Dragon").
		OrError(t)
}
```

does not have "Dungeon", so failed.

```
--- FAIL: TestAll (0.00s)
    /path/to/example_test.go:47: 
        ✘ // all: (1 ok / 2 matchers)
            ✔ strings.Contains(/* got */ "Elmer and the Dragon", /* want */ "Dragon")
            ✘ strings.Contains(/* got */ "Elmer and the Dragon", /* want */ "Dungeon")
```

Stateful Matcher
------------------

Almost all matchers are "stateless", means they do not remember got values in the past.

But, some mathcers are "stateful".
They remembers got values and deteremines match or not with the history.

> **Warning**
>
> DO NOT use stateful mathers in slice-related matcher, like `its.Slice` or `its.SliceContaining`.
>
> Because slice-related matcher tries for each submatchers repeatedly to make diff.
> It will lead unclear result.

`its.Monotonic` and `its.Singuler` are stateful matcher.

### `its.Monotonic`

`its.Monotonic` matches if new got value matches with the last got value.

```go
func TestMonotonic(t *testing.T) {
	itsLexicalOrder := its.Monotonic(its.GreaterThan[string])
	itsLexicalOrder.Match("abc").OrError(t)
	itsLexicalOrder.Match("def").OrError(t)
	itsLexicalOrder.Match("ddf").OrError(t)
	itsLexicalOrder.Match("dee").OrError(t)
}
```

Note that `its.Monotonic(its.GreaterThan[string])`, not `its.Monotonic(its.GreaterThan[string]())`.
`Monotonic` receives a matcher factory to create matchers for each match.

It goes failed, and show you

```
--- FAIL: TestMonotonic (0.00s)
    /path/to/example_test.go:64: 
        ✘ // monotonic
            ✔ (always pass)
            ✔ /* want */ abc < /* got */ def
            ✘ /* want */ def < /* got */ ddf
```

This error message is shown because the third match is failed. So, match history is at then.

### `its.Singuler`

`its.Singuler` matches if new got value does not match with past got values.

```go
func TestSinguler(t *testing.T) {
	itsUniqueId := its.Singuler(its.EqEq[string])

	itsUniqueId.Match("id: aaa").OrError(t)
	itsUniqueId.Match("id: bbb").OrError(t)
	itsUniqueId.Match("id: ccc").OrError(t)
	itsUniqueId.Match("id: bbb").OrError(t) // duplicated!
	itsUniqueId.Match("id: ddd").OrError(t)
}
```

```
--- FAIL: TestSinguler (0.00s)
    /path/to/example_test.go:56: 
        ✘ //do not match with values have been gotten
            ✔ (always pass)
            ✔ // none of:
                ~ /* got */ id: bbb == /* want */ id: aaa
            ✔ // none of:
                ~ /* got */ id: ccc == /* want */ id: aaa
                ~ /* got */ id: ccc == /* want */ id: bbb
            ✘ // none of:
                ✘ /* got */ id: bbb == /* want */ id: aaa
                ✔ /* got */ id: bbb == /* want */ id: bbb
                ✘ /* got */ id: bbb == /* want */ id: ccc
```

Conclusion
-----------

- its has many `its.Matcher[T]`, it works as `matcher.Match(got_value).OrError(t)`.
- matchers can be combined with `All`, `Some`, `None` or negated by `Not`.
- some matchers are stateful.
