package mockkit_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/config"
	"github.com/youta-t/its/mocker/internal/example"
	. "github.com/youta-t/its/mocker/internal/example/gen_mock"
	. "github.com/youta-t/its/mocker/internal/example/gen_structer"
	"github.com/youta-t/its/mocker/mockkit"
)

func TestMain(m *testing.M) {
	config.ReplaceProjectRoot()
	os.Exit(m.Run())
}

type FakeTest struct{}

func (FakeTest) Error(param ...any) {
	fmt.Println(param...)
}
func (FakeTest) Errorf(pattern string, param ...any) {
	fmt.Printf(pattern+"\n", param...)
}

func (FakeTest) Fatal(param ...any) {
	fmt.Println(param...)
	panic("fatal")
}
func (FakeTest) Fatalf(pattern string, param ...any) {
	fmt.Printf(pattern+"\n", param...)
	panic("fatalf")
}

func (FakeTest) Helper() {}

var t = FakeTest{}

func Example_mock_passed() {

	registry := UserRegistry_Build(t, UserRegistry_Spec{
		Get: UserRegistry_Get_Expects(its.StringHavingPrefix("user-id:")).
			ThenReturn(
				example.User{Id: "user-id:abcde", Name: "John Doe"},
				nil,
			),
	})

	registry2 := UserRegistry_Build(t, UserRegistry_Spec{
		Get: UserRegistry_Get_Expects(its.StringHavingPrefix("user-id:")).
			ThenEffect(func(userId string) (example.User, error) {
				return example.User{Id: "user-id:abcde", Name: "John Doe"}, nil
			}),
	})

	user, _ := registry.Get("user-id:abcdde")
	ItsUser(UserSpec{
		Id:   its.EqEq("user-id:abcde"),
		Name: its.EqEq("John Doe"),
	}).
		Match(user).
		OrError(t)

	user2, _ := registry2.Get("user-id:abcde")
	its.EqEq(example.User{Id: "user-id:abcde", Name: "John Doe"}).Match(user2).OrError(t)
	// Output:
}

func Example_behaviour_passed() {

	registry := UserRegistry_Spec{
		Get: UserRegistry_Get_Expects(its.StringHavingPrefix("user-id:")).
			ThenReturn(
				example.User{Id: "user-id:abcde", Name: "John Doe"},
				nil,
			),
	}

	registry2 := UserRegistry_Spec{
		Get: UserRegistry_Get_Expects(its.StringHavingPrefix("user-id:")).
			ThenEffect(func(userId string) (example.User, error) {
				return example.User{Id: "user-id:abcde", Name: "John Doe"}, nil
			}),
	}

	user, _ := UserRegistry_Build(t, registry).Get("user-id:abcdde")
	its.EqEq(example.User{Id: "user-id:abcde", Name: "John Doe"}).Match(user).OrError(t)

	user2, _ := UserRegistry_Build(t, registry2).Get("user-id:abcde")
	its.EqEq(example.User{Id: "user-id:abcde", Name: "John Doe"}).Match(user2).OrError(t)
	// Output:
}

func Example_mock_failed_by_assertion() {

	registry := UserRegistry_Build(
		t,
		UserRegistry_Spec{
			Get: UserRegistry_Get_Expects(its.StringHavingPrefix("superuser:")).
				ThenReturn(
					example.User{Id: "user-id:abcde", Name: "John Doe"},
					nil,
				),
		},
	)

	registry.Get("user-id:abcdde")
	// Output:
	// ✘ func UserRegistry_Get		--- @ ./mocker/mockkit/scenario_test.go:103
	//     ✘ userId :
	//         ✘ strings.HasPrefix(/* got */ "user-id:abcdde", /* want */ "superuser:")		--- @ ./mocker/mockkit/scenario_test.go:103
}

func Example_mock_failed_by_calling_not_mocked_method() {
	defer func() {
		recover()
	}()

	registry := UserRegistry_Build(
		t,
		UserRegistry_Spec{
			Get: UserRegistry_Get_Expects(its.StringHavingPrefix("superuser:")).
				ThenReturn(
					example.User{Id: "user-id:abcde", Name: "John Doe"},
					nil,
				),
		},
	)

	registry.Update(example.User{Id: "user-id:abcde", Name: "John Doe"})
	// Output:
	// ✘ UserRegistry.Update is not mocked		--- @ ./mocker/mockkit/scenario_test.go:134
}

