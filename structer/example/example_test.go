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

func Example_match_struct_ng() {
	ItsMyStruct(MyStructSpec{
		Name:      its.StringHavingPrefix("its"),
		Value:     its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
		Timestamp: its.Always[time.Time](),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 22, 30},
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
		Name:      its.StringHavingPrefix("its"),
		Value:     its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
		Timestamp: its.Equal(wantTimestamp),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 20, 30},
			Timestamp: wantTimestamp.Add(+5 * time.Minute),
		}).
		OrError(t)

	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:42
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:42
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:43
	//     ✘ .Value :		--- @ ./structer/example/example_test.go:42
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:44
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:44
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:44
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:44
	//     ✔ .Timestamp :		--- @ ./structer/example/example_test.go:42
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:45
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:60
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:60
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:61
	//     ✔ .Value :		--- @ ./structer/example/example_test.go:60
	//         ✔ []int{ ... (len: /* got */ 3, /* want */ 3; +0, -0)		--- @ ./structer/example/example_test.go:62
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:62
	//             ✔ /* got */ 20 == /* want */ 20		--- @ ./structer/example/example_test.go:62
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:62
	//     ✘ .Timestamp :		--- @ ./structer/example/example_test.go:60
	//         ✘ (/* want */ 2024-01-02 13:14:15 +0000 +0000).Equal(/* got */ 2024-01-02 13:19:15 +0000 +0000)		--- @ ./structer/example/example_test.go:63
}

// nil field in spec fallbacks to its.Always()
func Example_match_partial_spec_default() {
	ItsMyStruct(MyStructSpec{
		Name:  its.StringHavingPrefix("its"),
		Value: its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 22, 30},
			Timestamp: time.Now(),
		}).
		OrError(t)

	ItsMyStruct(MyStructSpec{
		Name:  its.StringHavingPrefix("its"),
		Value: its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 22, 30},
			Timestamp: time.Now(),
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:99
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:99
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:100
	//     ✘ .Value :		--- @ ./structer/example/example_test.go:99
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:101
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:101
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:101
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:101
	//     ✔ .Timestamp :		--- @ ./structer/example/example_test.go:99
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:99
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:110
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:110
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:111
	//     ✘ .Value :		--- @ ./structer/example/example_test.go:110
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:112
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:112
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:112
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:112
	//     ✔ .Timestamp :		--- @ ./structer/example/example_test.go:110
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:110
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
		Name:  its.StringHavingPrefix("its"),
		Value: its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 22, 30},
			Timestamp: gotTimestamp,
		}).
		OrError(t)

	ItsMyStruct(MyStructSpec{
		Name:  its.StringHavingPrefix("its"),
		Value: its.ForItems(its.Slice, its.EqEq, []int{10, 20, 30}),
	}).
		Match(internal.MyStruct{
			Name:      "its a matching library",
			Value:     []int{10, 22, 30},
			Timestamp: gotTimestamp,
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:158
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:158
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:159
	//     ✘ .Value :		--- @ ./structer/example/example_test.go:158
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:160
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:160
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:160
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:160
	//     ✘ .Timestamp :		--- @ ./structer/example/example_test.go:158
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:158
	//
	// ✘ type MyStruct:		--- @ ./structer/example/example_test.go:169
	//     ✔ .Name :		--- @ ./structer/example/example_test.go:169
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:170
	//     ✘ .Value :		--- @ ./structer/example/example_test.go:169
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:171
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:171
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:171
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:171
	//     ✘ .Timestamp :		--- @ ./structer/example/example_test.go:169
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:169
}
