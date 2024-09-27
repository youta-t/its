// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	pkg1 "github.com/youta-t/its/mocker/internal/generate_test"
	pkg2 "github.com/youta-t/its/mocker/internal/generate_test/dot"
	mockkit "github.com/youta-t/its/mocker/mockkit"
)

type _F1CallSpec struct {
}

type _F1Call struct {
	name itskit.Label
	spec _F1CallSpec
}

func F1_Expects() _F1Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F1Call{
		name: itskit.NewLabelWithLocation("func F1"),
		spec: _F1CallSpec{},
	}
}

type _F1Behavior struct {
	name   itskit.Label
	spec   _F1CallSpec
	effect func()
}

func (b *_F1Behavior) Fn(t mockkit.TestLike) func() {
	return func() {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect()
	}
}

func (c _F1Call) ThenReturn() mockkit.FuncBehavior[func()] {
	return c.ThenEffect(func() {

	})
}

func (c _F1Call) ThenEffect(effect func()) mockkit.FuncBehavior[func()] {
	return &_F1Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F2CallSpec struct {
}

type _F2Call struct {
	name itskit.Label
	spec _F2CallSpec
}

func F2_Expects() _F2Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F2Call{
		name: itskit.NewLabelWithLocation("func F2"),
		spec: _F2CallSpec{},
	}
}

type _F2Behavior struct {
	name   itskit.Label
	spec   _F2CallSpec
	effect func() int
}

func (b *_F2Behavior) Fn(t mockkit.TestLike) func() int {
	return func() int {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect()
	}
}

func (c _F2Call) ThenReturn(

	ret0 int,

) mockkit.FuncBehavior[func() int] {
	return c.ThenEffect(func() int {

		return ret0

	})
}

func (c _F2Call) ThenEffect(effect func() int) mockkit.FuncBehavior[func() int] {
	return &_F2Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F3CallSpec struct {
}

type _F3Call struct {
	name itskit.Label
	spec _F3CallSpec
}

func F3_Expects() _F3Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F3Call{
		name: itskit.NewLabelWithLocation("func F3"),
		spec: _F3CallSpec{},
	}
}

type _F3Behavior struct {
	name   itskit.Label
	spec   _F3CallSpec
	effect func() (int, string)
}

func (b *_F3Behavior) Fn(t mockkit.TestLike) func() (int, string) {
	return func() (
		int,
		string,

	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect()
	}
}

func (c _F3Call) ThenReturn(

	ret0 int,

	ret1 string,

) mockkit.FuncBehavior[func() (int, string)] {
	return c.ThenEffect(func() (
		int,
		string,

	) {

		return ret0, ret1

	})
}

func (c _F3Call) ThenEffect(effect func() (int, string)) mockkit.FuncBehavior[func() (int, string)] {
	return &_F3Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F4CallSpec struct {
}

type _F4Call struct {
	name itskit.Label
	spec _F4CallSpec
}

func F4_Expects() _F4Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F4Call{
		name: itskit.NewLabelWithLocation("func F4"),
		spec: _F4CallSpec{},
	}
}

type _F4Behavior struct {
	name   itskit.Label
	spec   _F4CallSpec
	effect func() (i int, s string)
}

func (b *_F4Behavior) Fn(t mockkit.TestLike) func() (i int, s string) {
	return func() (
		int,
		string,

	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect()
	}
}

func (c _F4Call) ThenReturn(

	ret0 int,

	ret1 string,

) mockkit.FuncBehavior[func() (i int, s string)] {
	return c.ThenEffect(func() (
		int,
		string,

	) {

		return ret0, ret1

	})
}

func (c _F4Call) ThenEffect(effect func() (i int, s string)) mockkit.FuncBehavior[func() (i int, s string)] {
	return &_F4Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F5CallSpec struct {
	arg0 its.Matcher[int]

	arg1 its.Matcher[string]
}

type _F5Call struct {
	name itskit.Label
	spec _F5CallSpec
}

func F5_Expects(
	arg0 its.Matcher[int],

	arg1 its.Matcher[string],

) _F5Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F5Call{
		name: itskit.NewLabelWithLocation("func F5"),
		spec: _F5CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),

			arg1: itskit.Named(
				"arg1",
				arg1,
			),
		},
	}
}