func Example_scenario_passed() {
	// // assume them:
	//
	// type User struct {
	// 	Id   string
	// 	Name string
	// }
	//
	// type UserRegistry interface {
	// 	Get(userId string) (User, error)
	// 	Update(User) error
	// 	Delete(User) error
	// }
	//
	// type SessionStore func(cookie string) (userId string, ok bool)
	//
	// // Let us test this web-like feature.
	//
	// func UpdateUser(
	// 	sess SessionStore,
	// 	registry UserRegistry,
	// ) func(cookie string, newName string) error {
	// 	return func(cookie, newName string) error {
	// 		userId, ok := sess(cookie)
	// 		if !ok {
	// 			return errors.New("you are not logged in")
	// 		}
	// 		user, err := registry.Get(userId)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		user.Name = newName
	// 		return registry.Update(user)
	// 	}
	// }
	//
	// // SessionStore and UserRegistry will be injected to UpdateUser.
	// // I want to test if these injected objects are handled propery.

	sc := mockkit.BeginScenario(t)
	defer sc.End()

	sess := mockkit.Next( // add a function into scenario
		sc,
		SessionStore_Expects(
			its.EqEq("fake-cookie")).           // expectation for arguments
			ThenReturn("sample-user-id", true), // fixture for retrun valeus
	)

	registry := UserRegistry_Spec{}
	registry.Get = mockkit.Next(
		sc,
		UserRegistry_Get_Expects(its.EqEq("sample-user-id")).
			ThenReturn(
				example.User{
					Id:   "sample-user-id",
					Name: "John Doe",
				},
				nil,
			),
	)

	registry.Update = mockkit.Next(
		sc,
		UserRegistry_Update_Expects(
			ItsUser(
				UserSpec{
					Id:   its.EqEq("sample-user-id"),
					Name: its.EqEq("Richard Roe"),
				},
			),
		).
			ThenReturn(nil),
	)

	testee := example.UpdateUser(
		sess.Fn(t),
		UserRegistry_Build(t, registry),
	)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
}

func Example_scenario_failed_by_wrong_arg() {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	sess := mockkit.Next(
		sc,
		SessionStore_Expects(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true),
	)

	registry := UserRegistry_Spec{}
	registry.Get = mockkit.Next(
		sc,
		UserRegistry_Get_Expects(its.EqEq("sample-user-id")).
			ThenReturn(
				example.User{
					Id:   "wrong-user-id", // WRONG!!!
					Name: "John Doe",
				},
				nil,
			),
	)

	registry.Update = mockkit.Next(
		sc,
		UserRegistry_Update_Expects(
			ItsUser(UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			}),
		).
			ThenReturn(nil),
	)

	testee := example.UpdateUser(
		sess.Fn(t),
		UserRegistry_Build(t, registry),
	)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
	// ✘ func UserRegistry_Update		--- @ ./mocker/mockkit/scenario_test.go:248
	//     ✘ arg0 :
	//         ✘ type User:		--- @ ./mocker/mockkit/scenario_test.go:249
	//             ✘ .Id :		--- @ ./mocker/mockkit/scenario_test.go:249
	//                 ✘ /* got */ wrong-user-id == /* want */ sample-user-id		--- @ ./mocker/mockkit/scenario_test.go:250
	//             ✔ .Name :		--- @ ./mocker/mockkit/scenario_test.go:249
	//                 ✔ /* got */ Richard Roe == /* want */ Richard Roe		--- @ ./mocker/mockkit/scenario_test.go:251
}

