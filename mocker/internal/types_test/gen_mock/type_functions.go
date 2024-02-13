// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	testee "github.com/youta-t/its/mocker/internal/types_test"
	
)

type _F1ReturnFixture struct {
}

type _F1Return struct {
	fixture _F1ReturnFixture
}

func (rfx _F1Return) Get() (
) {
	return 
}

type _F1CallSpec struct {
	
}

type _F1Call struct {
	name itskit.Label
	spec _F1CallSpec
}

func NewF1Call(
) _F1Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F1CallSpec {}
	
	return _F1Call{
		name: itskit.NewLabelWithLocation("func F1"),
		spec: spec,
	}
}

type _F1Behaviour  struct {
	name itskit.Label
	spec _F1CallSpec
	effect func()
}

func (b _F1Behaviour) Mock(t interface { Error(...any) }) func() {
	return func (
		
		
	)  {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		b.effect(
			
			
		)
	}
}

func (c _F1Call) ThenReturn(

)_F1Behaviour {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _F1Call) ThenEffect(effect func()) _F1Behaviour {
	return _F1Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F2ReturnFixture struct {
	ret0 int
	
}

type _F2Return struct {
	fixture _F2ReturnFixture
}

func (rfx _F2Return) Get() (
	int,
	
) {
	return rfx.fixture.ret0
}

type _F2CallSpec struct {
	
}

type _F2Call struct {
	name itskit.Label
	spec _F2CallSpec
}

func NewF2Call(
) _F2Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F2CallSpec {}
	
	return _F2Call{
		name: itskit.NewLabelWithLocation("func F2"),
		spec: spec,
	}
}

type _F2Behaviour  struct {
	name itskit.Label
	spec _F2CallSpec
	effect func() int
}

func (b _F2Behaviour) Mock(t interface { Error(...any) }) func() int {
	return func (
		
		
	) (
		int,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			
		)
	}
}

func (c _F2Call) ThenReturn(

	ret0 int,

)_F2Behaviour {
	return c.ThenEffect(func(
		
		
	)(
		int,
		
	){
		
		return ret0
		
	})
}

func (c _F2Call) ThenEffect(effect func() int) _F2Behaviour {
	return _F2Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F3ReturnFixture struct {
	ret0 int
	
	ret1 string
	
}

type _F3Return struct {
	fixture _F3ReturnFixture
}

func (rfx _F3Return) Get() (
	int,
	
	string,
	
) {
	return rfx.fixture.ret0, rfx.fixture.ret1
}

type _F3CallSpec struct {
	
}

type _F3Call struct {
	name itskit.Label
	spec _F3CallSpec
}

func NewF3Call(
) _F3Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F3CallSpec {}
	
	return _F3Call{
		name: itskit.NewLabelWithLocation("func F3"),
		spec: spec,
	}
}

type _F3Behaviour  struct {
	name itskit.Label
	spec _F3CallSpec
	effect func() ( int,  string)
}

func (b _F3Behaviour) Mock(t interface { Error(...any) }) func() ( int,  string) {
	return func (
		
		
	) (
		int,
		string,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			
		)
	}
}