type _F5Behavior struct {
	name   itskit.Label
	spec   _F5CallSpec
	effect func(arg0 int, arg1 string)
}

func (b *_F5Behavior) Fn(t mockkit.TestLike) func(arg0 int, arg1 string) {
	return func(

		arg0 int,

		arg1 string,

	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.arg1
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg1)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			arg1,
		)
	}
}

func (c _F5Call) ThenReturn() mockkit.FuncBehavior[func(arg0 int, arg1 string)] {
	return c.ThenEffect(func(

		int,

		string,

	) {

	})
}

func (c _F5Call) ThenEffect(effect func(arg0 int, arg1 string)) mockkit.FuncBehavior[func(arg0 int, arg1 string)] {
	return &_F5Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F6CallSpec struct {
	arg0 its.Matcher[int]

	vararg its.Matcher[[]string]
}

type _F6Call struct {
	name itskit.Label
	spec _F6CallSpec
}

func F6_Expects(
	arg0 its.Matcher[int],
	vararg its.Matcher[[]string],
) _F6Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F6Call{
		name: itskit.NewLabelWithLocation("func F6"),
		spec: _F6CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),

			vararg: itskit.Named(
				"vararg",
				vararg,
			),
		},
	}
}

type _F6Behavior struct {
	name   itskit.Label
	spec   _F6CallSpec
	effect func(arg0 int, vararg ...string)
}

func (b *_F6Behavior) Fn(t mockkit.TestLike) func(arg0 int, vararg ...string) {
	return func(

		arg0 int,

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.vararg
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			vararg...,
		)
	}
}

func (c _F6Call) ThenReturn() mockkit.FuncBehavior[func(arg0 int, vararg ...string)] {
	return c.ThenEffect(func(

		int,

		...string,
	) {

	})
}

func (c _F6Call) ThenEffect(effect func(arg0 int, vararg ...string)) mockkit.FuncBehavior[func(arg0 int, vararg ...string)] {
	return &_F6Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F7CallSpec struct {
	vararg its.Matcher[[]string]
}

type _F7Call struct {
	name itskit.Label
	spec _F7CallSpec
}

func F7_Expects(vararg its.Matcher[[]string],
) _F7Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F7Call{
		name: itskit.NewLabelWithLocation("func F7"),
		spec: _F7CallSpec{
			vararg: itskit.Named(
				"vararg",
				vararg,
			),
		},
	}
}

type _F7Behavior struct {
	name   itskit.Label
	spec   _F7CallSpec
	effect func(vararg ...string)
}

func (b *_F7Behavior) Fn(t mockkit.TestLike) func(vararg ...string) {
	return func(

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.vararg
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			vararg...,
		)
	}
}

func (c _F7Call) ThenReturn() mockkit.FuncBehavior[func(vararg ...string)] {
	return c.ThenEffect(func(

		...string,
	) {

	})
}

func (c _F7Call) ThenEffect(effect func(vararg ...string)) mockkit.FuncBehavior[func(vararg ...string)] {
	return &_F7Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F8CallSpec struct {
	cancel its.Matcher[int]

	spec its.Matcher[string]
}

type _F8Call struct {
	name itskit.Label
	spec _F8CallSpec
}

func F8_Expects(
	cancel its.Matcher[int],

	spec its.Matcher[string],

) _F8Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F8Call{
		name: itskit.NewLabelWithLocation("func F8"),
		spec: _F8CallSpec{
			cancel: itskit.Named(
				"cancel",
				cancel,
			),

			spec: itskit.Named(
				"spec",
				spec,
			),
		},
	}
}

type _F8Behavior struct {
	name   itskit.Label
	spec   _F8CallSpec
	effect func(cancel int, spec string)
}

func (b *_F8Behavior) Fn(t mockkit.TestLike) func(cancel int, spec string) {
	return func(

		arg0 int,

		arg1 string,

	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.cancel
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.spec
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg1)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			arg1,
		)
	}
}

