package its_test

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/youta-t/its"
)

func ExampleEqEq() {
	its.EqEq(42).Match(42).OrError(t) // pass. so no messages are output
	its.EqEq(42).Match(49).OrError(t) // fail!

	type MyType struct {
		Foo int
	}

	its.EqEq(MyType{Foo: 42}).Match(MyType{Foo: 24}).OrError(t) // also fail!

	// Output:
	// ✘ /* got */ 49 == /* want */ 42
	//
	// ✘ /* got */ {Foo:24} == /* want */ {Foo:42}
}

func ExampleEqEqPtr() {
	a := 42
	b := 42
	c := 24

	its.EqEqPtr(&a).Match(&a).OrError(t) // pass.
	its.EqEqPtr(&a).Match(&b).OrError(t) // pass. EqEqPtr compairs pointed value.
	its.EqEqPtr(&a).Match(&c).OrError(t) // fail.

	its.EqEqPtr[int](nil).Match(&c).OrError(t)  // fail.
	its.EqEqPtr(&a).Match(nil).OrError(t)       // fail.
	its.EqEqPtr[int](nil).Match(nil).OrError(t) // pass.

	// Output:
	// ✘ /* got */ &(24) == /* want */ &(42)
	//
	// ✘ /* got */ &(24) == /* want */ nil
	//
	// ✘ /* got */ nil == /* want */ &(42)
}

func ExampleGreaterThan() {
	//       <
	its.GreaterThan(10).Match(11).OrError(t) // pass.
	its.GreaterThan(11).Match(10).OrError(t) // fail.
	its.GreaterThan(10).Match(10).OrError(t) // pass.

	its.GreaterThan(1.0).Match(1.1).OrError(t) // pass.
	its.GreaterThan(1.0).Match(1.0).OrError(t) // pass.
	its.GreaterThan(1.1).Match(1.0).OrError(t) // fail.

	its.GreaterThan("aaa").Match("aab").OrError(t) // pass.
	its.GreaterThan("aaa").Match("aaa").OrError(t) // fail.
	its.GreaterThan("aab").Match("aaa").OrError(t) // fail.

	// Output:
	// ✘ /* want */ 11 < /* got */ 10
	//
	// ✘ /* want */ 10 < /* got */ 10
	//
	// ✘ /* want */ 1 < /* got */ 1
	//
	// ✘ /* want */ 1.1 < /* got */ 1
	//
	// ✘ /* want */ aaa < /* got */ aaa
	//
	// ✘ /* want */ aab < /* got */ aaa
}

func ExampleGreaterEq() {
	//      <=
	its.GreaterEq(10).Match(11).OrError(t) // pass.
	its.GreaterEq(11).Match(10).OrError(t) // fail.
	its.GreaterEq(10).Match(10).OrError(t) // fail.

	its.GreaterEq(1.0).Match(1.1).OrError(t) // pass.
	its.GreaterEq(1.0).Match(1.0).OrError(t) // fail.
	its.GreaterEq(1.1).Match(1.0).OrError(t) // fail.

	its.GreaterEq("aaa").Match("aab").OrError(t) // pass.
	its.GreaterEq("aaa").Match("aaa").OrError(t) // pass.
	its.GreaterEq("aab").Match("aaa").OrError(t) // fail.

	// Output:
	// ✘ /* want */ 11 <= /* got */ 10
	//
	// ✘ /* want */ 1.1 <= /* got */ 1
	//
	// ✘ /* want */ aab <= /* got */ aaa
}

func ExampleLesserThan() {
	//       >
	its.LesserThan(10).Match(11).OrError(t) // fail.
	its.LesserThan(11).Match(10).OrError(t) // pass.
	its.LesserThan(10).Match(10).OrError(t) // fail.

	its.LesserThan(1.0).Match(1.1).OrError(t) // fail.
	its.LesserThan(1.0).Match(1.0).OrError(t) // pass.
	its.LesserThan(1.1).Match(1.0).OrError(t) // fail.

	its.LesserThan("aaa").Match("aab").OrError(t) // fail.
	its.LesserThan("aaa").Match("aaa").OrError(t) // fail.
	its.LesserThan("aab").Match("aaa").OrError(t) // pass.

	// Output:
	// ✘ /* want */ 10 > /* got */ 11
	//
	// ✘ /* want */ 10 > /* got */ 10
	//
	// ✘ /* want */ 1 > /* got */ 1.1
	//
	// ✘ /* want */ 1 > /* got */ 1
	//
	// ✘ /* want */ aaa > /* got */ aab
	//
	// ✘ /* want */ aaa > /* got */ aaa
}

func ExampleLesserEq() {
	//      >=
	its.LesserEq(10).Match(11).OrError(t) // pass.
	its.LesserEq(11).Match(10).OrError(t) // fail.
	its.LesserEq(10).Match(10).OrError(t) // fail.

	its.LesserEq(1.0).Match(1.1).OrError(t) // pass.
	its.LesserEq(1.0).Match(1.0).OrError(t) // fail.
	its.LesserEq(1.1).Match(1.0).OrError(t) // fail.

	its.LesserEq("aaa").Match("aab").OrError(t) // pass.
	its.LesserEq("aaa").Match("aaa").OrError(t) // pass.
	its.LesserEq("aab").Match("aaa").OrError(t) // fail.

	// Output:
	// ✘ /* want */ 10 >= /* got */ 11
	//
	// ✘ /* want */ 1 >= /* got */ 1.1
	//
	// ✘ /* want */ aaa >= /* got */ aab
}

