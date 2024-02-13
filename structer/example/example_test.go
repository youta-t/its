package example_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/youta-t/its"
	"github.com/youta-t/its/config"
	"github.com/youta-t/its/structer/example/internal"
	. "github.com/youta-t/its/structer/example/internal/gen_structer"
)

type FakeT struct {
	c func()
}

func (*FakeT) Error(param ...any) {
	fmt.Println(param...)
}

func (f *FakeT) Cleanup(fn func()) {
	f.c = fn
}

func (f *FakeT) DoCleanup() {
	if f.c != nil {
		f.c()
		f.c = nil
	}
}

var t = &FakeT{}

func TestMain(m *testing.M) {
	config.ReplaceProjectRoot()
	os.Exit(m.Run())
}

func Example_match_struct() {
	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
		Timestamp: its.Always[time.Time](),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 22, 30,
			},
			Timestamp: time.Now(),
		}).
		OrError(t)

	wantTimestamp, err := time.Parse(
		time.RFC3339, "2024-01-02T13:14:15+00:00",
	)
	if err != nil {
		panic(err)
	}
	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
		Timestamp: its.Equal(wantTimestamp),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 20, 30,
			},
			Timestamp: wantTimestamp.Add(+5 * time.Minute),
		}).
		OrError(t)

	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:42
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:43
	//     ✘ .Value :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:44
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:45
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:46
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:47
	//     ✔ .Timestamp :
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:49
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:66
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:67
	//     ✔ .Value :
	//         ✔ []int{ ... (len: /* got */ 3, /* want */ 3; +0, -0)		--- @ ./structer/example/example_test.go:68
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:69
	//             ✔ /* got */ 20 == /* want */ 20		--- @ ./structer/example/example_test.go:70
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:71
	//     ✘ .Timestamp :
	//         ✘ (/* want */ 2024-01-02 13:14:15 +0000 +0000).Equal(/* got */ 2024-01-02 13:19:15 +0000 +0000)		--- @ ./structer/example/example_test.go:73
}

// nil field in spec fallbacks to its.Always()
func Example_match_partial_spec_default() {
	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 22, 30,
			},
			Timestamp: time.Now(),
		}).
		OrError(t)

	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 22, 30,
			},
			Timestamp: time.Now(),
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:111
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:112
	//     ✘ .Value :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:113
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:114
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:115
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:116
	//     ✔ .Timestamp :
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:111
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:128
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:129
	//     ✘ .Value :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:130
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:131
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:132
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:133
	//     ✔ .Timestamp :
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:128
}

// In strict mode,
func Example_match_partial_spec_strict() {
	config.StrictStructMatch(t, true)
	defer t.DoCleanup() // emulate *testing.T

	gotTimestamp, err := time.Parse(
		time.RFC3339, "2024-01-02T13:14:15+00:00",
	)
	if err != nil {
		panic(err)
	}

	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 22, 30,
			},
			Timestamp: gotTimestamp,
		}).
		OrError(t)

	ItsMyStruct(MyStructSpec{
		Name: its.StringHavingPrefix("its"),
		Value: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(internal.MyStruct{
			Name: "its a matching library",
			Value: []int{
				10, 22, 30,
			},
			Timestamp: gotTimestamp,
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:182
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:183
	//     ✘ .Value :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:184
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:185
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:186
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:187
	//     ✘ .Timestamp :
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:182
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:199
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:200
	//     ✘ .Value :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:201
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:202
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:203
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:204
	//     ✘ .Timestamp :
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:199
}
