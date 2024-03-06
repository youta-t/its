package scenario_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/config"
	"github.com/youta-t/its/mocker/internal/example"
	. "github.com/youta-t/its/mocker/internal/example/gen_mock"
	. "github.com/youta-t/its/mocker/internal/example/gen_structer"
	"github.com/youta-t/its/mocker/scenario"
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

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get: NewUserRegistry_GetCall(its.StringHavingPrefix("user-id:")).
			ThenReturn(
				example.User{Id: "user-id:abcde", Name: "John Doe"},
				nil,
			).
			Mock(t),
	})

	registry2 := NewMockedUserRegistry(t, UserRegistryImpl{
		Get: NewUserRegistry_GetCall(its.StringHavingPrefix("user-id:")).
			ThenEffect(func(userId string) (example.User, error) {
				return example.User{Id: "user-id:abcde", Name: "John Doe"}, nil
			}).
			Mock(t),
	})

	user, _ := registry.Get("user-id:abcdde")
	its.EqEq(example.User{Id: "user-id:abcde", Name: "John Doe"}).Match(user).OrError(t)

	user2, _ := registry2.Get("user-id:abcde")
	its.EqEq(example.User{Id: "user-id:abcde", Name: "John Doe"}).Match(user2).OrError(t)
	// Output:
}

func Example_mock_failed_by_assertion() {

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get: NewUserRegistry_GetCall(its.StringHavingPrefix("superuser:")).
			ThenReturn(
				example.User{Id: "user-id:abcde", Name: "John Doe"},
				nil,
			).
			Mock(t),
	})

	registry.Get("user-id:abcdde")
	// Output:
	// ✘ func UserRegistry_Get		--- @ ./mocker/scenario/scenario_test.go:73
	//     ✘ strings.HasPrefix(/* got */ "user-id:abcdde", /* want */ "superuser:")		--- @ ./mocker/scenario/scenario_test.go:73
}

func Example_mock_failed_by_calling_not_mocked_method() {
	defer func() {
		recover()
	}()

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get: NewUserRegistry_GetCall(its.StringHavingPrefix("superuser:")).
			ThenReturn(
				example.User{Id: "user-id:abcde", Name: "John Doe"},
				nil,
			).
			Mock(t),
	})

	registry.Update(example.User{
		Id: "user-id:abcde", Name: "John Doe",
	})
	// Output:
	// ✘ UserRegistry.Update is not mocked		--- @ ./mocker/scenario/scenario_test.go:101
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

	sc := scenario.Begin(t)
	defer sc.End()

	sess := scenario.Next( // add a function into scenario
		sc, NewSessionStoreCall(
			its.EqEq("fake-cookie")).           // expectation for arguments
			ThenReturn("sample-user-id", true). // fixture for retrun valeus
			Mock(t),                            // build as mock function
	)

	getUser := scenario.Next(sc, NewUserRegistry_GetCall(
		its.EqEq("sample-user-id"),
	).
		ThenReturn(
			example.User{
				Id:   "sample-user-id",
				Name: "John Doe",
			},
			nil,
		).
		Mock(t),
	)

	updateUser := scenario.Next(sc, NewUserRegistry_UpdateCall(
		ItsUser(
			UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			},
		),
	).
		ThenReturn(nil).
		Mock(t),
	)

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get:    getUser,
		Update: updateUser,
	})

	testee := example.UpdateUser(sess, registry)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
}

