package its_test

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/youta-t/its"
)

func ExampleEqEq_ok() {
	its.EqEq(42).Match(42).OrError(t) // pass. so no messages are output
}

func ExampleEqEq_ng() {
	its.EqEq(42).Match(49).OrError(t) // fail!
	// Output:
	// ✘ /* got */ 49 == /* want */ 42		--- @ ./general_test.go:18
}

func ExampleEqEq_ng_non_primitive_type() {
	type MyType struct {
		Foo int
	}

	its.EqEq(MyType{Foo: 42}).Match(MyType{Foo: 24}).OrError(t) // also fail!

	// Output:
	// ✘ /* got */ {Foo:24} == /* want */ {Foo:42}		--- @ ./general_test.go:28
}

func ExampleGreaterThan_int() {
	its.GreaterThan(10).Match(10).OrError(t)
	its.GreaterThan(11).Match(10).OrError(t)
	its.GreaterThan(10).Match(11).OrError(t)
	// Output:
	// ✘ /* want */ 10 < /* got */ 10		--- @ ./general_test.go:35
	//
	// ✘ /* want */ 11 < /* got */ 10		--- @ ./general_test.go:36
}

func ExampleGreaterThan_float() {
	its.GreaterThan(1.0).Match(1.0).OrError(t)
	its.GreaterThan(1.1).Match(1.0).OrError(t)
	its.GreaterThan(1.0).Match(1.1).OrError(t)
	// Output:
	// ✘ /* want */ 1 < /* got */ 1		--- @ ./general_test.go:45
	//
	// ✘ /* want */ 1.1 < /* got */ 1		--- @ ./general_test.go:46
}

func ExampleGreaterThan_string() {
	its.GreaterThan("aaa").Match("aaa").OrError(t)
	its.GreaterThan("aab").Match("aaa").OrError(t)
	its.GreaterThan("aaa").Match("aab").OrError(t)

	// Output:
	// ✘ /* want */ aaa < /* got */ aaa		--- @ ./general_test.go:55
	//
	// ✘ /* want */ aab < /* got */ aaa		--- @ ./general_test.go:56
}

func ExampleGreaterEq_int() {
	//      <=
	its.GreaterEq(10).Match(11).OrError(t)
	its.GreaterEq(11).Match(10).OrError(t)
	its.GreaterEq(10).Match(10).OrError(t)
	// Output:
	// ✘ /* want */ 11 <= /* got */ 10		--- @ ./general_test.go:68
}

func ExampleGreaterEq_float() {

	its.GreaterEq(1.0).Match(1.1).OrError(t)
	its.GreaterEq(1.0).Match(1.0).OrError(t)
	its.GreaterEq(1.1).Match(1.0).OrError(t)
	// Output:
	// ✘ /* want */ 1.1 <= /* got */ 1		--- @ ./general_test.go:78
}

