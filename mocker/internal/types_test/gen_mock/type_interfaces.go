// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	testee "github.com/youta-t/its/mocker/internal/types_test"
	u_sub "github.com/youta-t/its/mocker/internal/example/sub"
	
)

type _I0_M0ReturnFixture struct {
}

type _I0_M0Return struct {
	fixture _I0_M0ReturnFixture
}

func (rfx _I0_M0Return) Get() (
) {
	return 
}

type _I0_M0CallSpec struct {
	
}

type _I0_M0Call struct {
	name itskit.Label
	spec _I0_M0CallSpec
}

func NewI0_M0Call(
) _I0_M0Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M0CallSpec {}
	
	return _I0_M0Call{
		name: itskit.NewLabelWithLocation("func I0_M0"),
		spec: spec,
	}
}

type _I0_M0Behaviour  struct {
	name itskit.Label
	spec _I0_M0CallSpec
	effect func()
}

func (b _I0_M0Behaviour) Mock(t interface { Error(...any) }) func() {
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

func (c _I0_M0Call) ThenReturn(

)_I0_M0Behaviour {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _I0_M0Call) ThenEffect(effect func()) _I0_M0Behaviour {
	return _I0_M0Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I0_M1ReturnFixture struct {
	ret0 bool
	
	ret1 error
	
}

type _I0_M1Return struct {
	fixture _I0_M1ReturnFixture
}

func (rfx _I0_M1Return) Get() (
	bool,
	
	error,
	
) {
	return rfx.fixture.ret0, rfx.fixture.ret1
}

type _I0_M1CallSpec struct {
	arg0 its.Matcher[int]
	
	arg1 its.Matcher[string]
	
	
}

type _I0_M1Call struct {
	name itskit.Label
	spec _I0_M1CallSpec
}

func NewI0_M1Call(
	arg0 its.Matcher[int],
	
	arg1 its.Matcher[string],
	
) _I0_M1Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M1CallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.arg1 = itskit.Named(
		"arg1",
		arg1,
	)
	
	
	return _I0_M1Call{
		name: itskit.NewLabelWithLocation("func I0_M1"),
		spec: spec,
	}
}

type _I0_M1Behaviour  struct {
	name itskit.Label
	spec _I0_M1CallSpec
	effect func(arg0 int, arg1 string) ( bool,  error)
}

func (b _I0_M1Behaviour) Mock(t interface { Error(...any) }) func(arg0 int, arg1 string) ( bool,  error) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	) (
		bool,
		error,
		
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

func (c _I0_M1Call) ThenReturn(

	ret0 bool,

	ret1 error,

)_I0_M1Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	)(
		bool,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _I0_M1Call) ThenEffect(effect func(arg0 int, arg1 string) ( bool,  error)) _I0_M1Behaviour {
	return _I0_M1Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I0_M2ReturnFixture struct {
	ok bool
	
	err error
	
}

type _I0_M2Return struct {
	fixture _I0_M2ReturnFixture
}

func (rfx _I0_M2Return) Get() (
	bool,
	
	error,
	
) {
	return rfx.fixture.ok, rfx.fixture.err
}

type _I0_M2CallSpec struct {
	i its.Matcher[int]
	
	s its.Matcher[string]
	
	
}

type _I0_M2Call struct {
	name itskit.Label
	spec _I0_M2CallSpec
}

func NewI0_M2Call(
	i its.Matcher[int],
	
	s its.Matcher[string],
	
) _I0_M2Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M2CallSpec {}
	spec.i = itskit.Named(
		"i",
		i,
	)
	
	spec.s = itskit.Named(
		"s",
		s,
	)
	
	
	return _I0_M2Call{
		name: itskit.NewLabelWithLocation("func I0_M2"),
		spec: spec,
	}
}

type _I0_M2Behaviour  struct {
	name itskit.Label
	spec _I0_M2CallSpec
	effect func(i int, s string) (ok bool, err error)
}

func (b _I0_M2Behaviour) Mock(t interface { Error(...any) }) func(i int, s string) (ok bool, err error) {
	return func (
		
		arg0 int,
		
		arg1 string,
		
		
	) (
		bool,
		error,
		
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
		return b.effect(
			
			arg0,
			
			arg1,
			
			
		)
	}
}