func (c _F8Call) ThenReturn() mockkit.FuncBehavior[func(cancel int, spec string)] {
	return c.ThenEffect(func(

		int,

		string,

	) {

	})
}

func (c _F8Call) ThenEffect(effect func(cancel int, spec string)) mockkit.FuncBehavior[func(cancel int, spec string)] {
	return &_F8Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F9CallSpec struct {
	i its.Matcher[int]

	ss its.Matcher[[]string]
}

type _F9Call struct {
	name itskit.Label
	spec _F9CallSpec
}

func F9_Expects(
	i its.Matcher[int],
	ss its.Matcher[[]string],
) _F9Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F9Call{
		name: itskit.NewLabelWithLocation("func F9"),
		spec: _F9CallSpec{
			i: itskit.Named(
				"i",
				i,
			),

			ss: itskit.Named(
				"ss",
				ss,
			),
		},
	}
}

type _F9Behavior struct {
	name   itskit.Label
	spec   _F9CallSpec
	effect func(i int, ss ...string)
}

func (b *_F9Behavior) Fn(t mockkit.TestLike) func(i int, ss ...string) {
	return func(

		arg0 int,

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.i
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.ss
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			vararg...,
		)
	}
}

func (c _F9Call) ThenReturn() mockkit.FuncBehavior[func(i int, ss ...string)] {
	return c.ThenEffect(func(

		int,

		...string,
	) {

	})
}

func (c _F9Call) ThenEffect(effect func(i int, ss ...string)) mockkit.FuncBehavior[func(i int, ss ...string)] {
	return &_F9Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F10CallSpec struct {
	ss its.Matcher[[]string]
}

type _F10Call struct {
	name itskit.Label
	spec _F10CallSpec
}

func F10_Expects(ss its.Matcher[[]string],
) _F10Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F10Call{
		name: itskit.NewLabelWithLocation("func F10"),
		spec: _F10CallSpec{
			ss: itskit.Named(
				"ss",
				ss,
			),
		},
	}
}

type _F10Behavior struct {
	name   itskit.Label
	spec   _F10CallSpec
	effect func(ss ...string)
}

func (b *_F10Behavior) Fn(t mockkit.TestLike) func(ss ...string) {
	return func(

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.ss
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			vararg...,
		)
	}
}

func (c _F10Call) ThenReturn() mockkit.FuncBehavior[func(ss ...string)] {
	return c.ThenEffect(func(

		...string,
	) {

	})
}

func (c _F10Call) ThenEffect(effect func(ss ...string)) mockkit.FuncBehavior[func(ss ...string)] {
	return &_F10Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F11CallSpec struct {
	arg0 its.Matcher[int]

	arg1 its.Matcher[string]
}

type _F11Call struct {
	name itskit.Label
	spec _F11CallSpec
}

func F11_Expects(
	arg0 its.Matcher[int],

	arg1 its.Matcher[string],

) _F11Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F11Call{
		name: itskit.NewLabelWithLocation("func F11"),
		spec: _F11CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),

			arg1: itskit.Named(
				"arg1",
				arg1,
			),
		},
	}
}

type _F11Behavior struct {
	name   itskit.Label
	spec   _F11CallSpec
	effect func(arg0 int, arg1 string)
}

func (b *_F11Behavior) Fn(t mockkit.TestLike) func(arg0 int, arg1 string) {
	return func(

		arg0 int,

		arg1 string,

	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.arg1
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg1)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			arg1,
		)
	}
}

func (c _F11Call) ThenReturn() mockkit.FuncBehavior[func(arg0 int, arg1 string)] {
	return c.ThenEffect(func(

		int,

		string,

	) {

	})
}

