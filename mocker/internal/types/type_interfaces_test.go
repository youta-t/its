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
	var testee types.I0 = gen_mock.I0_Build(t, gen_mock.I0_Spec{
		M0: gen_mock.I0_M0_Expects().ThenReturn(),
		M1: gen_mock.I0_M1_Expects(its.EqEq(42), its.EqEq("abc")).
			ThenReturn(true, wantErr),
		M2: gen_mock.I0_M2_Expects(its.EqEq(42), its.EqEq("abc")).
			ThenReturn(true, wantErr),
		M3: gen_mock.I0_M3_Expects(
			its.EqEq(42),
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn(true, wantErr),
		M4: gen_mock.I0_M4_Expects(
			its.ForItems(its.Slice, its.EqEq, []string{"a", "b", "c"}),
		).
			ThenReturn(true),
	})

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
	testee := gen_mock.I1_Build(t, gen_mock.I1_Spec[sub.T, sub.T, types.X[sub.T]]{
		M0: gen_mock.I1_M0_Expects(its.EqEq(sub.T{}), its.EqEq(sub.T{})).
			ThenReturn(types.X[sub.T]{}, wanterr),
		M1: gen_mock.I1_M1_Expects(its.EqEq(sub.T{})).ThenReturn(),
		M2: gen_mock.I1_M2_Expects().ThenReturn(types.X[sub.T]{}),
	})
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

	var testee types.I2[mystr] = gen_mock.I2_Build(t, gen_mock.I2_Spec[mystr]{
		M0: gen_mock.I2_M0_Expects(its.EqEq(mystr("foo"))).
			ThenReturn(mystr("bar")),
	})

	r1 := testee.M0("foo")
	its.EqEq(mystr("bar")).Match(r1).OrError(t)
}
