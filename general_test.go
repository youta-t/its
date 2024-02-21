package its_test

import (
	"errors"
	"fmt"
	"math"
	"regexp"
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
	// ✘ /* got */ 49 == /* want */ 42		--- @ ./general_test.go:15
	//
	// ✘ /* got */ {Foo:24} == /* want */ {Foo:42}		--- @ ./general_test.go:21
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
	// ✘ /* got */ &(24) == /* want */ &(42)		--- @ ./general_test.go:36
	//
	// ✘ /* got */ &(24) == /* want */ nil		--- @ ./general_test.go:38
	//
	// ✘ /* got */ nil == /* want */ &(42)		--- @ ./general_test.go:39
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
	// ✘ /* want */ 11 < /* got */ 10		--- @ ./general_test.go:53
	//
	// ✘ /* want */ 10 < /* got */ 10		--- @ ./general_test.go:54
	//
	// ✘ /* want */ 1 < /* got */ 1		--- @ ./general_test.go:57
	//
	// ✘ /* want */ 1.1 < /* got */ 1		--- @ ./general_test.go:58
	//
	// ✘ /* want */ aaa < /* got */ aaa		--- @ ./general_test.go:61
	//
	// ✘ /* want */ aab < /* got */ aaa		--- @ ./general_test.go:62
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
	// ✘ /* want */ 11 <= /* got */ 10		--- @ ./general_test.go:81
	//
	// ✘ /* want */ 1.1 <= /* got */ 1		--- @ ./general_test.go:86
	//
	// ✘ /* want */ aab <= /* got */ aaa		--- @ ./general_test.go:90
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
	// ✘ /* want */ 10 > /* got */ 11		--- @ ./general_test.go:102
	//
	// ✘ /* want */ 10 > /* got */ 10		--- @ ./general_test.go:104
	//
	// ✘ /* want */ 1 > /* got */ 1.1		--- @ ./general_test.go:106
	//
	// ✘ /* want */ 1 > /* got */ 1		--- @ ./general_test.go:107
	//
	// ✘ /* want */ aaa > /* got */ aab		--- @ ./general_test.go:110
	//
	// ✘ /* want */ aaa > /* got */ aaa		--- @ ./general_test.go:111
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
	// ✘ /* want */ 10 >= /* got */ 11		--- @ ./general_test.go:130
	//
	// ✘ /* want */ 1 >= /* got */ 1.1		--- @ ./general_test.go:134
	//
	// ✘ /* want */ aaa >= /* got */ aab		--- @ ./general_test.go:138
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
	// ✘ (/* got */ 2023-10-11 12:18:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:160
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:161
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
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:181
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./general_test.go:182
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
	// ✘ (/* want */ 2023-10-11 12:13:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./general_test.go:200
	//
	// ✘ (/* want */ 2023-10-11 12:18:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:202
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
	// ✘ (/* want */ 42) equiv. (/* got */ 40)		--- @ ./general_test.go:216
}

func ExampleError() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e1).Match(e2).OrError(t)
	its.Error(e1).Match(e1).OrError(t)
	its.Error(e2).Match(e1).OrError(t)

	// Output:
	// ✘ errors.Is(/* got */ error, /* want */ wrapped: error)		--- @ ./general_test.go:230
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
	// ✘ want := new(its_test.CustomError); errors.As(/* got */ error by error.New, want)		--- @ ./general_test.go:247
}

func ExampleAlways() {
	its.Always[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
}

func ExampleNever() {
	its.Never[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
	// ✘ (never pass)		--- @ ./general_test.go:259
}

func ExampleStringHavingPrefix() {
	its.StringHavingPrefix("abc").Match("abcde").OrError(t)
	its.StringHavingPrefix("abc").Match("adcbe").OrError(t)

	// Output:
	// ✘ strings.HasPrefix(/* got */ "adcbe", /* want */ "abc")		--- @ ./general_test.go:266
}

func ExampleStringHavingSuffix() {
	its.StringHavingSuffix("cde").Match("abcde").OrError(t)
	its.StringHavingSuffix("cde").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.HasSuffix(/* got */ "adcbe", /* want */ "cde")		--- @ ./general_test.go:274
}

func ExampleStringContaining() {
	its.StringContaining("bcd").Match("abcde").OrError(t)
	its.StringContaining("bcd").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.Contains(/* got */ "adcbe", /* want */ "bcd")		--- @ ./general_test.go:281
}

func ExampleStringEqualFold() {
	its.StringEqualFold("abc").Match("abc").OrError(t)
	its.StringEqualFold("aBc").Match("AbC").OrError(t)
	its.StringEqualFold("abc").Match("αβγ").OrError(t)
	// Output:
	// ✘ strings.EqualFold(/* got */ "αβγ", /* want */ "abc")		--- @ ./general_test.go:289
}

func ExampleBytesEqual() {
	its.BytesEqual([]byte("abc")).Match([]byte("abc")).OrError(t)
	its.BytesEqual([]byte("abc")).Match([]byte("acb")).OrError(t)
	// Output:
	// ✘ bytes.Equal(/* got */ []byte{0x61, 0x63, 0x62}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./general_test.go:296
}

func ExampleBytesHavingPrefix() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("abcde")).OrError(t)
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasPrefix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./general_test.go:303
}