func (c _F11Call) ThenEffect(effect func(arg0 int, arg1 string)) mockkit.FuncBehavior[func(arg0 int, arg1 string)] {
	return &_F11Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F12CallSpec struct {
	arg0 its.Matcher[int]

	vararg its.Matcher[[]string]
}

type _F12Call struct {
	name itskit.Label
	spec _F12CallSpec
}

func F12_Expects(
	arg0 its.Matcher[int],
	vararg its.Matcher[[]string],
) _F12Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F12Call{
		name: itskit.NewLabelWithLocation("func F12"),
		spec: _F12CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),

			vararg: itskit.Named(
				"vararg",
				vararg,
			),
		},
	}
}

type _F12Behavior struct {
	name   itskit.Label
	spec   _F12CallSpec
	effect func(arg0 int, vararg ...string)
}

func (b *_F12Behavior) Fn(t mockkit.TestLike) func(arg0 int, vararg ...string) {
	return func(

		arg0 int,

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.vararg
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			arg0,

			vararg...,
		)
	}
}

func (c _F12Call) ThenReturn() mockkit.FuncBehavior[func(arg0 int, vararg ...string)] {
	return c.ThenEffect(func(

		int,

		...string,
	) {

	})
}

func (c _F12Call) ThenEffect(effect func(arg0 int, vararg ...string)) mockkit.FuncBehavior[func(arg0 int, vararg ...string)] {
	return &_F12Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F13CallSpec struct {
	vararg its.Matcher[[]string]
}

type _F13Call struct {
	name itskit.Label
	spec _F13CallSpec
}

func F13_Expects(vararg its.Matcher[[]string],
) _F13Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F13Call{
		name: itskit.NewLabelWithLocation("func F13"),
		spec: _F13CallSpec{
			vararg: itskit.Named(
				"vararg",
				vararg,
			),
		},
	}
}

type _F13Behavior struct {
	name   itskit.Label
	spec   _F13CallSpec
	effect func(vararg ...string)
}

func (b *_F13Behavior) Fn(t mockkit.TestLike) func(vararg ...string) {
	return func(

		vararg ...string,
	) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.vararg
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(

			vararg...,
		)
	}
}

func (c _F13Call) ThenReturn() mockkit.FuncBehavior[func(vararg ...string)] {
	return c.ThenEffect(func(

		...string,
	) {

	})
}

func (c _F13Call) ThenEffect(effect func(vararg ...string)) mockkit.FuncBehavior[func(vararg ...string)] {
	return &_F13Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F14CallSpec struct {
	arg0 its.Matcher[int]

	arg1 its.Matcher[string]
}

type _F14Call struct {
	name itskit.Label
	spec _F14CallSpec
}

func F14_Expects(
	arg0 its.Matcher[int],

	arg1 its.Matcher[string],

) _F14Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F14Call{
		name: itskit.NewLabelWithLocation("func F14"),
		spec: _F14CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),

			arg1: itskit.Named(
				"arg1",
				arg1,
			),
		},
	}
}

type _F14Behavior struct {
	name   itskit.Label
	spec   _F14CallSpec
	effect func(arg0 int, arg1 string) (f float64)
}

func (b *_F14Behavior) Fn(t mockkit.TestLike) func(arg0 int, arg1 string) (f float64) {
	return func(

		arg0 int,

		arg1 string,

	) float64 {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.arg1
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg1)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,

			arg1,
		)
	}
}

func (c _F14Call) ThenReturn(

	ret0 float64,

) mockkit.FuncBehavior[func(arg0 int, arg1 string) (f float64)] {
	return c.ThenEffect(func(

		int,

		string,

	) float64 {

		return ret0

	})
}

func (c _F14Call) ThenEffect(effect func(arg0 int, arg1 string) (f float64)) mockkit.FuncBehavior[func(arg0 int, arg1 string) (f float64)] {
	return &_F14Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F15CallSpec struct {
	i its.Matcher[int]

	ss its.Matcher[[]string]
}

type _F15Call struct {
	name itskit.Label
	spec _F15CallSpec
}

func F15_Expects(
	i its.Matcher[int],
	ss its.Matcher[[]string],
) _F15Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F15Call{
		name: itskit.NewLabelWithLocation("func F15"),
		spec: _F15CallSpec{
			i: itskit.Named(
				"i",
				i,
			),

			ss: itskit.Named(
				"ss",
				ss,
			),
		},
	}
}