func (c _F3Call) ThenReturn(

	ret0 int,

	ret1 string,

)_F3Behaviour {
	return c.ThenEffect(func(
		
		
	)(
		int,
		string,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _F3Call) ThenEffect(effect func() ( int,  string)) _F3Behaviour {
	return _F3Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F4ReturnFixture struct {
	i int
	
	s string
	
}

type _F4Return struct {
	fixture _F4ReturnFixture
}

func (rfx _F4Return) Get() (
	int,
	
	string,
	
) {
	return rfx.fixture.i, rfx.fixture.s
}

type _F4CallSpec struct {
	
}

type _F4Call struct {
	name itskit.Label
	spec _F4CallSpec
}

func NewF4Call(
) _F4Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F4CallSpec {}
	
	return _F4Call{
		name: itskit.NewLabelWithLocation("func F4"),
		spec: spec,
	}
}

type _F4Behaviour  struct {
	name itskit.Label
	spec _F4CallSpec
	effect func() (i int, s string)
}

func (b _F4Behaviour) Mock(t interface { Error(...any) }) func() (i int, s string) {
	return func (
		
		
	) (
		int,
		string,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			
		)
	}
}

func (c _F4Call) ThenReturn(

	ret0 int,

	ret1 string,

)_F4Behaviour {
	return c.ThenEffect(func(
		
		
	)(
		int,
		string,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _F4Call) ThenEffect(effect func() (i int, s string)) _F4Behaviour {
	return _F4Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F5ReturnFixture struct {
}

type _F5Return struct {
	fixture _F5ReturnFixture
}

func (rfx _F5Return) Get() (
) {
	return 
}

type _F5CallSpec struct {
	arg0 its.Matcher[int]
	
	arg1 its.Matcher[string]
	
	
}

type _F5Call struct {
	name itskit.Label
	spec _F5CallSpec
}

func NewF5Call(
	arg0 its.Matcher[int],
	
	arg1 its.Matcher[string],
	
) _F5Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F5CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.arg1 = itskit.Named(
		"arg1",
		arg1,
	)
	
	
	return _F5Call{
		name: itskit.NewLabelWithLocation("func F5"),
		spec: spec,
	}
}

type _F5Behaviour  struct {
	name itskit.Label
	spec _F5CallSpec
	effect func(arg0 int, arg1 string)
}

func (b _F5Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, arg1 string) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F5Call) ThenReturn(

)_F5Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	){
		
	})
}

func (c _F5Call) ThenEffect(effect func(arg0 int, arg1 string)) _F5Behaviour {
	return _F5Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F6ReturnFixture struct {
}

type _F6Return struct {
	fixture _F6ReturnFixture
}

func (rfx _F6Return) Get() (
) {
	return 
}

type _F6CallSpec struct {
	arg0 its.Matcher[int]
	
	vararg its.Matcher[[]string]
}

type _F6Call struct {
	name itskit.Label
	spec _F6CallSpec
}

func NewF6Call(
	arg0 its.Matcher[int],
	vararg its.Matcher[[]string],
) _F6Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F6CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.vararg = itskit.Named(
		"vararg",
		vararg,
	)
	return _F6Call{
		name: itskit.NewLabelWithLocation("func F6"),
		spec: spec,
	}
}

type _F6Behaviour  struct {
	name itskit.Label
	spec _F6CallSpec
	effect func(arg0 int, vararg ...string)
}

func (b _F6Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, vararg ...string) {
	return func (
		
		arg0 int,
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F6Call) ThenReturn(

)_F6Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		...string,
	){
		
	})
}

func (c _F6Call) ThenEffect(effect func(arg0 int, vararg ...string)) _F6Behaviour {
	return _F6Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F7ReturnFixture struct {
}

type _F7Return struct {
	fixture _F7ReturnFixture
}

func (rfx _F7Return) Get() (
) {
	return 
}

type _F7CallSpec struct {
	vararg its.Matcher[[]string]
}

type _F7Call struct {
	name itskit.Label
	spec _F7CallSpec
}

func NewF7Call(vararg its.Matcher[[]string],
) _F7Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F7CallSpec {}
	spec.vararg = itskit.Named(
		"vararg",
		vararg,
	)
	return _F7Call{
		name: itskit.NewLabelWithLocation("func F7"),
		spec: spec,
	}
}

type _F7Behaviour  struct {
	name itskit.Label
	spec _F7CallSpec
	effect func(vararg ...string)
}

func (b _F7Behaviour) Mock(t interface { Error(...any) }) func(vararg ...string) {
	return func (
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F7Call) ThenReturn(

)_F7Behaviour {
	return c.ThenEffect(func(
		
		...string,
	){
		
	})
}

func (c _F7Call) ThenEffect(effect func(vararg ...string)) _F7Behaviour {
	return _F7Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F8ReturnFixture struct {
}

type _F8Return struct {
	fixture _F8ReturnFixture
}

func (rfx _F8Return) Get() (
) {
	return 
}

type _F8CallSpec struct {
	i its.Matcher[int]
	
	s its.Matcher[string]
	
	
}

type _F8Call struct {
	name itskit.Label
	spec _F8CallSpec
}

func NewF8Call(
	i its.Matcher[int],
	
	s its.Matcher[string],
	
) _F8Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F8CallSpec {}
	spec.i = itskit.Named(
		"i",
		i,
	)
	
	spec.s = itskit.Named(
		"s",
		s,
	)
	
	
	return _F8Call{
		name: itskit.NewLabelWithLocation("func F8"),
		spec: spec,
	}
}

type _F8Behaviour  struct {
	name itskit.Label
	spec _F8CallSpec
	effect func(i int, s string)
}

func (b _F8Behaviour) Mock(t interface { Error(...any) }) func(i int, s string) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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
			matcher := b.spec.s
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

func (c _F8Call) ThenReturn(

)_F8Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	){
		
	})
}

