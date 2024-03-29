package types_test

import (
	"testing"

	"github.com/youta-t/its"
	types "github.com/youta-t/its/mocker/internal/types"
	"github.com/youta-t/its/mocker/internal/types/gen_mock"
)

// This test file tests not only behaviour but also compilability.

func TestF1(t *testing.T) {
	// compilability test
	{
		testee := gen_mock.NewF1Call().ThenReturn().Mock(t)
		testee()
	}

	{
		testee := gen_mock.NewF1Call().ThenEffect(func() {}).Mock(t)
		testee()
	}
}

func TestF2(t *testing.T) {
	{
		testee := gen_mock.NewF2Call().ThenReturn(42).Mock(t)
		its.EqEq(42).Match(testee()).OrError(t)
	}

	{
		testee := gen_mock.NewF2Call().
			ThenEffect(func() int { return 42 }).
			Mock(t)
		its.EqEq(42).Match(testee()).OrError(t)
	}
}

func TestF3(t *testing.T) {
	{
		testee := gen_mock.NewF3Call().ThenReturn(1, "a").Mock(t)

		r1, r2 := testee()

		its.EqEq(r1).Match(1).OrError(t)
		its.EqEq(r2).Match("a").OrError(t)
	}

	{
		testee := gen_mock.NewF3Call().ThenEffect(func() (int, string) {
			return 1, "a"
		}).Mock(t)

		r1, r2 := testee()
		its.EqEq(r1).Match(1).OrError(t)
		its.EqEq(r2).Match("a").OrError(t)
	}
}

func TestF4(t *testing.T) {
	{
		testee := gen_mock.NewF4Call().ThenReturn(1, "a").Mock(t)

		r1, r2 := testee()

		its.EqEq(r1).Match(1).OrError(t)
		its.EqEq(r2).Match("a").OrError(t)
	}

	{
		testee := gen_mock.NewF4Call().ThenEffect(func() (int, string) {
			return 1, "a"
		}).Mock(t)

		r1, r2 := testee()
		its.EqEq(r1).Match(1).OrError(t)
		its.EqEq(r2).Match("a").OrError(t)
	}
}

func TestF5(t *testing.T) {
	{
		testee := gen_mock.NewF5Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenReturn().
			Mock(t)

		testee(15, "abc")
	}
	{
		testee := gen_mock.NewF5Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenEffect(func(arg0 int, arg1 string) {}).
			Mock(t)

		testee(15, "abc")
	}
}

func TestF6(t *testing.T) {
	{
		testee := gen_mock.NewF6Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee(42, "a", "b", "c")
	}
	{
		testee := gen_mock.NewF6Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(arg0 int, vararg ...string) {}).
			Mock(t)
		testee(42, "a", "b", "c")
	}
}

func TestF7(t *testing.T) {
	{
		testee := gen_mock.NewF7Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee("a", "b", "c")
	}
	{
		testee := gen_mock.NewF7Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(vararg ...string) {}).
			Mock(t)
		testee("a", "b", "c")
	}
}

func TestF8(t *testing.T) {
	{
		testee := gen_mock.NewF8Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenReturn().
			Mock(t)

		testee(15, "abc")
	}
	{
		testee := gen_mock.NewF8Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenEffect(func(arg0 int, arg1 string) {}).
			Mock(t)

		testee(15, "abc")
	}
}

func TestF9(t *testing.T) {
	{
		testee := gen_mock.NewF9Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee(42, "a", "b", "c")
	}
	{
		testee := gen_mock.NewF9Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(arg0 int, vararg ...string) {}).
			Mock(t)
		testee(42, "a", "b", "c")
	}
}

func TestF10(t *testing.T) {
	{
		testee := gen_mock.NewF10Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee("a", "b", "c")
	}
	{
		testee := gen_mock.NewF10Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(vararg ...string) {}).
			Mock(t)
		testee("a", "b", "c")
	}
}

func TestF11(t *testing.T) {
	{
		testee := gen_mock.NewF11Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenReturn().
			Mock(t)

		testee(15, "abc")
	}
	{
		testee := gen_mock.NewF11Call(
			its.EqEq(15), its.EqEq("abc"),
		).
			ThenEffect(func(arg0 int, arg1 string) {}).
			Mock(t)

		testee(15, "abc")
	}
}

func TestF12(t *testing.T) {
	{
		testee := gen_mock.NewF12Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee(42, "a", "b", "c")
	}
	{
		testee := gen_mock.NewF12Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(arg0 int, vararg ...string) {}).
			Mock(t)
		testee(42, "a", "b", "c")
	}
}

