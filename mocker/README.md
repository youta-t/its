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

When you do `go generate ./...`, its-mocker creates mock function builders and
mock implementation of interfaces.

Function Mocks of its
-------------

Its has an opinion about mocks.

When mocks are needed, they mostly are used for

- checking arguments are fine, and
- giving fixed return values.

So, its mock has such features.

The its-mocker generates `New${func name}Call` function for each functions and methods of interfaces.
This `New...Call` is the entrypoint of function mock builder, and receives "what arguments to be passed" as `its.Matcher`

For example,

```go
mockFn := NewFooCall(its.EqEq(10), its.GreaterThan(0)).
    // ...
```

declares, the mock function is expected to be called as `mockFn(10, 1)` or `mockFn(10, 100)` (or something).

This mock automatically test arguments when called, no needs to retreive arguments later.
If not the test is passed, it records test error.

Then, let us declare return values and get a mock.

```go
mockFn := NewFooCall(its.EqEq(10), its.GreaterThan(0)).
    ThenReturn(42, false).
    Mock(t)  // t is *testing.T
```

In this method chain, `ThenReturn` declares "what values should be returned".
When do that, `mockFn` returns `(42, false)` always.

And, `Mock(t)` builds the method chain into a mock function.

`New...Call` function is generated for each functions and methods of interfaces,
so arguments of `New...Call` and the method chain are tailored for each functions/methods.

### Mock with Behavior

When you need to give behavior to mock, use `.ThenEffect` for an alternative of `.ThenReturn`.

```go
mockFn := NewFooCall(its.EqEq(10), its.GreaterThan(0)).
    ThenEffect(func(int, int) (int, bool) {
        fmt.Printf("called!")
        return 42, false
    }).
    Mock(t)  // t is *testing.T
```

The return values of callback are the return values of the mock.

Interface Mocks
----------------

For each interface, its-mock generates mock implementations of them.

Mock implementation can be gain with `New${interface-name}(t, ${interface-name}Impl{ ... })`.

For example, given an interface `Fizz` as below:

```go
type FizzBazz interface {
    Say (int) (int, string, bool)
}
```

Then, with its-mocker, you can mock of that with

```go
mock := NewFizzBazz(t, FizzImpl{  // t is *testing.T
    Say: func(n int) (int, string, bool) { ... },
})
```

Fields of `FizzImpl` defines mocked implementations for each methods.
Of course, to create mocked methods we can use function mock. So, we can...

```go
mock := NewFizzBazz(t, FizzImpl{  // t is *testing.T
    Say: NewFizzBazz_SayCall(its.EqEq(3)).ThenReturn(3, "fizz", true).Mock(t),
})

mock.Say(3)  // 3, "fizz", true
```

If you call not-mocked method, the interface mock records test error.

Scenario Test
-------------

When mocking, mock is not enough.

Testing with mock are often intersted with call order of mocked functions.
For such case, its mock supports "scenario test".

Scenario test feature is exported from `github.com/youta-t/its/mocker/scenario`.

In scenario test, the interested call order is described as a scenario.

```go
sc := scenario.Begin(t)  // t is *testing.T
defer sc.End()

_, chap1 := scenario.Next(sc, func(...) { ... })
_, chap2 := scenario.Next(sc, func(...) { ... })
// ...
```

At first, `scenario.Begin`. This creates an empty scenario.
Then, register functions with `scenario.Next` in expected call order.

Returned `chapX`s are functions tracked by scenario.
If `chap`s are called out of order, scenario reports test errors.
If all `chap` are not called until `sc.End`, also scenario reports test errors.

So, with scenario, "it has been called?" flags are not needed. Scenario tracks that.

`scenario.Next` can receive function, so we can pass function-mock.
And, returned `chap` can be set as implementations of interface.

So, building them up all and we get...

```go
sc := scenario.Begin(t)  // t is *testing.T
defer sc.End()

_, chap1 := scenario.Next(
    sc,
    NewFizzBazz_SayCall(its.EqEq(3)).
        ThenReturn(3, "fizz", true).
        Mock(t),
)

mock := NewFizzBazz(t, FizzBazzImpl{
    Say: chap1,
})

mock.Say(3)
```

This tests them all:

- `mock.Say` is called? (by scenario)
- `mock.Say` is planned to be called? (by interface mock)
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
	sc := scenario.Begin(t)
	defer sc.End()

	_, sess := scenario.Next( // add a function into scenario
		sc, NewSessionStoreCall(
			its.EqEq("fake-cookie")).           // expectation for arguments
			ThenReturn("sample-user-id", true). // fixture for retrun valeus
			Mock(t),                            // build as mock function
	)

	_, getUser := scenario.Next(sc, NewUserRegistry_GetCall(
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

	_, updateUser := scenario.Next(sc, NewUserRegistry_UpdateCall(
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

	its.Nil().Match( // no error has been returned?
		testee("fake-cookie", "Richard Roe"),
	).OrError(t)
}
```

Concusion
-----------

- its-mock generates function-mock and interface-mock, they tests for each layer.
- its has scenario test feature, it tests calling order of functions.
- its features are composable.