func (c _I0_M2Call) ThenReturn(

	ret0 bool,

	ret1 error,

)_I0_M2Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		string,
		
		
	)(
		bool,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _I0_M2Call) ThenEffect(effect func(i int, s string) (ok bool, err error)) _I0_M2Behaviour {
	return _I0_M2Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I0_M3ReturnFixture struct {
	ok bool
	
	err error
	
}

type _I0_M3Return struct {
	fixture _I0_M3ReturnFixture
}

func (rfx _I0_M3Return) Get() (
	bool,
	
	error,
	
) {
	return rfx.fixture.ok, rfx.fixture.err
}

type _I0_M3CallSpec struct {
	i its.Matcher[int]
	
	s its.Matcher[[]string]
}

type _I0_M3Call struct {
	name itskit.Label
	spec _I0_M3CallSpec
}

func NewI0_M3Call(
	i its.Matcher[int],
	s its.Matcher[[]string],
) _I0_M3Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M3CallSpec {}
	spec.i = itskit.Named(
		"i",
		i,
	)
	
	spec.s = itskit.Named(
		"s",
		s,
	)
	return _I0_M3Call{
		name: itskit.NewLabelWithLocation("func I0_M3"),
		spec: spec,
	}
}

type _I0_M3Behaviour  struct {
	name itskit.Label
	spec _I0_M3CallSpec
	effect func(i int, s ...string) (ok bool, err error)
}