func Example_scenario_failed_by_wrong_arg() {
	sc := scenario.Begin(t)
	defer sc.End()

	sess := scenario.Next(
		sc,
		NewSessionStoreCall(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true).
			Mock(t),
	)

	getUser := scenario.Next(sc, NewUserRegistry_GetCall(
		its.EqEq("sample-user-id"),
	).
		ThenReturn(
			example.User{
				Id:   "wrong-user-id", // WRONG!!!
				Name: "John Doe",
			},
			nil,
		).
		Mock(t),
	)

	updateUser := scenario.Next(sc, NewUserRegistry_UpdateCall(
		ItsUser(
			UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			},
		),
	).
		ThenReturn(nil).
		Mock(t),
	)

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get:    getUser,
		Update: updateUser,
	})

	testee := example.UpdateUser(sess, registry)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
	// ✘ func UserRegistry_Update		--- @ ./mocker/scenario/scenario_test.go:217
	//     ✘ type User:		--- @ ./mocker/scenario/scenario_test.go:218
	//         ✘ .Id :		--- @ ./mocker/scenario/scenario_test.go:218
	//             ✘ /* got */ wrong-user-id == /* want */ sample-user-id		--- @ ./mocker/scenario/scenario_test.go:220
	//         ✔ .Name :		--- @ ./mocker/scenario/scenario_test.go:218
	//             ✔ /* got */ Richard Roe == /* want */ Richard Roe		--- @ ./mocker/scenario/scenario_test.go:221
}

func Example_scenario_failed_by_wrong_call_order() {
	sc := scenario.Begin(t)
	defer sc.End()

	getUser := scenario.Next(sc, NewUserRegistry_GetCall(
		its.EqEq("sample-user-id"),
	).
		ThenReturn(
			example.User{
				Id:   "sample-user-id",
				Name: "John Doe",
			},
			nil,
		).
		Mock(t),
	)

	sess := scenario.Next(
		sc,
		NewSessionStoreCall(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true).
			Mock(t),
	)

	updateUser := scenario.Next(sc, NewUserRegistry_UpdateCall(
		ItsUser(
			UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			},
		),
	).
		ThenReturn(nil).
		Mock(t),
	)

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get:    getUser,
		Update: updateUser,
	})

	testee := example.UpdateUser(sess, registry)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
	// ✘ // scenario error: call order is out of plan		--- @ ./mocker/internal/example/scenario.go:26
	//     ✘ wanted to be called: ./mocker/scenario/scenario_test.go:250
	//
	// ✘ // scenario error: call order is out of plan		--- @ ./mocker/internal/example/scenario.go:35
	//     ✘ wanted to be called: ./mocker/scenario/scenario_test.go:263
	//
	// ✘ // there are functions planned but not called		--- @ ./mocker/scenario/scenario_test.go:300
	//     ✘ ./mocker/scenario/scenario_test.go:263
	//     ✘ ./mocker/scenario/scenario_test.go:270
}

func Example_scenario_failed_by_plans_incompleted() {
	sc := scenario.Begin(t)
	defer sc.End()

	sess := scenario.Next(
		sc,
		NewSessionStoreCall(its.EqEq("fake-cookie")).
			ThenReturn("sample-user-id", true).
			Mock(t),
	)

	getUser := scenario.Next(sc, NewUserRegistry_GetCall(
		its.EqEq("sample-user-id"),
	).
		ThenReturn(
			example.User{
				Id:   "sample-user-id",
				Name: "John Doe",
			},
			nil,
		).
		Mock(t),
	)

	updateUser := scenario.Next(sc, NewUserRegistry_UpdateCall(
		ItsUser(
			UserSpec{
				Id:   its.EqEq("sample-user-id"),
				Name: its.EqEq("Richard Roe"),
			},
		),
	).
		ThenReturn(nil).
		Mock(t),
	)

	scenario.Next(sc, NewUserRegistry_DeleteCall(
		its.Always[example.User](),
	).
		ThenReturn(nil).
		Mock(t),
	)

	registry := NewMockedUserRegistry(t, UserRegistryImpl{
		Get:    getUser,
		Update: updateUser,
	})

	testee := example.UpdateUser(sess, registry)

	its.Nil[error]().Match(testee("fake-cookie", "Richard Roe")).OrError(t)
	// Output:
	// ✘ // there are functions planned but not called		--- @ ./mocker/scenario/scenario_test.go:356
	//     ✘ ./mocker/scenario/scenario_test.go:338
}