func ExampleGreaterEq_sring() {
	its.GreaterEq("aaa").Match("aab").OrError(t)
	its.GreaterEq("aaa").Match("aaa").OrError(t)
	its.GreaterEq("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aab <= /* got */ aaa		--- @ ./general_test.go:86
}

func ExampleLesserThan_int() {
	//       >
	its.LesserThan(10).Match(11).OrError(t)
	its.LesserThan(11).Match(10).OrError(t)
	its.LesserThan(10).Match(10).OrError(t)
	// Output:
	// ✘ /* want */ 10 > /* got */ 11		--- @ ./general_test.go:94
	//
	// ✘ /* want */ 10 > /* got */ 10		--- @ ./general_test.go:96
}

func ExampleLesserThan_float() {
	its.LesserThan(1.0).Match(1.1).OrError(t)
	its.LesserThan(1.0).Match(1.0).OrError(t)
	its.LesserThan(1.1).Match(1.0).OrError(t)
	// Output:
	// ✘ /* want */ 1 > /* got */ 1.1		--- @ ./general_test.go:104
	//
	// ✘ /* want */ 1 > /* got */ 1		--- @ ./general_test.go:105
}

func ExampleLesserThan_string() {

	its.LesserThan("aaa").Match("aab").OrError(t)
	its.LesserThan("aaa").Match("aaa").OrError(t)
	its.LesserThan("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aaa > /* got */ aab		--- @ ./general_test.go:115
	//
	// ✘ /* want */ aaa > /* got */ aaa		--- @ ./general_test.go:116
}

func ExampleLesserEq_int() {
	//      >=
	its.LesserEq(10).Match(11).OrError(t)
	its.LesserEq(11).Match(10).OrError(t)
	its.LesserEq(10).Match(10).OrError(t)

	// Output:
	// ✘ /* want */ 10 >= /* got */ 11		--- @ ./general_test.go:127
}

func ExampleLesserEq_float() {
	its.LesserEq(1.0).Match(1.1).OrError(t)
	its.LesserEq(1.0).Match(1.0).OrError(t)
	its.LesserEq(1.1).Match(1.0).OrError(t)

	// Output:
	//　✘ /* want */ 1 >= /* got */ 1.1		--- @ ./general_test.go:136
}

func ExampleLesserEq_string() {
	its.LesserEq("aaa").Match("aab").OrError(t)
	its.LesserEq("aaa").Match("aaa").OrError(t)
	its.LesserEq("aab").Match("aaa").OrError(t)

	// Output:
	// ✘ /* want */ aaa >= /* got */ aab		--- @ ./general_test.go:145
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
	// ✘ (/* got */ 2023-10-11 12:18:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:163
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:164
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
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:184
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./general_test.go:185
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
	// ✘ (/* want */ 2023-10-11 12:13:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./general_test.go:203
	//
	// ✘ (/* want */ 2023-10-11 12:18:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./general_test.go:205
}

func ExampleEquivWith_ok() {
	its.EquivWith(
		42,
		func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).
		Match("42").
		OrError(t)
	// Output:
}

func ExampleEquivWith_ng() {
	its.EquivWith(
		42,
		func(want int, got string) bool { return fmt.Sprintf("%d", want) == got },
	).
		Match("40").
		OrError(t)

	// Output:
	// ✘ (/* want */ 42) equiv. (/* got */ 40)		--- @ ./general_test.go:224
}

func ExampleError() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e1).Match(e2).OrError(t)
	its.Error(e1).Match(e1).OrError(t)
	// Output:
}

func ExampleError_ng() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e2).Match(e1).OrError(t)

	// Output:
	// ✘ errors.Is(/* got */ error, /* want */ wrapped: error)		--- @ ./general_test.go:248
}

type CustomError string

func (c CustomError) Error() string {
	return string(c)
}

func ExampleErrorAs_ok() {
	customErr := CustomError("custom error")

	its.ErrorAs[CustomError]().Match(customErr).OrError(t)
	// Output:
}

func ExampleErrorAs_ng() {
	otherErr := errors.New("error by error.New")
	its.ErrorAs[CustomError]().Match(otherErr).OrError(t)

	// Output:
	// ✘ want := new(its_test.CustomError); errors.As(/* got */ error by error.New, want)		--- @ ./general_test.go:269
}

func ExampleAlways() {
	its.Always[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
}

func ExampleNever() {
	its.Never[any]().Match(struct{ Arbitary string }{}).OrError(t)
	// Output:
	// ✘ (never pass)		--- @ ./general_test.go:281
}

func ExampleStringHavingPrefix_ok() {
	its.StringHavingPrefix("abc").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringHavingPrefix_ng() {
	its.StringHavingPrefix("abc").Match("adcbe").OrError(t)

	// Output:
	// ✘ strings.HasPrefix(/* got */ "adcbe", /* want */ "abc")		--- @ ./general_test.go:292
}

func ExampleStringHavingSuffix_ok() {
	its.StringHavingSuffix("cde").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringHavingSuffix_ng() {
	its.StringHavingSuffix("cde").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.HasSuffix(/* got */ "adcbe", /* want */ "cde")		--- @ ./general_test.go:304
}

func ExampleStringContaining_ok() {
	its.StringContaining("bcd").Match("abcde").OrError(t)
	// Output:
}

func ExampleStringContaining_ng() {
	its.StringContaining("bcd").Match("adcbe").OrError(t)
	// Output:
	// ✘ strings.Contains(/* got */ "adcbe", /* want */ "bcd")		--- @ ./general_test.go:315
}

func ExampleStringEqualFold_ok() {
	its.StringEqualFold("abc").Match("abc").OrError(t)
	its.StringEqualFold("aBc").Match("AbC").OrError(t)
	// Output:
}

func ExampleStringEqualFold_ng() {
	its.StringEqualFold("abc").Match("αβγ").OrError(t)
	// Output:
	// ✘ strings.EqualFold(/* got */ "αβγ", /* want */ "abc")		--- @ ./general_test.go:327
}

func ExampleBytesEqual_ok() {
	its.BytesEqual([]byte("abc")).Match([]byte("abc")).OrError(t)
	// Output:
}

func ExampleBytesEqual_ng() {
	its.BytesEqual([]byte("abc")).Match([]byte("acb")).OrError(t)
	// Output:
	// ✘ bytes.Equal(/* got */ []byte{0x61, 0x63, 0x62}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./general_test.go:338
}

func ExampleBytesHavingPrefix_ok() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesHavingPrefix_ng() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasPrefix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./general_test.go:349
}

func ExampleBytesHavingSuffix_ok() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesHavingSuffix_ng() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasSuffix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x63, 0x64, 0x65})		--- @ ./general_test.go:360
}

