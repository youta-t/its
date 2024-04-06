// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	mockkit "github.com/youta-t/its/mocker/mockkit"
	testee "github.com/youta-t/its/mocker/internal/types"
	u_sub "github.com/youta-t/its/mocker/internal/example/sub"
	
)
type _I0_M0CallSpec struct {
	
}

type _I0_M0Call struct {
	name itskit.Label
	spec _I0_M0CallSpec
}

func I0_M0_Expects(
) _I0_M0Call {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I0_M0CallSpec {}
	
	return _I0_M0Call{
		name: itskit.NewLabelWithLocation("func I0_M0"),
		spec: spec,
	}
}

type _I0_M0Behavior  struct {
	name itskit.Label
	spec _I0_M0CallSpec
	effect func()
}

func (b *_I0_M0Behavior) Fn(t mockkit.TestLike) func() {
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

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _I0_M0Call) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_I0_M0Behavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I0_M1CallSpec struct {
	arg0 its.Matcher[int]
	
	arg1 its.Matcher[string]
	
	
}

type _I0_M1Call struct {
	name itskit.Label
	spec _I0_M1CallSpec
}

func I0_M1_Expects(
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

type _I0_M1Behavior  struct {
	name itskit.Label
	spec _I0_M1CallSpec
	effect func(arg0 int, arg1 string) ( bool,  error)
}

func (b *_I0_M1Behavior) Fn(t mockkit.TestLike) func(arg0 int, arg1 string) ( bool,  error) {
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

) mockkit.FuncBehavior[ func (arg0 int, arg1 string) ( bool,  error)  ] {
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

func (c _I0_M1Call) ThenEffect(effect func(arg0 int, arg1 string) ( bool,  error)) mockkit.FuncBehavior[ func (arg0 int, arg1 string) ( bool,  error) ] {
	return &_I0_M1Behavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I0_M2CallSpec struct {
	i its.Matcher[int]
	
	s its.Matcher[string]
	
	
}

type _I0_M2Call struct {
	name itskit.Label
	spec _I0_M2CallSpec
}

func I0_M2_Expects(
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

type _I0_M2Behavior  struct {
	name itskit.Label
	spec _I0_M2CallSpec
	effect func(i int, s string) (ok bool, err error)
}

func (b *_I0_M2Behavior) Fn(t mockkit.TestLike) func(i int, s string) (ok bool, err error) {
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

) mockkit.FuncBehavior[ func (i int, s string) (ok bool, err error)  ] {
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

func (c _I0_M2Call) ThenEffect(effect func(i int, s string) (ok bool, err error)) mockkit.FuncBehavior[ func (i int, s string) (ok bool, err error) ] {
	return &_I0_M2Behavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I0_M3CallSpec struct {
	i its.Matcher[int]
	
	s its.Matcher[[]string]
}

type _I0_M3Call struct {
	name itskit.Label
	spec _I0_M3CallSpec
}

func I0_M3_Expects(
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

type _I0_M3Behavior  struct {
	name itskit.Label
	spec _I0_M3CallSpec
	effect func(i int, s ...string) (ok bool, err error)
}

func (b *_I0_M3Behavior) Fn(t mockkit.TestLike) func(i int, s ...string) (ok bool, err error) {
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

) mockkit.FuncBehavior[ func (i int, s ...string) (ok bool, err error)  ] {
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

func (c _I0_M3Call) ThenEffect(effect func(i int, s ...string) (ok bool, err error)) mockkit.FuncBehavior[ func (i int, s ...string) (ok bool, err error) ] {
	return &_I0_M3Behavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


<<<<<<< HEAD:mocker/internal/types/gen_mock/type_interfaces.go
type _I0_M4CallSpec struct {
	s its.Matcher[[]string]
}

type _I0_M4Call struct {
	name itskit.Label
	spec _I0_M4CallSpec
}

func I0_M4_Expects(s its.Matcher[[]string],
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

type _I0_M4Behavior  struct {
	name itskit.Label
	spec _I0_M4CallSpec
	effect func(s ...string) bool
}

func (b *_I0_M4Behavior) Fn(t mockkit.TestLike) func(s ...string) bool {
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

) mockkit.FuncBehavior[ func (s ...string) bool  ] {
	return c.ThenEffect(func(
		
		...string,
	)(
		bool,
		
	){
		
		return ret0
		
	})
}

func (c _I0_M4Call) ThenEffect(effect func(s ...string) bool) mockkit.FuncBehavior[ func (s ...string) bool ] {
	return &_I0_M4Behavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I1_M0CallSpec[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
=======
type _I1_M0CallSpec[S pkg1.T, T pkg1.T, U pkg2.X[T]] struct {
>>>>>>> ae47cc0 (fixup! 862407dbe81a7f34e1dc36302d89ed60d4a7cdc2):mocker/internal/generate_test/gen_mock/type_interfaces.go
	arg0 its.Matcher[S]
	
	arg1 its.Matcher[T]
	
	
}

type _I1_M0Call[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M0CallSpec[S, T, U]
}

func I1_M0_Expects[S u_sub.T, T u_sub.T, U testee.X[T]](
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

type _I1_M0Behavior [S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M0CallSpec[S, T, U]
	effect func(arg0 S, arg1 T) ( U,  error)
}

func (b *_I1_M0Behavior[S, T, U]) Fn(t mockkit.TestLike) func(arg0 S, arg1 T) ( U,  error) {
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

) mockkit.FuncBehavior[ func (arg0 S, arg1 T) ( U,  error)  ] {
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

func (c _I1_M0Call[S, T, U]) ThenEffect(effect func(arg0 S, arg1 T) ( U,  error)) mockkit.FuncBehavior[ func (arg0 S, arg1 T) ( U,  error) ] {
	return &_I1_M0Behavior[S, T, U] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I1_M1CallSpec[S u_sub.T] struct {
	arg0 its.Matcher[S]
	
	
}

type _I1_M1Call[S u_sub.T] struct {
	name itskit.Label
	spec _I1_M1CallSpec[S]
}

func I1_M1_Expects[S u_sub.T](
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

type _I1_M1Behavior [S u_sub.T] struct {
	name itskit.Label
	spec _I1_M1CallSpec[S]
	effect func(arg0 S)
}

func (b *_I1_M1Behavior[S]) Fn(t mockkit.TestLike) func(arg0 S) {
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

) mockkit.FuncBehavior[ func (arg0 S)  ] {
	return c.ThenEffect(func(
		
		S,
		
		
	){
		
	})
}

func (c _I1_M1Call[S]) ThenEffect(effect func(arg0 S)) mockkit.FuncBehavior[ func (arg0 S) ] {
	return &_I1_M1Behavior[S] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I1_M2CallSpec[T u_sub.T, U testee.X[T]] struct {
	
}

type _I1_M2Call[T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M2CallSpec[T, U]
}

func I1_M2_Expects[T u_sub.T, U testee.X[T]](
) _I1_M2Call[T, U] {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _I1_M2CallSpec[T, U] {}
	
	return _I1_M2Call[T, U]{
		name: itskit.NewLabelWithLocation("func I1_M2"),
		spec: spec,
	}
}

type _I1_M2Behavior [T u_sub.T, U testee.X[T]] struct {
	name itskit.Label
	spec _I1_M2CallSpec[T, U]
	effect func() U
}

func (b *_I1_M2Behavior[T, U]) Fn(t mockkit.TestLike) func() U {
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

) mockkit.FuncBehavior[ func () U  ] {
	return c.ThenEffect(func(
		
		
	)(
		U,
		
	){
		
		return ret0
		
	})
}

func (c _I1_M2Call[T, U]) ThenEffect(effect func() U) mockkit.FuncBehavior[ func () U ] {
	return &_I1_M2Behavior[T, U] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _I2_M0CallSpec[T ~string] struct {
	arg0 its.Matcher[T]
	
	
}

type _I2_M0Call[T ~string] struct {
	name itskit.Label
	spec _I2_M0CallSpec[T]
}

func I2_M0_Expects[T ~string](
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

type _I2_M0Behavior [T ~string] struct {
	name itskit.Label
	spec _I2_M0CallSpec[T]
	effect func(arg0 T) T
}

func (b *_I2_M0Behavior[T]) Fn(t mockkit.TestLike) func(arg0 T) T {
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

) mockkit.FuncBehavior[ func (arg0 T) T  ] {
	return c.ThenEffect(func(
		
		T,
		
		
	)(
		T,
		
	){
		
		return ret0
		
	})
}

func (c _I2_M0Call[T]) ThenEffect(effect func(arg0 T) T) mockkit.FuncBehavior[ func (arg0 T) T ] {
	return &_I2_M0Behavior[T] {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


<<<<<<< HEAD:mocker/internal/types/gen_mock/type_interfaces.go
=======
type _C3_ReadCallSpec struct {
	p its.Matcher[[]byte]
	
	
}

type _C3_ReadCall struct {
	name itskit.Label
	spec _C3_ReadCallSpec
}

func C3_Read_Expects(
	p its.Matcher[[]byte],
	
) _C3_ReadCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C3_ReadCallSpec {}
	spec.p = itskit.Named(
		"p",
		p,
	)
	
	
	return _C3_ReadCall{
		name: itskit.NewLabelWithLocation("func C3_Read"),
		spec: spec,
	}
}

type _C3_ReadBehavior  struct {
	name itskit.Label
	spec _C3_ReadCallSpec
	effect func(p []byte) (n int, err error)
}

func (b *_C3_ReadBehavior) Fn(t mockkit.TestLike) func(p []byte) (n int, err error) {
	return func (
		
		arg0 []byte,
		
		
	) (
		int,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.p
			if matcher == nil {
				matcher = its.Never[[]byte]()
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

func (c _C3_ReadCall) ThenReturn(

	ret0 int,

	ret1 error,

) mockkit.FuncBehavior[ func (p []byte) (n int, err error)  ] {
	return c.ThenEffect(func(
		
		[]byte,
		
		
	)(
		int,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _C3_ReadCall) ThenEffect(effect func(p []byte) (n int, err error)) mockkit.FuncBehavior[ func (p []byte) (n int, err error) ] {
	return &_C3_ReadBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C3_WriteCallSpec struct {
	p its.Matcher[[]byte]
	
	
}

type _C3_WriteCall struct {
	name itskit.Label
	spec _C3_WriteCallSpec
}

func C3_Write_Expects(
	p its.Matcher[[]byte],
	
) _C3_WriteCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C3_WriteCallSpec {}
	spec.p = itskit.Named(
		"p",
		p,
	)
	
	
	return _C3_WriteCall{
		name: itskit.NewLabelWithLocation("func C3_Write"),
		spec: spec,
	}
}

type _C3_WriteBehavior  struct {
	name itskit.Label
	spec _C3_WriteCallSpec
	effect func(p []byte) (n int, err error)
}

func (b *_C3_WriteBehavior) Fn(t mockkit.TestLike) func(p []byte) (n int, err error) {
	return func (
		
		arg0 []byte,
		
		
	) (
		int,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.p
			if matcher == nil {
				matcher = its.Never[[]byte]()
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

func (c _C3_WriteCall) ThenReturn(

	ret0 int,

	ret1 error,

) mockkit.FuncBehavior[ func (p []byte) (n int, err error)  ] {
	return c.ThenEffect(func(
		
		[]byte,
		
		
	)(
		int,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _C3_WriteCall) ThenEffect(effect func(p []byte) (n int, err error)) mockkit.FuncBehavior[ func (p []byte) (n int, err error) ] {
	return &_C3_WriteBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C4_AnotherMethodCallSpec struct {
	
}

type _C4_AnotherMethodCall struct {
	name itskit.Label
	spec _C4_AnotherMethodCallSpec
}

func C4_AnotherMethod_Expects(
) _C4_AnotherMethodCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C4_AnotherMethodCallSpec {}
	
	return _C4_AnotherMethodCall{
		name: itskit.NewLabelWithLocation("func C4_AnotherMethod"),
		spec: spec,
	}
}

type _C4_AnotherMethodBehavior  struct {
	name itskit.Label
	spec _C4_AnotherMethodCallSpec
	effect func()
}

func (b *_C4_AnotherMethodBehavior) Fn(t mockkit.TestLike) func() {
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

func (c _C4_AnotherMethodCall) ThenReturn(

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _C4_AnotherMethodCall) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_C4_AnotherMethodBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C4_CloseCallSpec struct {
	
}

type _C4_CloseCall struct {
	name itskit.Label
	spec _C4_CloseCallSpec
}

func C4_Close_Expects(
) _C4_CloseCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C4_CloseCallSpec {}
	
	return _C4_CloseCall{
		name: itskit.NewLabelWithLocation("func C4_Close"),
		spec: spec,
	}
}

type _C4_CloseBehavior  struct {
	name itskit.Label
	spec _C4_CloseCallSpec
	effect func() error
}

func (b *_C4_CloseBehavior) Fn(t mockkit.TestLike) func() error {
	return func (
		
		
	) (
		error,
		
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

func (c _C4_CloseCall) ThenReturn(

	ret0 error,

) mockkit.FuncBehavior[ func () error  ] {
	return c.ThenEffect(func(
		
		
	)(
		error,
		
	){
		
		return ret0
		
	})
}

func (c _C4_CloseCall) ThenEffect(effect func() error) mockkit.FuncBehavior[ func () error ] {
	return &_C4_CloseBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C4_ReadCallSpec struct {
	p its.Matcher[[]byte]
	
	
}

type _C4_ReadCall struct {
	name itskit.Label
	spec _C4_ReadCallSpec
}

func C4_Read_Expects(
	p its.Matcher[[]byte],
	
) _C4_ReadCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C4_ReadCallSpec {}
	spec.p = itskit.Named(
		"p",
		p,
	)
	
	
	return _C4_ReadCall{
		name: itskit.NewLabelWithLocation("func C4_Read"),
		spec: spec,
	}
}

type _C4_ReadBehavior  struct {
	name itskit.Label
	spec _C4_ReadCallSpec
	effect func(p []byte) (n int, err error)
}

func (b *_C4_ReadBehavior) Fn(t mockkit.TestLike) func(p []byte) (n int, err error) {
	return func (
		
		arg0 []byte,
		
		
	) (
		int,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.p
			if matcher == nil {
				matcher = its.Never[[]byte]()
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

func (c _C4_ReadCall) ThenReturn(

	ret0 int,

	ret1 error,

) mockkit.FuncBehavior[ func (p []byte) (n int, err error)  ] {
	return c.ThenEffect(func(
		
		[]byte,
		
		
	)(
		int,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _C4_ReadCall) ThenEffect(effect func(p []byte) (n int, err error)) mockkit.FuncBehavior[ func (p []byte) (n int, err error) ] {
	return &_C4_ReadBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C4_DotMethodCallSpec struct {
	
}

type _C4_DotMethodCall struct {
	name itskit.Label
	spec _C4_DotMethodCallSpec
}

func C4_DotMethod_Expects(
) _C4_DotMethodCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C4_DotMethodCallSpec {}
	
	return _C4_DotMethodCall{
		name: itskit.NewLabelWithLocation("func C4_DotMethod"),
		spec: spec,
	}
}

type _C4_DotMethodBehavior  struct {
	name itskit.Label
	spec _C4_DotMethodCallSpec
	effect func()
}

func (b *_C4_DotMethodBehavior) Fn(t mockkit.TestLike) func() {
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

func (c _C4_DotMethodCall) ThenReturn(

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _C4_DotMethodCall) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_C4_DotMethodBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C4_MethodCallSpec struct {
	
}

type _C4_MethodCall struct {
	name itskit.Label
	spec _C4_MethodCallSpec
}

func C4_Method_Expects(
) _C4_MethodCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C4_MethodCallSpec {}
	
	return _C4_MethodCall{
		name: itskit.NewLabelWithLocation("func C4_Method"),
		spec: spec,
	}
}

type _C4_MethodBehavior  struct {
	name itskit.Label
	spec _C4_MethodCallSpec
	effect func()
}

func (b *_C4_MethodBehavior) Fn(t mockkit.TestLike) func() {
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

func (c _C4_MethodCall) ThenReturn(

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _C4_MethodCall) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_C4_MethodBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _C5_MCallSpec struct {
	
}

type _C5_MCall struct {
	name itskit.Label
	spec _C5_MCallSpec
}

func C5_M_Expects(
) _C5_MCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _C5_MCallSpec {}
	
	return _C5_MCall{
		name: itskit.NewLabelWithLocation("func C5_M"),
		spec: spec,
	}
}

type _C5_MBehavior  struct {
	name itskit.Label
	spec _C5_MCallSpec
	effect func()
}

func (b *_C5_MBehavior) Fn(t mockkit.TestLike) func() {
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

func (c _C5_MCall) ThenReturn(

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _C5_MCall) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_C5_MBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


>>>>>>> ae47cc0 (fixup! 862407dbe81a7f34e1dc36302d89ed60d4a7cdc2):mocker/internal/generate_test/gen_mock/type_interfaces.go


type _I0Impl struct {
	
	M4 func(s ...string) bool
	M0 func()
	M1 func(arg0 int, arg1 string) ( bool,  error)
	M2 func(i int, s string) (ok bool, err error)
	M3 func(i int, s ...string) (ok bool, err error)
}

func I0_Build(t mockkit.TestLike, spec I0_Spec) testee.I0 {
	impl := _I0Impl{}

	
	if spec.M4 != nil {
		impl.M4 = spec.M4.Fn(t)
	}
	
	if spec.M0 != nil {
		impl.M0 = spec.M0.Fn(t)
	}
	
	if spec.M1 != nil {
		impl.M1 = spec.M1.Fn(t)
	}
	
	if spec.M2 != nil {
		impl.M2 = spec.M2.Fn(t)
	}
	
	if spec.M3 != nil {
		impl.M3 = spec.M3.Fn(t)
	}
	

	return _I0Mock { t: t, impl: impl }
}

type _I0Mock struct {
	t mockkit.TestLike
	impl _I0Impl
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


type I0_Spec struct {
	
	M4 mockkit.FuncBehavior[func (s ...string) bool]
	
	M0 mockkit.FuncBehavior[func ()]
	
	M1 mockkit.FuncBehavior[func (arg0 int, arg1 string) ( bool,  error)]
	
	M2 mockkit.FuncBehavior[func (i int, s string) (ok bool, err error)]
	
	M3 mockkit.FuncBehavior[func (i int, s ...string) (ok bool, err error)]
	
}

type _I1Impl[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	
	M0 func(arg0 S, arg1 T) ( U,  error)
	M1 func(arg0 S)
	M2 func() U
}

func I1_Build[S u_sub.T, T u_sub.T, U testee.X[T]](t mockkit.TestLike, spec I1_Spec[S, T, U]) testee.I1[S, T, U] {
	impl := _I1Impl[S, T, U]{}

	
	if spec.M0 != nil {
		impl.M0 = spec.M0.Fn(t)
	}
	
	if spec.M1 != nil {
		impl.M1 = spec.M1.Fn(t)
	}
	
	if spec.M2 != nil {
		impl.M2 = spec.M2.Fn(t)
	}
	

	return _I1Mock[S, T, U] { t: t, impl: impl }
}

type _I1Mock[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	t mockkit.TestLike
	impl _I1Impl[S, T, U]
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


type I1_Spec[S u_sub.T, T u_sub.T, U testee.X[T]] struct {
	
	M0 mockkit.FuncBehavior[func (arg0 S, arg1 T) ( U,  error)]
	
	M1 mockkit.FuncBehavior[func (arg0 S)]
	
	M2 mockkit.FuncBehavior[func () U]
	
}

type _I2Impl[T ~string] struct {
	
	M0 func(arg0 T) T
}

func I2_Build[T ~string](t mockkit.TestLike, spec I2_Spec[T]) testee.I2[T] {
	impl := _I2Impl[T]{}

	
	if spec.M0 != nil {
		impl.M0 = spec.M0.Fn(t)
	}
	

	return _I2Mock[T] { t: t, impl: impl }
}

type _I2Mock[T ~string] struct {
	t mockkit.TestLike
	impl _I2Impl[T]
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


type I2_Spec[T ~string] struct {
	
	M0 mockkit.FuncBehavior[func (arg0 T) T]
	
}

<<<<<<< HEAD:mocker/internal/types/gen_mock/type_interfaces.go
=======
type _C3Impl struct {
	
	Read func(p []byte) (n int, err error)
	Write func(p []byte) (n int, err error)
}

func C3_Build(t mockkit.TestLike, spec C3_Spec) pkg2.C3 {
	impl := _C3Impl{}

	
	if spec.Read != nil {
		impl.Read = spec.Read.Fn(t)
	}
	
	if spec.Write != nil {
		impl.Write = spec.Write.Fn(t)
	}
	

	return _C3Mock { t: t, impl: impl }
}

type _C3Mock struct {
	t mockkit.TestLike
	impl _C3Impl
}

func (m _C3Mock) Read (
	p []byte,
) (
	n int,
	err error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Read == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C3.Read is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Read(
		p,
	)
}


func (m _C3Mock) Write (
	p []byte,
) (
	n int,
	err error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Write == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C3.Write is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Write(
		p,
	)
}


type C3_Spec struct {
	
	Read mockkit.FuncBehavior[func (p []byte) (n int, err error)]
	
	Write mockkit.FuncBehavior[func (p []byte) (n int, err error)]
	
}

type _C4Impl struct {
	
	AnotherMethod func()
	Close func() error
	Read func(p []byte) (n int, err error)
	DotMethod func()
	Method func()
}

func C4_Build(t mockkit.TestLike, spec C4_Spec) pkg2.C4 {
	impl := _C4Impl{}

	
	if spec.AnotherMethod != nil {
		impl.AnotherMethod = spec.AnotherMethod.Fn(t)
	}
	
	if spec.Close != nil {
		impl.Close = spec.Close.Fn(t)
	}
	
	if spec.Read != nil {
		impl.Read = spec.Read.Fn(t)
	}
	
	if spec.DotMethod != nil {
		impl.DotMethod = spec.DotMethod.Fn(t)
	}
	
	if spec.Method != nil {
		impl.Method = spec.Method.Fn(t)
	}
	

	return _C4Mock { t: t, impl: impl }
}

type _C4Mock struct {
	t mockkit.TestLike
	impl _C4Impl
}

func (m _C4Mock) AnotherMethod (
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.AnotherMethod == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C4.AnotherMethod is not mocked").String(),
		).OrFatal(m.t)
	}

	m.impl.AnotherMethod(
	)
}


func (m _C4Mock) Close (
) (
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Close == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C4.Close is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Close(
	)
}


func (m _C4Mock) Read (
	p []byte,
) (
	n int,
	err error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Read == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C4.Read is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Read(
		p,
	)
}


func (m _C4Mock) DotMethod (
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.DotMethod == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C4.DotMethod is not mocked").String(),
		).OrFatal(m.t)
	}

	m.impl.DotMethod(
	)
}


func (m _C4Mock) Method (
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Method == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("C4.Method is not mocked").String(),
		).OrFatal(m.t)
	}

	m.impl.Method(
	)
}


type C4_Spec struct {
	
	AnotherMethod mockkit.FuncBehavior[func ()]
	
	Close mockkit.FuncBehavior[func () error]
	
	Read mockkit.FuncBehavior[func (p []byte) (n int, err error)]
	
	DotMethod mockkit.FuncBehavior[func ()]
	
	Method mockkit.FuncBehavior[func ()]
	
}

>>>>>>> ae47cc0 (fixup! 862407dbe81a7f34e1dc36302d89ed60d4a7cdc2):mocker/internal/generate_test/gen_mock/type_interfaces.go
