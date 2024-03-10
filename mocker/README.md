its-mocker
============

its-mocker is a code generator to create mocks for interface or function types.

```
Usage of .../mocker:
mocker is a mock generator for any interfaces or functions.
This is designed to be used as go:generate.

It generates a file with same name as a file having go:generate directive.

  -as-package
        handle -source as package path
  -dest string
        directory where new file to be created at (default "./gen_mock")
  -p    alias of -as-package
  -source string
        recognise source as package path. If not set, use environmental variable GOFILE.
  -t value
        alias of -type
  -type value
        Type names to generate Mock. Repeatable. By default, all interfaces and type func are target.
```

Typical Usage
-------------

its-mocker is designed to be used as 

```go
//go:generate go run github.com/youta-t/its/mocker
```

Generated files are stored in a subpackage named `gen_mock`.

When you do `go generate ./...`, its-mocker creates mock function builders and
mock implementation of interfaces.


Function Mocks of its
-------------

Its has an opinion about mocks.

When mocks are needed, they mostly are used for

- checking arguments are fine, and
- giving fixed return values.

So, its mock has such features.

The its-mocker generates `${func name}_Expecs` function for each functions and methods of interfaces.
This `..._Expects` is the entrypoint of function mock builder, and receives "what arguments to be passed" as `its.Matcher`

`..._Expects` function is generated for each functions and methods of interfaces,
so arguments of `..._Expects` and the method chain are tailored for each functions/methods.

For example about a type

```go
//go:generate github.com/youta-t/its/mocker

// ...

type Foo func(int, int) (int, bool)
```

After `go generate ./...`, we can create mock of a `Foo` function.

```go
mockFoo := Foo_Expects(its.EqEq(10), its.GreaterThan(0)).
    // cont....
```

It declares, the mock function is expected to be called as `mockFoo(10, 1)` or `mockFoo(10, 100)` (or something).

This mock automatically test arguments when called, no needs to retreive arguments later.
If not the test is passed, it records test error.

Then, let us declare return values and get a mock.

```go
mockFoo := Foo_Expects(its.EqEq(10), its.GreaterThan(0)).
    ThenReturn(42, false).
    // cont....
```

In this method chain, `.ThenReturn` declares "what values should be returned".
When do that, the mock function `mockFoo` returns `(42, false)` always.

`.ThenReturn` at here returns `its/mocker/mockkit.FuncBehavior[Foo]`, not a function itself.
To get a mocked function, do

```go
mockFoo := Foo_Expects(its.EqEq(10), its.GreaterThan(0)).
    ThenReturn(42, false).
    Fn(t)  // t is *testing.T.
```

Now we get a mock.

### Mock with Side Effect

When you need to give side effects to mock, use `.ThenEffect` for an alternative of `.ThenReturn`.

```go
mockFoo := Foo_Expects(its.EqEq(10), its.GreaterThan(0)).
    ThenEffect(func(int, int) (int, bool) {
        fmt.Printf("called!")
        return 42, false
    }).
    Fn(t)  // t is *testing.T
```

The return values of callback are the return values of the mock.

Interface Mocks
----------------

For each interface, its/mocker generates mock implementations of them.

Mock implementation can be gain with `${interface name}_Build(t, ${interface name}_Spec{ ... })`.

For example, given an interface `FizzBazz` as below:

```go
type FizzBazz interface {
    Say (int) (n int, s string, seeString bool)
}
```

Then, with its/mocker, you can mock of that with

```go
var mock FizzBazz = FizzBazz_Build(
    t,  // t is *testing.T
    FizzBazz_Spec{
        Say: FizzBazz_Say_Expects(...).ThenReturn(...),
    },
)

mock.Say(3)  // 3, "fizz", true
```

Fields of `FizzBazz_Spec` holds "behavior" for each methods.
Behavior, in this context, is a value defines that "how does it test the parameters?" and "what values will be returned?".