func ExampleBytesContaining_ok() {
	its.BytesContaining([]byte("bcd")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesContaining_ng() {
	its.BytesContaining([]byte("bcd")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.Contains(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x62, 0x63, 0x64})		--- @ ./general_test.go:371
}

func ExampleNaN_ok() {
	its.NaN().Match(math.NaN()).OrError(t)
	// Output:
}

func ExampleNaN_ng() {
	its.NaN().Match(42).OrError(t)
	// Output:
	// ✘ math.IsNaN(/* got */ 42.000000)		--- @ ./general_test.go:382
}

func ExampleInf() {
	its.Inf().Match(math.Inf(1)).OrError(t)
	its.Inf().Match(math.Inf(-1)).OrError(t)
	// Output:
}

func ExampleInf_ng() {
	its.Inf().Match(0).OrError(t)
	// Output:
	// ✘ math.IsInf(/* got */ 0.000000, 0)		--- @ ./general_test.go:394
}

func ExampleClosedChan_ok() {
	ch1 := make(chan int, 1)
	close(ch1)
	its.ClosedChan[int]().Match(ch1).OrError(t)
	// Output:
}
func ExampleClosedChan_ng() {
	ch2 := make(chan string, 1)
	its.ClosedChan[string]().Match(ch2).OrError(t)
	// Output:
	// ✘ chan string is not closed.		--- @ ./general_test.go:407
}

func ExampleType_primitive() {
	its.Type[string]().Match("text value").OrError(t)
	its.Type[string]().Match(42).OrError(t)
	// Output:
	// ✘ /* got */ 42 is a string		--- @ ./general_test.go:414
}

func ExampleType_non_primitive() {
	type T struct {
		Foo int
	}
	type U struct {
		Bar int
	}
	its.Type[T]().Match(T{Foo: 42}).OrError(t)
	its.Type[U]().Match(T{Foo: 42}).OrError(t)
	// Output:
	// ✘ /* got */ {Foo:42} is a its_test.U		--- @ ./general_test.go:427
}

func ExampleMatch_ok() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.Match(pattern).Match([]byte("github.com")).OrError(t)
	// Output:
}

func ExampleMatch_ng() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.Match(pattern).Match([]byte("github.com/youta-t/its")).OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).Match(/* got */ [103 105 116 104 117 98 46 99 111 109 47 121 111 117 116 97 45 116 47 105 116 115])		--- @ ./general_test.go:440
}

func ExampleMatchString() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.MatchString(pattern).Match("github.com").OrError(t)
	// Output:
}

func ExampleMatchString_ng() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.MatchString(pattern).Match("github.com/youta-t/its").OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).MatchString(/* got */ "github.com/youta-t/its")		--- @ ./general_test.go:453
}

func ExampleNil_ok() {
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

func ExampleNil_ng_value_ptr() {
	three := 3
	its.Nil[*int]().Match(&three).OrError(t)

	// Output:
	// ✘ (/* got */ 3) is nil		--- @ ./general_test.go:482
}

func ExampleNil_ng_chan() {
	chNonNil := make(chan int)
	its.Nil[chan int]().Match(chNonNil).OrError(t)
	// Output:
	// ✘ (/* got */ chan int) is nil		--- @ ./general_test.go:490
}

func ExampleNil_ng_func() {
	funNonNil := func() {}
	its.Nil[func()]().Match(funNonNil).OrError(t)
	// Output:
	//
	// ✘ (/* got */ func()) is nil		--- @ ./general_test.go:497
}

func ExampleNil_ng_nonpointer_value() {
	its.Nil[any]().Match(struct{}{}).OrError(t)
	its.Nil[int]().Match(3).OrError(t)
	// Output:
	// ✘ (/* got */ {}) is nil		--- @ ./general_test.go:504
	//
	// ✘ (/* got */ 3) is nil		--- @ ./general_test.go:505
}

func ExampleNil_ng_interface() {
	its.Nil[error]().Match(errors.New("error")).OrError(t)
	// Output:
	// ✘ (/* got */ error) is nil		--- @ ./general_test.go:513
}

func ExampleNil_ng_map() {
	mapNonNil := map[string]int{}
	its.Nil[map[string]int]().Match(mapNonNil).OrError(t)
	// Output:
	//✘ (/* got */ map[]) is nil		--- @ ./general_test.go:520
}

func ExampleNil_ng_slice() {
	sliceNonNil := []int{}
	its.Nil[[]int]().Match(sliceNonNil).OrError(t)

	// Output:
	//
	// ✘ (/* got */ []) is nil		--- @ ./general_test.go:527
}

func ExamplePointer_ok() {
	got := 42
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)
	// Output:
}