func (c _F8Call) ThenEffect(effect func(i int, s string)) _F8Behaviour {
	return _F8Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F9ReturnFixture struct {
}

type _F9Return struct {
	fixture _F9ReturnFixture
}

func (rfx _F9Return) Get() (
) {
	return 
}

type _F9CallSpec struct {
	i its.Matcher[int]
	
	ss its.Matcher[[]string]
}

type _F9Call struct {
	name itskit.Label
	spec _F9CallSpec
}

func NewF9Call(
	i its.Matcher[int],
	ss its.Matcher[[]string],
) _F9Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F9CallSpec {}
	spec.i = itskit.Named(
		"i",
		i,
	)
	
	spec.ss = itskit.Named(
		"ss",
		ss,
	)
	return _F9Call{
		name: itskit.NewLabelWithLocation("func F9"),
		spec: spec,
	}
}

type _F9Behaviour  struct {
	name itskit.Label
	spec _F9CallSpec
	effect func(i int, ss ...string)
}

func (b _F9Behaviour) Mock(t interface { Error(...any) }) func(i int, ss ...string) {
	return func (
		
		arg0 int,
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F9Call) ThenReturn(

)_F9Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		...string,
	){
		
	})
}

func (c _F9Call) ThenEffect(effect func(i int, ss ...string)) _F9Behaviour {
	return _F9Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F10ReturnFixture struct {
}

type _F10Return struct {
	fixture _F10ReturnFixture
}

func (rfx _F10Return) Get() (
) {
	return 
}

type _F10CallSpec struct {
	ss its.Matcher[[]string]
}

type _F10Call struct {
	name itskit.Label
	spec _F10CallSpec
}

func NewF10Call(ss its.Matcher[[]string],
) _F10Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F10CallSpec {}
	spec.ss = itskit.Named(
		"ss",
		ss,
	)
	return _F10Call{
		name: itskit.NewLabelWithLocation("func F10"),
		spec: spec,
	}
}

type _F10Behaviour  struct {
	name itskit.Label
	spec _F10CallSpec
	effect func(ss ...string)
}

func (b _F10Behaviour) Mock(t interface { Error(...any) }) func(ss ...string) {
	return func (
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F10Call) ThenReturn(

)_F10Behaviour {
	return c.ThenEffect(func(
		
		...string,
	){
		
	})
}

func (c _F10Call) ThenEffect(effect func(ss ...string)) _F10Behaviour {
	return _F10Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F11ReturnFixture struct {
}

type _F11Return struct {
	fixture _F11ReturnFixture
}

func (rfx _F11Return) Get() (
) {
	return 
}

type _F11CallSpec struct {
	arg0 its.Matcher[int]
	
	arg1 its.Matcher[string]
	
	
}

type _F11Call struct {
	name itskit.Label
	spec _F11CallSpec
}

func NewF11Call(
	arg0 its.Matcher[int],
	
	arg1 its.Matcher[string],
	
) _F11Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F11CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.arg1 = itskit.Named(
		"arg1",
		arg1,
	)
	
	
	return _F11Call{
		name: itskit.NewLabelWithLocation("func F11"),
		spec: spec,
	}
}

type _F11Behaviour  struct {
	name itskit.Label
	spec _F11CallSpec
	effect func(arg0 int, arg1 string)
}

func (b _F11Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, arg1 string) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F11Call) ThenReturn(

)_F11Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	){
		
	})
}