func ExampleBefore() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.Before(t1).Match(t2).OrError(t)
	its.Before(t1).Match(t1).OrError(t)
	its.Before(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* got */ 2023-10-11 12:18:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)
}

func ExampleAfter() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.After(t1).Match(t2).OrError(t)
	its.After(t1).Match(t1).OrError(t)
	its.After(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:13:14 +0000 +0000)
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:18:14 +0000 +0000)
}

func ExampleEqual() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.Equal(t1).Match(t2).OrError(t)
	its.Equal(t1).Match(t1).OrError(t)
	its.Equal(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* want */ 2023-10-11 12:13:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:18:14 +0000 +0000)
	//
	// ✘ (/* want */ 2023-10-11 12:18:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:13:14 +0000 +0000)
}

func ExampleEquivWith() {
	// testing equiveness on stringified value
	its.EquivWith(
		42, func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).Match("42").OrError(t)

	its.EquivWith(
		42, func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).Match("40").OrError(t)

	// Output:
	// ✘ (/* want */ 42) equiv. (/* got */ 40)
}

func ExampleError() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e1).Match(e2).OrError(t)
	its.Error(e1).Match(e1).OrError(t)
	its.Error(e2).Match(e1).OrError(t)

	// Output:
	// ✘ errors.Is(/* got */ error, /* want */ wrapped: error)
}

type CustomError string

func (c CustomError) Error() string {
	return string(c)
}

func ExampleErrorAs() {
	customErr := CustomError("custom error")
	otherErr := errors.New("error by error.New")

	its.ErrorAs[CustomError]().Match(customErr).OrError(t)
	its.ErrorAs[CustomError]().Match(otherErr).OrError(t)

	// Output:
	// ✘ want := new(its_test.CustomError); errors.As(/* got */ error by error.New, want)
}

func ExampleAlways() {
	its.Always[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
}

func ExampleNever() {
	its.Never[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
	// ✘ (never pass)
}

func ExampleStringHavingPrefix() {
	its.StringHavingPrefix[string]("abc").Match("abcde").OrError(t)
	its.StringHavingPrefix[string]("abc").Match("adcbe").OrError(t)

	type MyString string
	its.StringHavingPrefix[MyString]("abc").Match("abcde").OrError(t)
	its.StringHavingPrefix[MyString]("abc").Match("adcbe").OrError(t)

	// Output:
	// ✘ strings.HasPrefix(/* got */ "adcbe", /* want */ "abc")
	//
	// ✘ strings.HasPrefix(/* got */ "adcbe", /* want */ "abc")
}

func ExampleStringHavingSuffix() {
	its.StringHavingSuffix[string]("cde").Match("abcde").OrError(t)
	its.StringHavingSuffix[string]("cde").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.HasSuffix(/* got */ "adcbe", /* want */ "cde")
}

func ExampleStringContaining() {
	its.StringContaining[string]("bcd").Match("abcde").OrError(t)
	its.StringContaining[string]("bcd").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.Contains(/* got */ "adcbe", /* want */ "bcd")
}

func ExampleStringEqualFold() {
	its.StringEqualFold[string]("abc").Match("abc").OrError(t)
	its.StringEqualFold[string]("aBc").Match("AbC").OrError(t)
	its.StringEqualFold[string]("abc").Match("αβγ").OrError(t)
	// Output:
	// ✘ strings.EqualFold(/* got */ "αβγ", /* want */ "abc")
}

func ExampleBytesEqual() {
	its.BytesEqual([]byte("abc")).Match([]byte("abc")).OrError(t)
	its.BytesEqual([]byte("abc")).Match([]byte("acb")).OrError(t)
	// Output:
	// ✘ bytes.Equal(/* got */ []byte{0x61, 0x63, 0x62}, /* want */ []byte{0x61, 0x62, 0x63})
}

func ExampleBytesHavingPrefix() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("abcde")).OrError(t)
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasPrefix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x61, 0x62, 0x63})
}

func ExampleBytesHavingSuffix() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("abcde")).OrError(t)
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasSuffix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x63, 0x64, 0x65})
}

func ExampleBytesContaining() {
	its.BytesContaining([]byte("bcd")).Match([]byte("abcde")).OrError(t)
	its.BytesContaining([]byte("bcd")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.Contains(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x62, 0x63, 0x64})
}

func ExampleNaN() {
	its.NaN().Match(math.NaN()).OrError(t)
	its.NaN().Match(42).OrError(t)
	// Output:
	// ✘ math.IsNaN(/* got */ 42.000000)
}

func ExampleInf() {
	its.Inf().Match(math.Inf(1)).OrError(t)
	its.Inf().Match(math.Inf(-1)).OrError(t)
	its.Inf().Match(0).OrError(t)
	// Output:
	// ✘ math.IsInf(/* got */ 0.000000, 0)
}

func ExampleClosed() {
	ch1 := make(chan int, 1)
	close(ch1)
	its.Closed[int]().Match(ch1).OrError(t)
	ch2 := make(chan string, 1)
	its.Closed[string]().Match(ch2).OrError(t)
	// Output:
	// ✘ chan string is not closed.
}

func ExampleType() {
	its.Type[string]().Match("text value").OrError(t)
	its.Type[int]().Match("text value").OrError(t)
	type T struct {
		Foo int
	}
	its.Type[int]().Match(T{Foo: 42}).OrError(t)
	// Output:
	// ✘ /* got */ text value is a int
	//
	// ✘ /* got */ {Foo:42} is a int
}