func ExamplePointer_ng() {
	got := 40
	its.Pointer(its.EqEq(42)).Match(&got).OrError(t)

	var ptrgot *int
	its.Pointer(its.EqEq(42)).Match(ptrgot).OrError(t)
	// Output:
	// ✘ /* got */ *int is not nil,		--- @ ./general_test.go:542
	//     ✘ /* got */ 40 == /* want */ 42		--- @ ./general_test.go:542
	//
	// ✘ /* got */ nil is not nil,		--- @ ./general_test.go:545
	//     ✘ /* got */ ?? == /* want */ 42		--- @ ./general_test.go:545
}

func ExampleText_multibytes() {
	// "風景" (純銀もざいく; 山村暮鳥): https://www.aozora.gr.jp/cards/000136/files/52348_42039.html
	its.Text(`
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
かすかなるむぎぶえ
いちめんのなのはな
`).Match(`
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
ひばりのおしやべり
いちめんのなのはな
`).OrError(t)
	// Output:
	// ✘ (+ = got, - = want)		--- @ ./general_test.go:556
	//       |
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//     - | かすかなるむぎぶえ
	//     + | ひばりのおしやべり
	//       | いちめんのなのはな
	//       |
}

func ExampleText_ascii() {
	// London Bridge Is Falling Down
	its.Text(`
Build it up with bricks and mortar,
Bricks and mortar, bricks and mortar,
Build it up with bricks and mortar,
My fair lady.
`).Match(`
Build it up with iron and steel,
Iron and steel, iron and steel,
Build it up with iron and steel,
My fair lady.
`).OrError(t)
	// Output:
	//
	// ✘ (+ = got, - = want)		--- @ ./general_test.go:595
	//       |
	//     - | Build it up with bricks and mortar,
	//     - | Bricks and mortar, bricks and mortar,
	//     - | Build it up with bricks and mortar,
	//     + | Build it up with iron and steel,
	//     + | Iron and steel, iron and steel,
	//     + | Build it up with iron and steel,
	//       | My fair lady.
	//       |
}

func ExampleText() {
	its.Text(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
    nisi ut aliquip ex ea commodo consequat.
    Duis aute irure dolor in reprehenderit in voluptate velit
    esse cillum dolore eu fugiat nulla pariatur.

    Excepteur sint occaecat cupidatat non proident,
    sunt in culpa qui officia deserunt mollit anim id est laborum.
`).Match(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

    nisi ut aliquip ex ea commodo consequat.
    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
    Duis aute irure dolor in reprehenderit in voluptate velit
    esse cillum dolore eu fugiat nulla pariatur.

    sunt in culpa qui officia deserunt mollit anim id est laborum.
    Excepteur sint occaecat cupidatat non proident,
`).OrError(t)
	// Output:
	// ✘ (+ = got, - = want)		--- @ ./general_test.go:621
	//       |
	//       | Lorem Ipsum:
	//       |
	//       |     Lorem ipsum dolor sit amet,
	//       |     consectetur adipiscing elit,
	//       |     sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	//       |
	//     - |     Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
	//       |     nisi ut aliquip ex ea commodo consequat.
	//     + |     Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
	//       |     Duis aute irure dolor in reprehenderit in voluptate velit
	//       |     esse cillum dolore eu fugiat nulla pariatur.
	//       |
	//     - |     Excepteur sint occaecat cupidatat non proident,
	//       |     sunt in culpa qui officia deserunt mollit anim id est laborum.
	//     + |     Excepteur sint occaecat cupidatat non proident,
	//       |
}
