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

Trailing `.OrError(t)` is reporting. It calles `t.Error`, if and only if the match is failed.

Run the test above, you will show like this:

```
--- FAIL: TestExample (0.00s)
✘ /* got */ 12 == /* want */ 42		/path/to/example_test.go:10
```

`✘` tells you "this match is failed!"

There are more matchers.

- `its.Equal`: match with `want.Equal(got)` method, for like `time.Time`
- `its.DeepEqual`: match with `reflect.DeepEqual(got, want)`
- `its.StringHavingPrefix`: match with `strings.HasPrefix(got, want)`
- `its.GreaterEq`: match with `want <= got`
- `its.LesserThan`: match with `want > got`
- `its.ClosedChan`: tests the channel is closed
- `its.Always` and `its.Never`
- ...and more!

For all matchers, see https://pkg.go.dev/github.com/youta-t/its .

If you want to fail-fast, you call `Match(...).OrFatal(t)`.
`.OrFatal` likes `OrError`, but calling `t.Fatal`.

```go
func TestExample(t *testing.T) {
    its.EqEq(42).Match(12).OrFatal(t)
    its.EqEq(100).Match(100).OrError(t)  // does not reach here
}
```

If you need "match or not" in a value, `Match(...).Ok()`. This returns `true` if passing.

Slice and Map
--------------

Of cource its has matchers for slice and map.

### Example: `its.Slice`

`its.Slice` matches with slice and `[]its.Macther`.

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
✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ /path/to/example_test.go:21
    ✔ /* want */ 3 >= /* got */ 1		--- @ /path/to/example_test.go:22
    ✔ /* got */ 4 == /* want */ 4		--- @ /path/to/example_test.go:23
    ✘ - /* want */ 5 <= /* got */ ??		--- @ /path/to/example_test.go:24
    ✘ + /* got */ 4
```

It shows matching report with diff.
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
✘ map[string]string{... ( keys: /* got */ 4, /* want */ 4; +3, -3 )		 --- @ /path/to/example_test.go:38
    ✔ body:
        ✔ (/* want */ ^(t-shirt|jacket)$).MatchString(/* got */ "jacket")		 --- @ /path/to/example_test.go:40
    ✘ foot: (not in got)
        ✘ /* got */ ?? == /* want */ sneaker		--- @ /path/to/example_test.go:42
    ✘ hand: (not in want)
        ✘ /* got */ white glove, /* want */ ??
    ✘ head:
        ✘ /* got */ hat == /* want */ cap		--- @ /path/to/example_test.go:39
    ✘ leg:
        ✘ /* got */ slacks == /* want */ jeans		--- @ /path/to/example_test:41
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
✘ // all: (1 ok / 2 matchers)		--- @ /path/to/example_test.go:4
    ✔ strings.Contains(/* got */ "Elmer and the Dragon", /* want */ "Dragon")		--- @ /path/to/example_test.go:5
    ✘ strings.Contains(/* got */ "Elmer and the Dragon", /* want */ "Dungeon")		--- @ /path/to/example_test.go:6
```

Stateful Matcher
------------------

Almost all matchers are "stateless", means they do not remember got values in the past.

But, some mathcers are "stateful".
They remembers got values and deteremines match or not with the history.

> [!WARNING]
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
✘ // monotonic		--- @ /path/to/example_test.go:64
    ✔ (always pass)		--- @ /path/to/example_test.go:65
    ✔ /* want */ abc < /* got */ def		--- @ /path/to/example_test.go:66
    ✘ /* want */ def < /* got */ ddf		--- @ /path/to/example_test.go:67
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
✘ //do not match with values have been gotten%%--- @ /path/to/example_test.go:56
    ✔ (always pass)		--- @ /path/to/example_test.go:58
    ✔ // none of:		--- @ /path/to/example_test.go:59
        ~ /* got */ id: bbb == /* want */ id: aaa		--- @ /path/to/example_test.go:58
    ✔ // none of:		--- @ /path/to/example_test.go:60
        ~ /* got */ id: ccc == /* want */ id: aaa		--- @ /path/to/example_test.go:58
        ~ /* got */ id: ccc == /* want */ id: bbb		--- @ /path/to/example_test.go:59
    ✘ // none of:		--- @ /path/to/example_test.go:61
        ✘ /* got */ id: bbb == /* want */ id: aaa		--- @ /path/to/example_test.go:58
        ✔ /* got */ id: bbb == /* want */ id: bbb		--- @ /path/to/example_test.go:59
        ✘ /* got */ id: bbb == /* want */ id: ccc		--- @ /path/to/example_test.go:60
```

Conclusion
-----------

- its has many `its.Matcher[T]` and they works as `matcher.Match(got_value).OrError(t)`.
- matchers can be combined with `All`, `Some`, `None` or negated by `Not`.
- some matchers are stateful.