They, fields, are typed as `github.com/youta.t/its/mocker/mockkit.FuncBehavior[F]`, of course it represents a behavior.
Each of `F`s is a `func` type having same signature as the method for the field name,
and it expresses specification of the behavior the mocked funcion.

When you call not-mocked method of interface mock, the interface mock records test error.

### Behavior without parameter testing

Test-free `FuncBehavior[F]` can be created by `github.com/youta-t/its/mocker/mockkit.Effect`.

The package `github.com/yotua-t/mocker/mockkit` has utilities for mockking. `Effect` function is one of them.

`Effect` works like below:

```go
var eff mockkit.FuncBehavior[func(int, string) (float, error)] = mockkit.Effect(
    func(a int, b string) (f float, err error) { return 0, nil },
)
mockFn := eff.Fn(t)
mockFn()  // 0, nil
```

Scenario Test
-------------

When mocking, mock is not enough.

Testing with mock are often intersted with call order of mocked functions.
For such case, its mock supports "scenario test".

In scenario test, the interested call order is described as a scenario.

```go
sc := mockkit.BeginScenario(t)  // t is *testing.T
defer sc.End()

chap1 := mockkit.Next(sc, mockkit.Effect(func(...) { ... }))
chap2 := mockkit.Next(sc, mockkit.Effect(func(...) { ... }))
// ...
```

At first, `mockkit.BeginScenario`. This creates an empty scenario.
Then, register `FuncBehavior[F]` with `mockkit.Next` in expected call order.

Returned `chapX`s are behaviors tracked by scenario.
If `chap`s are called out of order or have not called until the end, scenario reports test errors.

`mockkit.Next` can receive `FuncBehavior`, so we can pass function-mock.
And, returned `chap` can be set as implementations of interface.

So, building them up all and we get...

```go
sc := mockkit.BeginScenario(t)  // t is *testing.T
defer sc.End()

spec := FizzBazz_Sppec{}
spec.Say = mockkit.Next(
    sc,
    NewFizzBazz_SayCall(its.EqEq(3)).ThenReturn(3, "fizz", true),
)

mock := FizzBazz_Build(t, spec)
mock.Say(3)
```

This tests them all:

- `mock.Say` is called? (by scenario)
- `mock.Say` is planned to be called? (by interface mock & scenario)
- `mock.Say` is called with `(3)`? (by function mock)

More advanced example
----------------------

Thinking about a tiny web service like function.

```go
// assume them:

type User struct {
    Id   string
    Name string
}

type UserRegistry interface {
    Get(userId string) (User, error)
    Update(User) error
    Delete(User) error
}

type SessionStore func(cookie string) (userId string, ok bool)

// Let us test this web-like feature.

func UpdateUser(
    sess SessionStore,
    registry UserRegistry,
) func(cookie string, newName string) error {
    return func(cookie, newName string) error {
        userId, ok := sess(cookie)
        if !ok {
            return errors.New("you are not logged in")
        }
        user, err := registry.Get(userId)
        if err != nil {
            return err
        }
        user.Name = newName
        return registry.Update(user)
    }
}
```

We should test `UpdateUser` feature.

This receives `SessionStore` (as function) and `UserRegistry` (as interface), and
returns "request handler".

To test that, we can do

```go
func TestUpdateUser(t *testing.T) {
	sc := mockkit.BeginScenario(t)
	defer sc.End()

	sess := mockkit.Next( // add a function into scenario
		sc,
		SessionStore_Expects(its.EqEq("fake-cookie")).  // expectation for arguments
			ThenReturn("sample-user-id", true),         // fixture for retrun valeus
	)

	userRegistry := UserRegistry_Spec{}
	userRegistry.Get = mockkit.Next(
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

	userRegistry.Update = mockkit.Next(
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
		UserRegistry_Build(t, userRegistry),
	)

	its.Nil().
		Match(testee("fake-cookie", "Richard Roe")).  // no error has been returned?
		OrError(t)
}
```

Concusion
-----------

- its-mock generates function-mock and interface-mock, they tests for each layer.
- its has scenario test feature, it tests calling order of functions.
- its features are composable.