type _F15Behavior struct {
	name   itskit.Label
	spec   _F15CallSpec
	effect func(i int, ss ...string) float64
}

func (b *_F15Behavior) Fn(t mockkit.TestLike) func(i int, ss ...string) float64 {
	return func(

		arg0 int,

		vararg ...string,
	) float64 {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.i
			if matcher == nil {
				matcher = its.Never[int]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		{
			matcher := b.spec.ss
			if matcher == nil {
				matcher = its.Never[[]string]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,

			vararg...,
		)
	}
}

func (c _F15Call) ThenReturn(

	ret0 float64,

) mockkit.FuncBehavior[func(i int, ss ...string) float64] {
	return c.ThenEffect(func(

		int,

		...string,
	) float64 {

		return ret0

	})
}

func (c _F15Call) ThenEffect(effect func(i int, ss ...string) float64) mockkit.FuncBehavior[func(i int, ss ...string) float64] {
	return &_F15Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F16CallSpec[T any] struct {
	arg0 its.Matcher[pkg1.F16[T]]
}

type _F16Call[T any] struct {
	name itskit.Label
	spec _F16CallSpec[T]
}

func F16_Expects[T any](
	arg0 its.Matcher[pkg1.F16[T]],

) _F16Call[T] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F16Call[T]{
		name: itskit.NewLabelWithLocation("func F16"),
		spec: _F16CallSpec[T]{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F16Behavior[T any] struct {
	name   itskit.Label
	spec   _F16CallSpec[T]
	effect func(arg0 pkg1.F16[T]) T
}

func (b *_F16Behavior[T]) Fn(t mockkit.TestLike) func(arg0 pkg1.F16[T]) T {
	return func(

		arg0 pkg1.F16[T],

	) T {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[pkg1.F16[T]]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F16Call[T]) ThenReturn(

	ret0 T,

) mockkit.FuncBehavior[func(arg0 pkg1.F16[T]) T] {
	return c.ThenEffect(func(

		pkg1.F16[T],

	) T {

		return ret0

	})
}

func (c _F16Call[T]) ThenEffect(effect func(arg0 pkg1.F16[T]) T) mockkit.FuncBehavior[func(arg0 pkg1.F16[T]) T] {
	return &_F16Behavior[T]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F17CallSpec[T ~struct {
	Foo int
}] struct {
	arg0 its.Matcher[T]
}

type _F17Call[T ~struct {
	Foo int
}] struct {
	name itskit.Label
	spec _F17CallSpec[T]
}

func F17_Expects[T ~struct {
	Foo int
}](
	arg0 its.Matcher[T],

) _F17Call[T] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F17Call[T]{
		name: itskit.NewLabelWithLocation("func F17"),
		spec: _F17CallSpec[T]{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F17Behavior[T ~struct {
	Foo int
}] struct {
	name   itskit.Label
	spec   _F17CallSpec[T]
	effect func(arg0 T) T
}

func (b *_F17Behavior[T]) Fn(t mockkit.TestLike) func(arg0 T) T {
	return func(

		arg0 T,

	) T {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[T]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F17Call[T]) ThenReturn(

	ret0 T,

) mockkit.FuncBehavior[func(arg0 T) T] {
	return c.ThenEffect(func(

		T,

	) T {

		return ret0

	})
}

func (c _F17Call[T]) ThenEffect(effect func(arg0 T) T) mockkit.FuncBehavior[func(arg0 T) T] {
	return &_F17Behavior[T]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F18CallSpec[T any, U any] struct {
	arg0 its.Matcher[func(arg0 T, arg1 U)]
}

type _F18Call[T any, U any] struct {
	name itskit.Label
	spec _F18CallSpec[T, U]
}

func F18_Expects[T any, U any](
	arg0 its.Matcher[func(arg0 T, arg1 U)],

) _F18Call[T, U] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F18Call[T, U]{
		name: itskit.NewLabelWithLocation("func F18"),
		spec: _F18CallSpec[T, U]{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F18Behavior[T any, U any] struct {
	name   itskit.Label
	spec   _F18CallSpec[T, U]
	effect func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)
}

func (b *_F18Behavior[T, U]) Fn(t mockkit.TestLike) func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U) {
	return func(

		arg0 func(arg0 T, arg1 U),

	) func(arg0 T, arg1 U) {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[func(arg0 T, arg1 U)]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F18Call[T, U]) ThenReturn(

	ret0 func(arg0 T, arg1 U),

) mockkit.FuncBehavior[func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)] {
	return c.ThenEffect(func(

		func(arg0 T, arg1 U),

	) func(arg0 T, arg1 U) {

		return ret0

	})
}

func (c _F18Call[T, U]) ThenEffect(effect func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)) mockkit.FuncBehavior[func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)] {
	return &_F18Behavior[T, U]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F19CallSpec[T ~int] struct {
	arg0 its.Matcher[pkg1.F16[T]]
}

type _F19Call[T ~int] struct {
	name itskit.Label
	spec _F19CallSpec[T]
}

func F19_Expects[T ~int](
	arg0 its.Matcher[pkg1.F16[T]],

) _F19Call[T] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F19Call[T]{
		name: itskit.NewLabelWithLocation("func F19"),
		spec: _F19CallSpec[T]{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F19Behavior[T ~int] struct {
	name   itskit.Label
	spec   _F19CallSpec[T]
	effect func(arg0 pkg1.F16[T]) T
}

func (b *_F19Behavior[T]) Fn(t mockkit.TestLike) func(arg0 pkg1.F16[T]) T {
	return func(

		arg0 pkg1.F16[T],

	) T {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[pkg1.F16[T]]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F19Call[T]) ThenReturn(

	ret0 T,

) mockkit.FuncBehavior[func(arg0 pkg1.F16[T]) T] {
	return c.ThenEffect(func(

		pkg1.F16[T],

	) T {

		return ret0

	})
}

func (c _F19Call[T]) ThenEffect(effect func(arg0 pkg1.F16[T]) T) mockkit.FuncBehavior[func(arg0 pkg1.F16[T]) T] {
	return &_F19Behavior[T]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F20CallSpec[T ~int | int8 | ~int16 | int32] struct {
	arg0 its.Matcher[T]
}

type _F20Call[T ~int | int8 | ~int16 | int32] struct {
	name itskit.Label
	spec _F20CallSpec[T]
}

func F20_Expects[T ~int | int8 | ~int16 | int32](
	arg0 its.Matcher[T],

) _F20Call[T] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F20Call[T]{
		name: itskit.NewLabelWithLocation("func F20"),
		spec: _F20CallSpec[T]{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F20Behavior[T ~int | int8 | ~int16 | int32] struct {
	name   itskit.Label
	spec   _F20CallSpec[T]
	effect func(arg0 T) T
}

func (b *_F20Behavior[T]) Fn(t mockkit.TestLike) func(arg0 T) T {
	return func(

		arg0 T,

	) T {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[T]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F20Call[T]) ThenReturn(

	ret0 T,

) mockkit.FuncBehavior[func(arg0 T) T] {
	return c.ThenEffect(func(

		T,

	) T {

		return ret0

	})
}

func (c _F20Call[T]) ThenEffect(effect func(arg0 T) T) mockkit.FuncBehavior[func(arg0 T) T] {
	return &_F20Behavior[T]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F25CallSpec struct {
	d its.Matcher[pkg2.DotStruct]
}

type _F25Call struct {
	name itskit.Label
	spec _F25CallSpec
}

func F25_Expects(
	d its.Matcher[pkg2.DotStruct],

) _F25Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F25Call{
		name: itskit.NewLabelWithLocation("func F25"),
		spec: _F25CallSpec{
			d: itskit.Named(
				"d",
				d,
			),
		},
	}
}

type _F25Behavior struct {
	name   itskit.Label
	spec   _F25CallSpec
	effect func(d pkg2.DotStruct) pkg2.DotInterface
}

func (b *_F25Behavior) Fn(t mockkit.TestLike) func(d pkg2.DotStruct) pkg2.DotInterface {
	return func(

		arg0 pkg2.DotStruct,

	) pkg2.DotInterface {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.d
			if matcher == nil {
				matcher = its.Never[pkg2.DotStruct]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F25Call) ThenReturn(

	ret0 pkg2.DotInterface,

) mockkit.FuncBehavior[func(d pkg2.DotStruct) pkg2.DotInterface] {
	return c.ThenEffect(func(

		pkg2.DotStruct,

	) pkg2.DotInterface {

		return ret0

	})
}

func (c _F25Call) ThenEffect(effect func(d pkg2.DotStruct) pkg2.DotInterface) mockkit.FuncBehavior[func(d pkg2.DotStruct) pkg2.DotInterface] {
	return &_F25Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F26CallSpec struct {
	d its.Matcher[pkg2.DotInterface]
}

type _F26Call struct {
	name itskit.Label
	spec _F26CallSpec
}

func F26_Expects(
	d its.Matcher[pkg2.DotInterface],

) _F26Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F26Call{
		name: itskit.NewLabelWithLocation("func F26"),
		spec: _F26CallSpec{
			d: itskit.Named(
				"d",
				d,
			),
		},
	}
}

type _F26Behavior struct {
	name   itskit.Label
	spec   _F26CallSpec
	effect func(d pkg2.DotInterface) pkg2.DotStruct
}

func (b *_F26Behavior) Fn(t mockkit.TestLike) func(d pkg2.DotInterface) pkg2.DotStruct {
	return func(

		arg0 pkg2.DotInterface,

	) pkg2.DotStruct {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.d
			if matcher == nil {
				matcher = its.Never[pkg2.DotInterface]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F26Call) ThenReturn(

	ret0 pkg2.DotStruct,

) mockkit.FuncBehavior[func(d pkg2.DotInterface) pkg2.DotStruct] {
	return c.ThenEffect(func(

		pkg2.DotInterface,

	) pkg2.DotStruct {

		return ret0

	})
}

func (c _F26Call) ThenEffect(effect func(d pkg2.DotInterface) pkg2.DotStruct) mockkit.FuncBehavior[func(d pkg2.DotInterface) pkg2.DotStruct] {
	return &_F26Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F27CallSpec[T pkg2.DotInterface] struct {
	d its.Matcher[T]
}

type _F27Call[T pkg2.DotInterface] struct {
	name itskit.Label
	spec _F27CallSpec[T]
}

func F27_Expects[T pkg2.DotInterface](
	d its.Matcher[T],

) _F27Call[T] {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F27Call[T]{
		name: itskit.NewLabelWithLocation("func F27"),
		spec: _F27CallSpec[T]{
			d: itskit.Named(
				"d",
				d,
			),
		},
	}
}

type _F27Behavior[T pkg2.DotInterface] struct {
	name   itskit.Label
	spec   _F27CallSpec[T]
	effect func(d T) T
}

func (b *_F27Behavior[T]) Fn(t mockkit.TestLike) func(d T) T {
	return func(

		arg0 T,

	) T {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.d
			if matcher == nil {
				matcher = its.Never[T]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F27Call[T]) ThenReturn(

	ret0 T,

) mockkit.FuncBehavior[func(d T) T] {
	return c.ThenEffect(func(

		T,

	) T {

		return ret0

	})
}

func (c _F27Call[T]) ThenEffect(effect func(d T) T) mockkit.FuncBehavior[func(d T) T] {
	return &_F27Behavior[T]{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F28CallSpec struct {
	arg0 its.Matcher[pkg2.DotMap]
}

type _F28Call struct {
	name itskit.Label
	spec _F28CallSpec
}

func F28_Expects(
	arg0 its.Matcher[pkg2.DotMap],

) _F28Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F28Call{
		name: itskit.NewLabelWithLocation("func F28"),
		spec: _F28CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F28Behavior struct {
	name   itskit.Label
	spec   _F28CallSpec
	effect func(arg0 pkg2.DotMap) pkg2.DotSlice
}

func (b *_F28Behavior) Fn(t mockkit.TestLike) func(arg0 pkg2.DotMap) pkg2.DotSlice {
	return func(

		arg0 pkg2.DotMap,

	) pkg2.DotSlice {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[pkg2.DotMap]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F28Call) ThenReturn(

	ret0 pkg2.DotSlice,

) mockkit.FuncBehavior[func(arg0 pkg2.DotMap) pkg2.DotSlice] {
	return c.ThenEffect(func(

		pkg2.DotMap,

	) pkg2.DotSlice {

		return ret0

	})
}

func (c _F28Call) ThenEffect(effect func(arg0 pkg2.DotMap) pkg2.DotSlice) mockkit.FuncBehavior[func(arg0 pkg2.DotMap) pkg2.DotSlice] {
	return &_F28Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F29CallSpec struct {
	arg0 its.Matcher[pkg2.DotSlice]
}

type _F29Call struct {
	name itskit.Label
	spec _F29CallSpec
}

func F29_Expects(
	arg0 its.Matcher[pkg2.DotSlice],

) _F29Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F29Call{
		name: itskit.NewLabelWithLocation("func F29"),
		spec: _F29CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F29Behavior struct {
	name   itskit.Label
	spec   _F29CallSpec
	effect func(arg0 pkg2.DotSlice) pkg2.DotMap
}

func (b *_F29Behavior) Fn(t mockkit.TestLike) func(arg0 pkg2.DotSlice) pkg2.DotMap {
	return func(

		arg0 pkg2.DotSlice,

	) pkg2.DotMap {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[pkg2.DotSlice]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F29Call) ThenReturn(

	ret0 pkg2.DotMap,

) mockkit.FuncBehavior[func(arg0 pkg2.DotSlice) pkg2.DotMap] {
	return c.ThenEffect(func(

		pkg2.DotSlice,

	) pkg2.DotMap {

		return ret0

	})
}

func (c _F29Call) ThenEffect(effect func(arg0 pkg2.DotSlice) pkg2.DotMap) mockkit.FuncBehavior[func(arg0 pkg2.DotSlice) pkg2.DotMap] {
	return &_F29Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}

type _F30CallSpec struct {
	arg0 its.Matcher[pkg2.DotGene[int]]
}

type _F30Call struct {
	name itskit.Label
	spec _F30CallSpec
}

func F30_Expects(
	arg0 its.Matcher[pkg2.DotGene[int]],

) _F30Call {
	{
		cancel := itskit.SkipStack()
		defer cancel()
	}

	return _F30Call{
		name: itskit.NewLabelWithLocation("func F30"),
		spec: _F30CallSpec{
			arg0: itskit.Named(
				"arg0",
				arg0,
			),
		},
	}
}

type _F30Behavior struct {
	name   itskit.Label
	spec   _F30CallSpec
	effect func(arg0 pkg2.DotGene[int]) pkg2.DotGene[int]
}

func (b *_F30Behavior) Fn(t mockkit.TestLike) func(arg0 pkg2.DotGene[int]) pkg2.DotGene[int] {
	return func(

		arg0 pkg2.DotGene[int],

	) pkg2.DotGene[int] {
		if h, ok := t.(interface{ Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}

		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[pkg2.DotGene[int]]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}

		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(

			arg0,
		)
	}
}

func (c _F30Call) ThenReturn(

	ret0 pkg2.DotGene[int],

) mockkit.FuncBehavior[func(arg0 pkg2.DotGene[int]) pkg2.DotGene[int]] {
	return c.ThenEffect(func(

		pkg2.DotGene[int],

	) pkg2.DotGene[int] {

		return ret0

	})
}

func (c _F30Call) ThenEffect(effect func(arg0 pkg2.DotGene[int]) pkg2.DotGene[int]) mockkit.FuncBehavior[func(arg0 pkg2.DotGene[int]) pkg2.DotGene[int]] {
	return &_F30Behavior{
		name:   c.name,
		spec:   c.spec,
		effect: effect,
	}
}