func (c _F11Call) ThenEffect(effect func(arg0 int, arg1 string)) _F11Behaviour {
	return _F11Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F12ReturnFixture struct {
}

type _F12Return struct {
	fixture _F12ReturnFixture
}

func (rfx _F12Return) Get() (
) {
	return 
}

type _F12CallSpec struct {
	arg0 its.Matcher[int]
	
	vararg its.Matcher[[]string]
}

type _F12Call struct {
	name itskit.Label
	spec _F12CallSpec
}

func NewF12Call(
	arg0 its.Matcher[int],
	vararg its.Matcher[[]string],
) _F12Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F12CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.vararg = itskit.Named(
		"vararg",
		vararg,
	)
	return _F12Call{
		name: itskit.NewLabelWithLocation("func F12"),
		spec: spec,
	}
}

type _F12Behaviour  struct {
	name itskit.Label
	spec _F12CallSpec
	effect func(arg0 int, vararg ...string)
}

func (b _F12Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, vararg ...string) {
	return func (
		
		arg0 int,
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F12Call) ThenReturn(

)_F12Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		...string,
	){
		
	})
}

func (c _F12Call) ThenEffect(effect func(arg0 int, vararg ...string)) _F12Behaviour {
	return _F12Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F13ReturnFixture struct {
}

type _F13Return struct {
	fixture _F13ReturnFixture
}

func (rfx _F13Return) Get() (
) {
	return 
}

type _F13CallSpec struct {
	vararg its.Matcher[[]string]
}

type _F13Call struct {
	name itskit.Label
	spec _F13CallSpec
}

func NewF13Call(vararg its.Matcher[[]string],
) _F13Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F13CallSpec {}
	spec.vararg = itskit.Named(
		"vararg",
		vararg,
	)
	return _F13Call{
		name: itskit.NewLabelWithLocation("func F13"),
		spec: spec,
	}
}

type _F13Behaviour  struct {
	name itskit.Label
	spec _F13CallSpec
	effect func(vararg ...string)
}

func (b _F13Behaviour) Mock(t interface { Error(...any) }) func(vararg ...string) {
	return func (
		
		vararg ...string,
	)  {
		if h, ok := t.(interface { Helper() }); ok {
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

func (c _F13Call) ThenReturn(

)_F13Behaviour {
	return c.ThenEffect(func(
		
		...string,
	){
		
	})
}

func (c _F13Call) ThenEffect(effect func(vararg ...string)) _F13Behaviour {
	return _F13Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F14ReturnFixture struct {
	f float64
	
}

type _F14Return struct {
	fixture _F14ReturnFixture
}

func (rfx _F14Return) Get() (
	float64,
	
) {
	return rfx.fixture.f
}

type _F14CallSpec struct {
	arg0 its.Matcher[int]
	
	arg1 its.Matcher[string]
	
	
}

type _F14Call struct {
	name itskit.Label
	spec _F14CallSpec
}

func NewF14Call(
	arg0 its.Matcher[int],
	
	arg1 its.Matcher[string],
	
) _F14Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F14CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.arg1 = itskit.Named(
		"arg1",
		arg1,
	)
	
	
	return _F14Call{
		name: itskit.NewLabelWithLocation("func F14"),
		spec: spec,
	}
}

type _F14Behaviour  struct {
	name itskit.Label
	spec _F14CallSpec
	effect func(arg0 int, arg1 string) (f float64)
}

func (b _F14Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, arg1 string) (f float64) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	) (
		float64,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
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

)_F14Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	)(
		float64,
		
	){
		
		return ret0
		
	})
}

func (c _F14Call) ThenEffect(effect func(arg0 int, arg1 string) (f float64)) _F14Behaviour {
	return _F14Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F15ReturnFixture struct {
	ret0 float64
	
}

type _F15Return struct {
	fixture _F15ReturnFixture
}

func (rfx _F15Return) Get() (
	float64,
	
) {
	return rfx.fixture.ret0
}

type _F15CallSpec struct {
	i its.Matcher[int]
	
	ss its.Matcher[[]string]
}

type _F15Call struct {
	name itskit.Label
	spec _F15CallSpec
}

func NewF15Call(
	i its.Matcher[int],
	ss its.Matcher[[]string],
) _F15Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F15CallSpec {}
	spec.i = itskit.Named(
		"i",
		i,
	)
	
	spec.ss = itskit.Named(
		"ss",
		ss,
	)
	return _F15Call{
		name: itskit.NewLabelWithLocation("func F15"),
		spec: spec,
	}
}

