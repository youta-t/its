package example_test

import (
	"fmt"

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
	// ✘ type MyStruct1:
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
	//             ✔ /* got */ 10 == /* want */ 10
	//             ✘ - /* got */ ?? == /* want */ 20
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30
	//     ✘ .Sub1 :
	//         ✘ type Sub1:
	//             ✘ .StringField :
	//                 ✘ /* got */ nested, tested? == /* want */ nested, tested!
	//
	// ✘ type MyStruct1:
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
	//             ✔ /* got */ 10 == /* want */ 10
	//             ✘ - /* got */ ?? == /* want */ 20
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30
	//     ✔ .Sub1 :
	//         ✔ type Sub1:
	//             ✔ .StringField :
	//                 ✔ // some: (1 ok / 2 matchers)
	//                     ~ strings.HasPrefix(/* got */ "nested, tested!", /* want */ "tested!")
	//                     ✔ strings.HasSuffix(/* got */ "nested, tested!", /* want */ "tested!")
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
	// ✘ type MyStruct1:
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
	//             ✔ /* got */ 10 == /* want */ 10
	//             ✘ - /* got */ ?? == /* want */ 20
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30
	//     ✔ .Sub1 :
	//         ✔ (always pass)
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
	// ✘ type MyStruct1:
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")
	//     ✘ .Values :
	//         ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)
	//             ✔ /* got */ 10 == /* want */ 10
	//             ✘ - /* got */ ?? == /* want */ 20
	//             ✘ + /* got */ 22
	//             ✔ /* got */ 30 == /* want */ 30
	//     ✘ .Sub1 :
	//         ✘ (never pass)
	//
	// ✘ type MyStruct1:
	//     ✔ .Name :
	//         ✔ strings.HasPrefix(/* got */ "its a matching library", /* want */ "its")
	//     ✘ .Values :
	//         ✘ (never pass)
	//     ✔ .Sub1 :
	//         ✔ type Sub1:
	//             ✔ .StringField :
	//                 ✔ // some: (1 ok / 2 matchers)
	//                     ~ strings.HasPrefix(/* got */ "nested, tested!", /* want */ "tested!")
	//                     ✔ strings.HasSuffix(/* got */ "nested, tested!", /* want */ "tested!")
}