func ExampleBytesHavingSuffix() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("abcde")).OrError(t)
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasSuffix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x63, 0x64, 0x65})		--- @ ./general_test.go:310
}

func ExampleBytesContaining() {
	its.BytesContaining([]byte("bcd")).Match([]byte("abcde")).OrError(t)
	its.BytesContaining([]byte("bcd")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.Contains(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x62, 0x63, 0x64})		--- @ ./general_test.go:317
}

func ExampleNaN() {
	its.NaN().Match(math.NaN()).OrError(t)
	its.NaN().Match(42).OrError(t)
	// Output:
	// ✘ math.IsNaN(/* got */ 42.000000)		--- @ ./general_test.go:324
}

func ExampleInf() {
	its.Inf().Match(math.Inf(1)).OrError(t)
	its.Inf().Match(math.Inf(-1)).OrError(t)
	its.Inf().Match(0).OrError(t)
	// Output:
	// ✘ math.IsInf(/* got */ 0.000000, 0)		--- @ ./general_test.go:332
}

func ExampleClosedChan() {
	ch1 := make(chan int, 1)
	close(ch1)
	its.ClosedChan[int]().Match(ch1).OrError(t)
	ch2 := make(chan string, 1)
	its.ClosedChan[string]().Match(ch2).OrError(t)
	// Output:
	// ✘ chan string is not closed.		--- @ ./general_test.go:342
}

func ExampleType() {
	its.Type[string]().Match("text value").OrError(t)
	its.Type[int]().Match("text value").OrError(t)
	type T struct {
		Foo int
	}
	its.Type[int]().Match(T{Foo: 42}).OrError(t)
	// Output:
	// ✘ /* got */ text value is a int		--- @ ./general_test.go:349
	//
	// ✘ /* got */ {Foo:42} is a int		--- @ ./general_test.go:353
}

func ExampleMatch() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.Match(pattern).Match([]byte("github.com")).OrError(t)
	its.Match(pattern).Match([]byte("github.com/youta-t/its")).OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).Match(/* got */ [103 105 116 104 117 98 46 99 111 109 47 121 111 117 116 97 45 116 47 105 116 115])		--- @ ./general_test.go:363
}

func ExampleMatchString() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.MatchString(pattern).Match("github.com").OrError(t)
	its.MatchString(pattern).Match("github.com/youta-t/its").OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).MatchString(/* got */ "github.com/youta-t/its")		--- @ ./general_test.go:371
}

func ExampleNil_pass() {
	var none *int
	its.Nil[*int]().Match(none).OrError(t)

	var chNil chan int = nil
	its.Nil[chan int]().Match(chNil).OrError(t)

	var fnNil func()
	its.Nil[func()]().Match(fnNil).OrError(t)

	its.Nil[any]().Match(nil).OrError(t)
	its.Nil[error]().Match(nil).OrError(t)

	var mapNil map[string]int
	its.Nil[map[string]int]().Match(mapNil).OrError(t)

	var sliceNil []int
	its.Nil[[]int]().Match(sliceNil).OrError(t)

	// Output:
}

func ExampleNil_fail() {
	three := 3
	its.Nil[*int]().Match(&three).OrError(t)

	chNonNil := make(chan int)
	its.Nil[chan int]().Match(chNonNil).OrError(t)

	funNonNil := func() {}
	its.Nil[func()]().Match(funNonNil).OrError(t)

	its.Nil[any]().Match(struct{}{}).OrError(t)
	its.Nil[error]().Match(errors.New("error")).OrError(t)

	mapNonNil := map[string]int{}
	its.Nil[map[string]int]().Match(mapNonNil).OrError(t)

	sliceNonNil := []int{}
	its.Nil[[]int]().Match(sliceNonNil).OrError(t)

	its.Nil[int]().Match(3).OrError(t)
	// Output:
	// ✘ (/* got */ 3) is nil		--- @ ./general_test.go:400
	//
	// ✘ (/* got */ chan int) is nil		--- @ ./general_test.go:403
	//
	// ✘ (/* got */ func()) is nil		--- @ ./general_test.go:406
	//
	// ✘ (/* got */ {}) is nil		--- @ ./general_test.go:408
	//
	// ✘ (/* got */ error) is nil		--- @ ./general_test.go:409
	//
	// ✘ (/* got */ map[]) is nil		--- @ ./general_test.go:412
	//
	// ✘ (/* got */ []) is nil		--- @ ./general_test.go:415
	//
	// ✘ (/* got */ 3) is nil		--- @ ./general_test.go:417
}

func ExamplePointer() {
	got := 42
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)

	got = 40
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)

	var ptrgot *int
	its.Pointer(its.EqEq(42)).Match(ptrgot).OrError(t)
	// Output:
	// ✘ /* got */ *int is not nil,		--- @ ./general_test.go:441
	//     ✘ /* got */ 40 == /* want */ 42		--- @ ./general_test.go:441
	//
	// ✘ /* got */ nil is not nil,		--- @ ./general_test.go:444
	//     ✘ /* got */ ?? == /* want */ 42		--- @ ./general_test.go:444
}