func Example_scenario_failed_by_wrong_call_order() {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	registry := UserRegistry_Spec{}

	registry.Get = mockkit.Next( // call #2
		sc,
		UserRegistry_Get_Expects(its.EqEq("sample-user-id")).
			ThenReturn(
				example.User{Id: "sample-user-id", Name: "John Doe"},
				nil,
			),
	)

	sess := mockkit.Next( // call #1
		sc,
		SessionStore_Expects(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true),
	)

	registry.Update = mockkit.Next( // call #3
		sc,
		UserRegistry_Update_Expects(
			ItsUser(UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			}),
		).
			ThenReturn(nil),
	)

	testee := example.UpdateUser(
		sess.Fn(t),
		UserRegistry_Build(t, registry),
	)
	err := testee("fake-cookie", "Richard Roe")

	its.Nil[error]().Match(err).OrError(t)
	// Output:
	// ✘ // call order :
	//     ✘ []*mockkit.call{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./mocker/mockkit/scenario_test.go:319
	//         ✘ - (/* want */ ./mocker/mockkit/scenario_test.go:279).Equal(/* got */ ??)		--- @ ./mocker/mockkit/scenario_test.go:319
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:288).Equal(/* got */ ./mocker/mockkit/scenario_test.go:288 (invoked at ./mocker/internal/example/scenario.go:26))		--- @ ./mocker/mockkit/scenario_test.go:319
	//         ✘ + /* got */ ./mocker/mockkit/scenario_test.go:279 (invoked at ./mocker/internal/example/scenario.go:30)
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:294).Equal(/* got */ ./mocker/mockkit/scenario_test.go:294 (invoked at ./mocker/internal/example/scenario.go:35))		--- @ ./mocker/mockkit/scenario_test.go:319
}

func Example_scenario_failed_by_plans_incompleted() {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	sess := mockkit.Next(
		sc,
		SessionStore_Expects(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true),
	)

	registry := UserRegistry_Spec{}
	registry.Get = mockkit.Next(
		sc,
		UserRegistry_Get_Expects(its.EqEq("sample-user-id")).
			ThenReturn(
				example.User{
					Id:   "sample-user-id",
					Name: "John Doe",
				},
				nil,
			),
	)

	registry.Update = mockkit.Next(
		sc,
		UserRegistry_Update_Expects(
			ItsUser(
				UserSpec{
					Id:   its.EqEq("sample-user-id"),
					Name: its.EqEq("Richard Roe"),
				},
			),
		).
			ThenReturn(nil),
	)

	registry.Delete = mockkit.Next(
		sc,
		UserRegistry_Delete_Expects(its.Always[example.User]()).
			ThenReturn(nil),
	)

	testee := example.UpdateUser(
		sess.Fn(t),
		UserRegistry_Build(t, registry),
	)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
	// ✘ // call order :
	//     ✘ []*mockkit.call{ ... (len: /* got */ 3, /* want */ 4; +0, -1)		--- @ ./mocker/mockkit/scenario_test.go:376
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:325).Equal(/* got */ ./mocker/mockkit/scenario_test.go:325 (invoked at ./mocker/internal/example/scenario.go:26))		--- @ ./mocker/mockkit/scenario_test.go:376
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:332).Equal(/* got */ ./mocker/mockkit/scenario_test.go:332 (invoked at ./mocker/internal/example/scenario.go:30))		--- @ ./mocker/mockkit/scenario_test.go:376
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:344).Equal(/* got */ ./mocker/mockkit/scenario_test.go:344 (invoked at ./mocker/internal/example/scenario.go:35))		--- @ ./mocker/mockkit/scenario_test.go:376
	//         ✘ - (/* want */ ./mocker/mockkit/scenario_test.go:357).Equal(/* got */ ??)		--- @ ./mocker/mockkit/scenario_test.go:376
}

func TestSequential(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	seq := mockkit.Sequential(
		mockkit.Next(sc, mockkit.Effect(func() int { return 10 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 11 })),
	)
	seq.Append(
		mockkit.Next(sc, mockkit.Effect(func() int { return 12 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 13 })),
	)

	testee := seq.Fn(t)

	got := []int{
		testee(), testee(), testee(), testee(),
	}

	its.ForItems(its.Slice, its.EqEq, []int{10, 11, 12, 13}).Match(got).OrError(t)
}

