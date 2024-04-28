package generatetest_test

import (
	"errors"
	"testing"

	"github.com/youta-t/its"
	"github.com/youta-t/its/mocker/internal/example/sub"
	types "github.com/youta-t/its/mocker/internal/generate_test"
	"github.com/youta-t/its/mocker/internal/generate_test/gen_mock"
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

func TestC3(t *testing.T) {
	readErr := errors.New("read error")
	writeErr := errors.New("write err")
	var testee types.C3 = gen_mock.C3_Build(t, gen_mock.C3_Spec{
		Read:  gen_mock.C3_Read_Expects(its.Not(its.Nil[[]byte]())).ThenReturn(7, readErr),
		Write: gen_mock.C3_Write_Expects(its.Not(its.Nil[[]byte]())).ThenReturn(42, writeErr),
	})

	{
		buf := make([]byte, 3)
		n, err := testee.Read(buf)
		its.EqEq(7).Match(n).OrError(t)
		its.Error(readErr).Match(err).OrError(t)
	}

	{
		buf := make([]byte, 3)
		n, err := testee.Write(buf)
		its.EqEq(42).Match(n).OrError(t)
		its.Error(writeErr).Match(err).OrError(t)
	}
}

func TestC4(t *testing.T) {
	readErr := errors.New("read err")
	closeErr := errors.New("close err")
	testee := gen_mock.C4_Build(t, gen_mock.C4_Spec{
		Read:          gen_mock.C4_Read_Expects(its.Not(its.Nil[[]byte]())).ThenReturn(10, readErr),
		Close:         gen_mock.C4_Close_Expects().ThenReturn(closeErr),
		Method:        gen_mock.C4_Method_Expects().ThenReturn(),
		AnotherMethod: gen_mock.C4_AnotherMethod_Expects().ThenReturn(),
		DotMethod:     gen_mock.C4_DotMethod_Expects().ThenReturn(),
	})

	{
		buf := make([]byte, 3)
		n, err := testee.Read(buf)
		its.EqEq(10).Match(n).OrError(t)
		its.EqEq(readErr).Match(err).OrError(t)
	}
	{
		err := testee.Close()
		its.Error(closeErr).Match(err).OrError(t)
	}
	testee.Method()
	testee.DotMethod()
	testee.AnotherMethod()

}