func (b _I0_M3Behaviour) Mock(t interface { Error(...any) }) func(i int, s ...string) (ok bool, err error) {
	return func (
		
		arg0 int,
		
		vararg ...string,
	) (
		bool,
		error,
		
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
			matcher := b.spec.s
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

func (c _I0_M3Call) ThenReturn(

	ret0 bool,

	ret1 error,

)_I0_M3Behaviour {
	return c.ThenEffect(func(
		
		int,
		
		...string,
	)(
		bool,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _I0_M3Call) ThenEffect(effect func(i int, s ...string) (ok bool, err error)) _I0_M3Behaviour {
	return _I0_M3Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I0_M4ReturnFixture struct {
	ret0 bool
	
}

type _I0_M4Return struct {
	fixture _I0_M4ReturnFixture
}

func (rfx _I0_M4Return) Get() (
	bool,
	
) {
	return rfx.fixture.ret0
}

type _I0_M4CallSpec struct {
	s its.Matcher[[]string]
}

type _I0_M4Call struct {
	name itskit.Label
	spec _I0_M4CallSpec
}

func NewI0_M4Call(s its.Matcher[[]string],
) _I0_M4Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M4CallSpec {}
	spec.s = itskit.Named(
		"s",
		s,
	)
	return _I0_M4Call{
		name: itskit.NewLabelWithLocation("func I0_M4"),
		spec: spec,
	}
}

type _I0_M4Behaviour  struct {
	name itskit.Label
	spec _I0_M4CallSpec
	effect func(s ...string) bool
}

func (b _I0_M4Behaviour) Mock(t interface { Error(...any) }) func(s ...string) bool {
	return func (
		
		vararg ...string,
	) (
		bool,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.s
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
			
			
			vararg...,
			
		)
	}
}

func (c _I0_M4Call) ThenReturn(

	ret0 bool,

)_I0_M4Behaviour {
	return c.ThenEffect(func(
		
		...string,
	)(
		bool,
		
	){
		
		return ret0
		
	})
}

func (c _I0_M4Call) ThenEffect(effect func(s ...string) bool) _I0_M4Behaviour {
	return _I0_M4Behaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I1_M0ReturnFixture[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	ret0 U
	
	ret1 error
	
}

type _I1_M0Return[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	fixture _I1_M0ReturnFixture[S, T, U]
}

func (rfx _I1_M0Return[S, T, U]) Get() (
	U,
	
	error,
	
) {
	return rfx.fixture.ret0, rfx.fixture.ret1
}

type _I1_M0CallSpec[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	arg0 its.Matcher[S]
	
	arg1 its.Matcher[T]
	
	
}

type _I1_M0Call[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M0CallSpec[S, T, U]
}

func NewI1_M0Call[S u_sub.T, T u_sub.T, U testee.X[T]](
	arg0 its.Matcher[S],
	
	arg1 its.Matcher[T],
	
) _I1_M0Call[S, T, U] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I1_M0CallSpec[S, T, U] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	spec.arg1 = itskit.Named(
		"arg1",
		arg1,
	)
	
	
	return _I1_M0Call[S, T, U]{
		name: itskit.NewLabelWithLocation("func I1_M0"),
		spec: spec,
	}
}

type _I1_M0Behaviour [S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M0CallSpec[S, T, U]
	effect func(arg0 S, arg1 T) ( U,  error)
}

func (b _I1_M0Behaviour[S, T, U]) Mock(t interface { Error(...any) }) func(arg0 S, arg1 T) ( U,  error) {
	return func (
		
		arg0 S,
		
		arg1 T,
		
		
	) (
		U,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[S]()
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
				matcher = its.Never[T]()
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

func (c _I1_M0Call[S, T, U]) ThenReturn(

	ret0 U,

	ret1 error,

)_I1_M0Behaviour[S, T, U] {
	return c.ThenEffect(func(
		
		S,
		
		T,
		
		
	)(
		U,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _I1_M0Call[S, T, U]) ThenEffect(effect func(arg0 S, arg1 T) ( U,  error)) _I1_M0Behaviour[S, T, U] {
	return _I1_M0Behaviour[S, T, U] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I1_M1ReturnFixture[S u_sub.T] struct {
}

type _I1_M1Return[S u_sub.T] struct {
	fixture _I1_M1ReturnFixture[S]
}

func (rfx _I1_M1Return[S]) Get() (
) {
	return 
}

type _I1_M1CallSpec[S u_sub.T] struct {
	arg0 its.Matcher[S]
	
	
}

type _I1_M1Call[S u_sub.T] struct {
	name itskit.Label
	spec _I1_M1CallSpec[S]
}

func NewI1_M1Call[S u_sub.T](
	arg0 its.Matcher[S],
	
) _I1_M1Call[S] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I1_M1CallSpec[S] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _I1_M1Call[S]{
		name: itskit.NewLabelWithLocation("func I1_M1"),
		spec: spec,
	}
}

type _I1_M1Behaviour [S u_sub.T] struct {
	name itskit.Label
	spec _I1_M1CallSpec[S]
	effect func(arg0 S)
}

func (b _I1_M1Behaviour[S]) Mock(t interface { Error(...any) }) func(arg0 S) {
	return func (
		
		arg0 S,
		
		
	)  {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[S]()
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
		b.effect(
			
			arg0,
			
			
		)
	}
}

func (c _I1_M1Call[S]) ThenReturn(

)_I1_M1Behaviour[S] {
	return c.ThenEffect(func(
		
		S,
		
		
	){
		
	})
}

func (c _I1_M1Call[S]) ThenEffect(effect func(arg0 S)) _I1_M1Behaviour[S] {
	return _I1_M1Behaviour[S] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I1_M2ReturnFixture[T u_sub.T, U testee.X[T]] struct {
	ret0 U
	
}

type _I1_M2Return[T u_sub.T, U testee.X[T]] struct {
	fixture _I1_M2ReturnFixture[T, U]
}

func (rfx _I1_M2Return[T, U]) Get() (
	U,
	
) {
	return rfx.fixture.ret0
}

type _I1_M2CallSpec[T u_sub.T, U testee.X[T]] struct {
	
}

type _I1_M2Call[T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M2CallSpec[T, U]
}

func NewI1_M2Call[T u_sub.T, U testee.X[T]](
) _I1_M2Call[T, U] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I1_M2CallSpec[T, U] {}
	
	return _I1_M2Call[T, U]{
		name: itskit.NewLabelWithLocation("func I1_M2"),
		spec: spec,
	}
}

type _I1_M2Behaviour [T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M2CallSpec[T, U]
	effect func() U
}

func (b _I1_M2Behaviour[T, U]) Mock(t interface { Error(...any) }) func() U {
	return func (
		
		
	) (
		U,
		
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

func (c _I1_M2Call[T, U]) ThenReturn(

	ret0 U,

)_I1_M2Behaviour[T, U] {
	return c.ThenEffect(func(
		
		
	)(
		U,
		
	){
		
		return ret0
		
	})
}

func (c _I1_M2Call[T, U]) ThenEffect(effect func() U) _I1_M2Behaviour[T, U] {
	return _I1_M2Behaviour[T, U] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _I2_M0ReturnFixture[T ~string] struct {
	ret0 T
	
}

type _I2_M0Return[T ~string] struct {
	fixture _I2_M0ReturnFixture[T]
}

func (rfx _I2_M0Return[T]) Get() (
	T,
	
) {
	return rfx.fixture.ret0
}

type _I2_M0CallSpec[T ~string] struct {
	arg0 its.Matcher[T]
	
	
}

type _I2_M0Call[T ~string] struct {
	name itskit.Label
	spec _I2_M0CallSpec[T]
}

func NewI2_M0Call[T ~string](
	arg0 its.Matcher[T],
	
) _I2_M0Call[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I2_M0CallSpec[T] {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _I2_M0Call[T]{
		name: itskit.NewLabelWithLocation("func I2_M0"),
		spec: spec,
	}
}

type _I2_M0Behaviour [T ~string] struct {
	name itskit.Label
	spec _I2_M0CallSpec[T]
	effect func(arg0 T) T
}

func (b _I2_M0Behaviour[T]) Mock(t interface { Error(...any) }) func(arg0 T) T {
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

func (c _I2_M0Call[T]) ThenReturn(

	ret0 T,

)_I2_M0Behaviour[T] {
	return c.ThenEffect(func(
		
		T,
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _I2_M0Call[T]) ThenEffect(effect func(arg0 T) T) _I2_M0Behaviour[T] {
	return _I2_M0Behaviour[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}




type I0Impl struct {
	
	M0 func()
	M1 func(arg0 int, arg1 string) ( bool,  error)
	M2 func(i int, s string) (ok bool, err error)
	M3 func(i int, s ...string) (ok bool, err error)
	M4 func(s ...string) bool
}

func NewMockedI0(t interface { Fatal(...any) } ,impl I0Impl) testee.I0 {
	return _I0Mock { t: t, impl: impl }
}

type _I0Mock struct {
	t interface { Fatal(...any) }
	impl I0Impl
}

func (m _I0Mock) M0 (
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M0 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I0.M0 is not mocked").String(),
		).OrFatal(m.t)
	}

	m.impl.M0(
	)
}


func (m _I0Mock) M1 (
	arg0 int,
	arg1 string,
) (
	bool,
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M1 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I0.M1 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M1(
		arg0,
		arg1,
	)
}


func (m _I0Mock) M2 (
	i int,
	s string,
) (
	ok bool,
	err error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M2 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I0.M2 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M2(
		i,
		s,
	)
}


func (m _I0Mock) M3 (
	i int,
	s ...string,
) (
	ok bool,
	err error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M3 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I0.M3 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M3(
		i,
		s...,
	)
}


func (m _I0Mock) M4 (
	s ...string,
) (
	bool,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M4 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I0.M4 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M4(
		s...,
	)
}


type I1Impl[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	
	M0 func(arg0 S, arg1 T) ( U,  error)
	M1 func(arg0 S)
	M2 func() U
}

func NewMockedI1[S u_sub.T, T u_sub.T, U testee.X[T]](t interface { Fatal(...any) } ,impl I1Impl[S, T, U]) testee.I1[S, T, U] {
	return _I1Mock[S, T, U] { t: t, impl: impl }
}

type _I1Mock[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	t interface { Fatal(...any) }
	impl I1Impl[S, T, U]
}

func (m _I1Mock[S, T, U]) M0 (
	arg0 S,
	arg1 T,
) (
	U,
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M0 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I1.M0 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M0(
		arg0,
		arg1,
	)
}


func (m _I1Mock[S, T, U]) M1 (
	arg0 S,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M1 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I1.M1 is not mocked").String(),
		).OrFatal(m.t)
	}

	m.impl.M1(
		arg0,
	)
}


func (m _I1Mock[S, T, U]) M2 (
) (
	U,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M2 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I1.M2 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M2(
	)
}


type I2Impl[T ~string] struct {
	
	M0 func(arg0 T) T
}

func NewMockedI2[T ~string](t interface { Fatal(...any) } ,impl I2Impl[T]) testee.I2[T] {
	return _I2Mock[T] { t: t, impl: impl }
}

type _I2Mock[T ~string] struct {
	t interface { Fatal(...any) }
	impl I2Impl[T]
}

func (m _I2Mock[T]) M0 (
	arg0 T,
) (
	T,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.M0 == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("I2.M0 is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.M0(
		arg0,
	)
}