type _F15Behaviour  struct {
	name itskit.Label
	spec _F15CallSpec
	effect func(i int, ss ...string) float64
}

func (b _F15Behaviour) Mock(t interface { Error(...any) }) func(i int, ss ...string) float64 {
	return func (
		
		arg0 int,
		
		vararg ...string,
	) (
		float64,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
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

)_F15Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		...string,
	)(
		float64,
		
	){
		
		return ret0
		
	})
}

func (c _F15Call) ThenEffect(effect func(i int, ss ...string) float64) _F15Behaviour {
	return _F15Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F16ReturnFixture[T any] struct {
	ret0 T
	
}

type _F16Return[T any] struct {
	fixture _F16ReturnFixture[T]
}

func (rfx _F16Return[T]) Get() (
	T,
	
) {
	return rfx.fixture.ret0
}

type _F16CallSpec[T any] struct {
	arg0 its.Matcher[testee.F16[T]]
	
	
}

type _F16Call[T any] struct {
	name itskit.Label
	spec _F16CallSpec[T]
}

func NewF16Call[T any](
	arg0 its.Matcher[testee.F16[T]],
	
) _F16Call[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F16CallSpec[T] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _F16Call[T]{
		name: itskit.NewLabelWithLocation("func F16"),
		spec: spec,
	}
}

type _F16Behaviour [T any] struct {
	name itskit.Label
	spec _F16CallSpec[T]
	effect func(arg0 testee.F16[T]) T
}

func (b _F16Behaviour[T]) Mock(t interface { Error(...any) }) func(arg0 testee.F16[T]) T {
	return func (
		
		arg0 testee.F16[T],
		
		
	) (
		T,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[testee.F16[T]]()
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

)_F16Behaviour[T] {
	return c.ThenEffect(func(
		
		testee.F16[T],
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _F16Call[T]) ThenEffect(effect func(arg0 testee.F16[T]) T) _F16Behaviour[T] {
	return _F16Behaviour[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F17ReturnFixture[T struct{
	Foo int
}] struct {
	ret0 T
	
}

type _F17Return[T struct{
	Foo int
}] struct {
	fixture _F17ReturnFixture[T]
}

func (rfx _F17Return[T]) Get() (
	T,
	
) {
	return rfx.fixture.ret0
}

type _F17CallSpec[T struct{
	Foo int
}] struct {
	arg0 its.Matcher[T]
	
	
}

type _F17Call[T struct{
	Foo int
}] struct {
	name itskit.Label
	spec _F17CallSpec[T]
}

func NewF17Call[T struct{
	Foo int
}](
	arg0 its.Matcher[T],
	
) _F17Call[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F17CallSpec[T] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _F17Call[T]{
		name: itskit.NewLabelWithLocation("func F17"),
		spec: spec,
	}
}

type _F17Behaviour [T struct{
	Foo int
}] struct {
	name itskit.Label
	spec _F17CallSpec[T]
	effect func(arg0 T) T
}

func (b _F17Behaviour[T]) Mock(t interface { Error(...any) }) func(arg0 T) T {
	return func (
		
		arg0 T,
		
		
	) (
		T,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
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

)_F17Behaviour[T] {
	return c.ThenEffect(func(
		
		T,
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _F17Call[T]) ThenEffect(effect func(arg0 T) T) _F17Behaviour[T] {
	return _F17Behaviour[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F18ReturnFixture[T any, U any] struct {
	ret0 func(arg0 T, arg1 U)
	
}

type _F18Return[T any, U any] struct {
	fixture _F18ReturnFixture[T, U]
}

func (rfx _F18Return[T, U]) Get() (
	func(arg0 T, arg1 U),
	
) {
	return rfx.fixture.ret0
}

type _F18CallSpec[T any, U any] struct {
	arg0 its.Matcher[func(arg0 T, arg1 U)]
	
	
}

type _F18Call[T any, U any] struct {
	name itskit.Label
	spec _F18CallSpec[T, U]
}

func NewF18Call[T any, U any](
	arg0 its.Matcher[func(arg0 T, arg1 U)],
	
) _F18Call[T, U] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F18CallSpec[T, U] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _F18Call[T, U]{
		name: itskit.NewLabelWithLocation("func F18"),
		spec: spec,
	}
}

type _F18Behaviour [T any, U any] struct {
	name itskit.Label
	spec _F18CallSpec[T, U]
	effect func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)
}

func (b _F18Behaviour[T, U]) Mock(t interface { Error(...any) }) func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U) {
	return func (
		
		arg0 func(arg0 T, arg1 U),
		
		
	) (
		func(arg0 T, arg1 U),
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
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

)_F18Behaviour[T, U] {
	return c.ThenEffect(func(
		
		func(arg0 T, arg1 U),
		
		
	)(
		func(arg0 T, arg1 U),
		
	){
		
		return ret0
		
	})
}

func (c _F18Call[T, U]) ThenEffect(effect func(arg0 func(arg0 T, arg1 U)) func(arg0 T, arg1 U)) _F18Behaviour[T, U] {
	return _F18Behaviour[T, U] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F19ReturnFixture[T ~int] struct {
	ret0 T
	
}

type _F19Return[T ~int] struct {
	fixture _F19ReturnFixture[T]
}

func (rfx _F19Return[T]) Get() (
	T,
	
) {
	return rfx.fixture.ret0
}

type _F19CallSpec[T ~int] struct {
	arg0 its.Matcher[testee.F16[T]]
	
	
}

type _F19Call[T ~int] struct {
	name itskit.Label
	spec _F19CallSpec[T]
}

func NewF19Call[T ~int](
	arg0 its.Matcher[testee.F16[T]],
	
) _F19Call[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F19CallSpec[T] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _F19Call[T]{
		name: itskit.NewLabelWithLocation("func F19"),
		spec: spec,
	}
}

type _F19Behaviour [T ~int] struct {
	name itskit.Label
	spec _F19CallSpec[T]
	effect func(arg0 testee.F16[T]) T
}

func (b _F19Behaviour[T]) Mock(t interface { Error(...any) }) func(arg0 testee.F16[T]) T {
	return func (
		
		arg0 testee.F16[T],
		
		
	) (
		T,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[testee.F16[T]]()
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

)_F19Behaviour[T] {
	return c.ThenEffect(func(
		
		testee.F16[T],
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _F19Call[T]) ThenEffect(effect func(arg0 testee.F16[T]) T) _F19Behaviour[T] {
	return _F19Behaviour[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _F20ReturnFixture[T ~int | int8 | ~int16 | int32] struct {
	ret0 T
	
}

type _F20Return[T ~int | int8 | ~int16 | int32] struct {
	fixture _F20ReturnFixture[T]
}

func (rfx _F20Return[T]) Get() (
	T,
	
) {
	return rfx.fixture.ret0
}

type _F20CallSpec[T ~int | int8 | ~int16 | int32] struct {
	arg0 its.Matcher[T]
	
	
}

type _F20Call[T ~int | int8 | ~int16 | int32] struct {
	name itskit.Label
	spec _F20CallSpec[T]
}

func NewF20Call[T ~int | int8 | ~int16 | int32](
	arg0 its.Matcher[T],
	
) _F20Call[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _F20CallSpec[T] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _F20Call[T]{
		name: itskit.NewLabelWithLocation("func F20"),
		spec: spec,
	}
}

type _F20Behaviour [T ~int | int8 | ~int16 | int32] struct {
	name itskit.Label
	spec _F20CallSpec[T]
	effect func(arg0 T) T
}

func (b _F20Behaviour[T]) Mock(t interface { Error(...any) }) func(arg0 T) T {
	return func (
		
		arg0 T,
		
		
	) (
		T,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
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

)_F20Behaviour[T] {
	return c.ThenEffect(func(
		
		T,
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _F20Call[T]) ThenEffect(effect func(arg0 T) T) _F20Behaviour[T] {
	return _F20Behaviour[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}