func TestF13(t *testing.T) {
	{
		testee := gen_mock.NewF13Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn().
			Mock(t)
		testee("a", "b", "c")
	}
	{
		testee := gen_mock.NewF13Call(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(vararg ...string) {}).
			Mock(t)
		testee("a", "b", "c")
	}
}

func TestF14(t *testing.T) {
	{
		testee := gen_mock.NewF14Call(
			its.EqEq(42), its.EqEq("foo"),
		).
			ThenReturn(12.125).
			Mock(t)

		r1 := testee(42, "foo")

		its.EqEq(12.125).Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF14Call(
			its.EqEq(42), its.EqEq("foo"),
		).
			ThenEffect(func(arg0 int, arg1 string) (f float64) {
				return 12.125
			}).
			Mock(t)

		r1 := testee(42, "foo")

		its.EqEq(12.125).Match(r1).OrError(t)
	}
}

func TestF15(t *testing.T) {
	{
		testee := gen_mock.NewF15Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn(12.125).
			Mock(t)

		r1 := testee(42, "a", "b", "c")

		its.EqEq(12.125).Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF15Call(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenEffect(func(i int, ss ...string) float64 {
				return 12.125
			}).
			Mock(t)

		r1 := testee(42, "a", "b", "c")

		its.EqEq(12.125).Match(r1).OrError(t)
	}
}

func TestF16(t *testing.T) {
	{
		testee := gen_mock.NewF16Call[string](
			its.Always[types.F16[string]](),
		).
			ThenReturn("foo").
			Mock(t)

		r1 := testee(func(f types.F16[string]) string {
			return "abc"
		})

		its.EqEq("foo").Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF16Call[string](
			its.Always[types.F16[string]](),
		).
			ThenEffect(func(arg0 types.F16[string]) string {
				return "foo"
			}).
			Mock(t)

		r1 := testee(func(f types.F16[string]) string {
			return "abc"
		})

		its.EqEq("foo").Match(r1).OrError(t)
	}
}

func TestF17(t *testing.T) {
	type Fooer struct {
		Foo int
	}
	{
		testee := gen_mock.NewF17Call[Fooer](
			its.EqEq(Fooer{Foo: 42}),
		).
			ThenReturn(Fooer{Foo: 99}).
			Mock(t)

		r1 := testee(Fooer{Foo: 42})
		its.EqEq(Fooer{Foo: 99}).Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF17Call[Fooer](
			its.EqEq(Fooer{Foo: 42}),
		).
			ThenEffect(func(arg0 Fooer) Fooer {
				return Fooer{Foo: 99}
			}).
			Mock(t)

		r1 := testee(Fooer{Foo: 42})
		its.EqEq(Fooer{Foo: 99}).Match(r1).OrError(t)
	}
}

func TestF18(t *testing.T) {
	// compilability test
	{
		testee := gen_mock.NewF18Call[int, string](
			its.Always[func(int, string)](),
		).
			ThenReturn(func(int, string) {}).
			Mock(t)

		r1 := testee(func(int, string) {})
		r1(1, "a")
	}
	{
		testee := gen_mock.NewF18Call[int, string](
			its.Always[func(int, string)](),
		).
			ThenEffect(func(arg0 func(arg0 int, arg1 string)) func(arg0 int, arg1 string) {
				return func(arg0 int, arg1 string) {}
			}).
			Mock(t)

		r1 := testee(func(int, string) {})
		r1(1, "a")
	}
}

func TestF19(t *testing.T) {
	type myint int
	{
		testee := gen_mock.NewF19Call[myint](
			its.Always[types.F16[myint]](),
		).
			ThenReturn(myint(16)).
			Mock(t)

		r1 := testee(func(f types.F16[myint]) myint { return 42 })
		its.EqEq(myint(16)).Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF19Call[myint](
			its.Always[types.F16[myint]](),
		).
			ThenEffect(func(arg0 types.F16[myint]) myint {
				return 16
			}).
			Mock(t)

		r1 := testee(func(f types.F16[myint]) myint { return 42 })
		its.EqEq(myint(16)).Match(r1).OrError(t)
	}
}

func TestF20(t *testing.T) {
	{
		testee := gen_mock.NewF20Call(its.EqEq(42)).ThenReturn(99).Mock(t)
		r1 := testee(42)
		its.EqEq(99).Match(r1).OrError(t)
	}
	{
		testee := gen_mock.NewF20Call(its.EqEq(42)).
			ThenEffect(func(arg0 int) int { return 99 }).
			Mock(t)
		r1 := testee(42)
		its.EqEq(99).Match(r1).OrError(t)
	}
}