func ExampleSequential_too_much_invoke() {
	sc := mockkit.BeginScenario(t)
	defer sc.End()
	defer func() {
		recover()
	}()

	seq := mockkit.Sequential(
		mockkit.Next(sc, mockkit.Effect(func() int { return 10 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 11 })),
	)
	inserted := mockkit.Next(sc, mockkit.Effect(func() int { return 99 })).Fn(t)
	seq.Append(
		mockkit.Next(sc, mockkit.Effect(func() int { return 12 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 13 })),
	)

	testee := seq.Fn(t)

	got := []int{
		testee(), testee(), inserted(), testee(), testee(),
	}

	its.ForItems(its.Slice, its.EqEq, []int{10, 11, 99, 12, 13}).Match(got).OrError(t)

	testee() // extra call!

	fmt.Println("does not reach here")

	// Output:
	// ✘ // invoke count :		--- @ ./mocker/mockkit/scenario_test.go:425
	//     ✘ /* want */ 4 >= /* got */ 5		--- @ ./mocker/mockkit/scenario_test.go:417
}

func ExampleSequential_out_of_order() {
	sc := mockkit.BeginScenario(t)
	defer sc.End()
	defer func() {
		recover()
	}()

	seq := mockkit.Sequential(
		mockkit.Next(sc, mockkit.Effect(func() int { return 10 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 11 })),
	)
	inserted := mockkit.Next(sc, mockkit.Effect(func() int { return 99 })).Fn(t)
	seq.Append(
		mockkit.Next(sc, mockkit.Effect(func() int { return 12 })),
		mockkit.Next(sc, mockkit.Effect(func() int { return 13 })),
	)

	testee := seq.Fn(t)

	got := []int{
		testee(),
		testee(),
		testee(),
		inserted(),
		testee(),
	}

	its.ForItems(its.Slice, its.EqEq, []int{10, 11, 12, 99, 13}).Match(got).OrError(t)

	// Output:
	// ✘ // call order :
	//     ✘ []*mockkit.call{ ... (len: /* got */ 5, /* want */ 5; +1, -1)		--- @ ./mocker/mockkit/scenario_test.go:472
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:442).Equal(/* got */ ./mocker/mockkit/scenario_test.go:442 (invoked at ./mocker/mockkit/scenario_test.go:454))		--- @ ./mocker/mockkit/scenario_test.go:472
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:443).Equal(/* got */ ./mocker/mockkit/scenario_test.go:443 (invoked at ./mocker/mockkit/scenario_test.go:455))		--- @ ./mocker/mockkit/scenario_test.go:472
	//         ✘ - (/* want */ ./mocker/mockkit/scenario_test.go:445).Equal(/* got */ ??)		--- @ ./mocker/mockkit/scenario_test.go:472
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:447).Equal(/* got */ ./mocker/mockkit/scenario_test.go:447 (invoked at ./mocker/mockkit/scenario_test.go:456))		--- @ ./mocker/mockkit/scenario_test.go:472
	//         ✘ + /* got */ ./mocker/mockkit/scenario_test.go:445 (invoked at ./mocker/mockkit/scenario_test.go:457)
	//         ✔ (/* want */ ./mocker/mockkit/scenario_test.go:448).Equal(/* got */ ./mocker/mockkit/scenario_test.go:448 (invoked at ./mocker/mockkit/scenario_test.go:458))		--- @ ./mocker/mockkit/scenario_test.go:472
}

func TestScenarioNext_WithoutArgs(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	f := mockkit.Next(sc, mockkit.Effect(func() bool { return false })).Fn(t)

	got := f()
	its.EqEq(false).Match(got).OrError(t)
}

func TestScenarioNext_WithoutVariadic(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	f := mockkit.Next(sc, mockkit.Effect(func(a int, b int, c int) bool {
		its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).Match([]int{a, b, c}).OrError(t)
		return false
	})).
		Fn(t)

	got := f(1, 2, 3)
	its.EqEq(false).Match(got).OrError(t)
}

func TestScenarioNext_WithSlice(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	f := mockkit.Next(sc, mockkit.Effect(func(args []int) bool {
		its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).Match(args).OrError(t)
		return false
	})).
		Fn(t)

	got := f([]int{1, 2, 3})
	its.EqEq(false).Match(got).OrError(t)
}

func TestScenarioNext_WithVariadic(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	f := mockkit.Next(sc, mockkit.Effect(func(args ...int) bool {
		its.ForItems(its.Slice, its.EqEq, []int{1, 2, 3}).Match(args).OrError(t)
		return false
	})).
		Fn(t)

	got := f(1, 2, 3)
	its.EqEq(false).Match(got).OrError(t)
}
