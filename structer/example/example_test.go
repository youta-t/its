package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/config"
	types "github.com/youta-t/its/structer/example/internal"
	in_gen "github.com/youta-t/its/structer/example/internal/gen"
	"github.com/youta-t/its/structer/example/internal/sub1"
	sub1_gen "github.com/youta-t/its/structer/example/internal/sub1/gen"
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
	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Values: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
		Sub1: sub1_gen.ItsSub1(sub1_gen.Sub1Spec{
			StringField: its.EqEq("nested, tested!"),
		}),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested?",
			},
		}).
		OrError(t)

	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Values: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
		Sub1: sub1_gen.ItsSub1(sub1_gen.Sub1Spec{
			StringField: its.EqEq("nested, tested!"),
		}),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 20, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested!",
			},
		}).
		OrError(t)

	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Values: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
		Sub1: sub1_gen.ItsSub1(sub1_gen.Sub1Spec{
			StringField: its.Some(
				its.StringHavingPrefix("tested!"),
				its.StringHavingSuffix("tested!"),
			),
		}),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested!",
			},
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct1:		--- @ ./structer/example/example_test.go:43
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:44
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:45
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:46
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:47
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:48
	//     ✘ .Sub1 :
	//         ✘ type Sub1:
	//             ✘ .StringField :
	//                 ✘ /* got */ nested, tested? == /* want */ nested, tested!		--- @ ./structer/example/example_test.go:51
	//
	// ✘ type MyStruct1:		--- @ ./structer/example/example_test.go:87
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:88
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:89
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:90
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:91
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:92
	//     ✔ .Sub1 :
	//         ✔ type Sub1:
	//             ✔ .StringField :
	//                 ✔ // some: (1 ok / 2 matchers)		--- @ ./structer/example/example_test.go:95
	//                     ~ strings.HasPrefix(/* got */ "nested, tested!", /* want */ "tested!")		--- @ ./structer/example/example_test.go:96
	//                     ✔ strings.HasSuffix(/* got */ "nested, tested!", /* want */ "tested!")		--- @ ./structer/example/example_test.go:97
}

// nil field in spec fallbacks to its.Always()
func Example_match_partial_spec_default() {
	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Values: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested?",
			},
		}).
		OrError(t)

	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Sub1: sub1_gen.ItsSub1(sub1_gen.Sub1Spec{
			StringField: its.Some(
				its.StringHavingPrefix("tested!"),
				its.StringHavingSuffix("tested!"),
			),
		}),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested!",
			},
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct1:		--- @ ./structer/example/example_test.go:145
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:146
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:147
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:148
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:149
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:150
	//     ✔ .Sub1 :
	//         ✔ (always pass)		--- @ ./structer/example/example_test.go:145
}

// In strict mode,
func Example_match_partial_spec_strict() {
	config.StrictStructMatch(t, true)
	defer t.DoCleanup() // emulate *testing.T

	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Values: its.Slice(
			its.EqEq(10),
			its.EqEq(20),
			its.EqEq(30),
		),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested?",
			},
		}).
		OrError(t)

	in_gen.ItsMyStruct1(in_gen.MyStruct1Spec{
		Name: its.StringHavingPrefix("its"),
		Sub1: sub1_gen.ItsSub1(sub1_gen.Sub1Spec{
			StringField: its.Some(
				its.StringHavingPrefix("tested!"),
				its.StringHavingSuffix("tested!"),
			),
		}),
	}).
		Match(types.MyStruct1{
			Name: "its a matching library",
			Values: []int{
				10, 22, 30,
			},
			Sub1: sub1.Sub1{
				StringField: "nested, tested!",
			},
		}).
		OrError(t)
	// Output:
	// ✘ type MyStruct1:		--- @ ./structer/example/example_test.go:202
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:203
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./structer/example/example_test.go:204
	//             ✔ /* got */ 10 == /* want */ 10		--- @ ./structer/example/example_test.go:205
	//             ✘ - /* got */ ?? == /* want */ 20		--- @ ./structer/example/example_test.go:206
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30		--- @ ./structer/example/example_test.go:207
	//     ✘ .Sub1 :
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:202
	//
	// ✘ type MyStruct1:		--- @ ./structer/example/example_test.go:221
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")		--- @ ./structer/example/example_test.go:222
	//     ✘ .Values :
	//         ✘ (never pass)		--- @ ./structer/example/example_test.go:221
	//     ✔ .Sub1 :
	//         ✔ type Sub1:
	//             ✔ .StringField :
	//                 ✔ // some: (1 ok / 2 matchers)		--- @ ./structer/example/example_test.go:224
	//                     ~ strings.HasPrefix(/* got */ "nested, tested!", /* want */ "tested!")		--- @ ./structer/example/example_test.go:225
	//                     ✔ strings.HasSuffix(/* got */ "nested, tested!", /* want */ "tested!")		--- @ ./structer/example/example_test.go:226
}
