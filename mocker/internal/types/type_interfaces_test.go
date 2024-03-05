package types_test

import (
	"errors"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/mocker/internal/example/sub"
	types "github.com/youta-t/its/mocker/internal/types"
	"github.com/youta-t/its/mocker/internal/types/gen_mock"
)

func TestI0(t *testing.T) {
	wantErr := errors.New("err")
	var testee types.I0 = gen_mock.NewMockedI0(
		t,
		gen_mock.I0Impl{
			M0: gen_mock.NewI0_M0Call().ThenReturn().Mock(t),
			M1: gen_mock.NewI0_M1Call(its.EqEq(42), its.EqEq("abc")).
				ThenReturn(true, wantErr).
				Mock(t),
			M2: gen_mock.NewI0_M2Call(its.EqEq(42), its.EqEq("abc")).
				ThenReturn(true, wantErr).
				Mock(t),
			M3: gen_mock.NewI0_M3Call(
				its.EqEq(42),
				its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
			).
				ThenReturn(true, wantErr).
				Mock(t),
			M4: gen_mock.NewI0_M4Call(
				its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
			).
				ThenReturn(true).
				Mock(t),
		},
	)

	{
		testee.M0()
	}

	{
		r1, r2 := testee.M1(42, "abc")
		its.EqEq(true).Match(r1).OrError(t)
		its.Error(wantErr).Match(r2).OrError(t)
	}

	{
		r1, r2 := testee.M2(42, "abc")
		its.EqEq(true).Match(r1).OrError(t)
		its.Error(wantErr).Match(r2).OrError(t)
	}

	{
		r1, r2 := testee.M3(42, "a", "b", "c")
		its.EqEq(true).Match(r1).OrError(t)
		its.Error(wantErr).Match(r2).OrError(t)
	}

	{
		r1 := testee.M4("a", "b", "c")
		its.EqEq(true).Match(r1).OrError(t)
	}
}

func TestI1(t *testing.T) {
	wanterr := errors.New("err")
	testee := gen_mock.NewMockedI1(
		t,
		gen_mock.I1Impl[sub.T, sub.T, types.X[sub.T]]{
			M0: gen_mock.NewI1_M0Call(its.EqEq(sub.T{}), its.EqEq(sub.T{})).
				ThenReturn(types.X[sub.T]{}, wanterr).
				Mock(t),
			M1: gen_mock.NewI1_M1Call(its.EqEq(sub.T{})).ThenReturn().Mock(t),
			M2: gen_mock.NewI1_M2Call().ThenReturn(types.X[sub.T]{}).Mock(t),
		},
	)
	{
		r1, r2 := testee.M0(sub.T{}, sub.T{})

		its.EqEq(types.X[sub.T]{}).Match(r1).OrError(t)
		its.Error(wanterr).Match(r2).OrError(t)
	}

	{
		testee.M1(sub.T{})
	}

	{
		r1 := testee.M2()
		its.EqEq(types.X[sub.T]{}).Match(r1).OrError(t)
	}
}

func TestI3(t *testing.T) {
	type mystr string

	var testee types.I2[mystr] = gen_mock.NewMockedI2(t, gen_mock.I2Impl[mystr]{
		M0: gen_mock.NewI2_M0Call(its.EqEq(mystr("foo"))).
			ThenReturn(mystr("bar")).
			Mock(t),
	})

	r1 := testee.M0("foo")
	its.EqEq(mystr("bar")).Match(r1).OrError(t)
}
