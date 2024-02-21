// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	testee "github.com/youta-t/its/mocker/internal/example"
	
)

type _SessionStoreReturnFixture struct {
	userId string
	
	ok bool
	
}

type _SessionStoreReturn struct {
	fixture _SessionStoreReturnFixture
}

func (rfx _SessionStoreReturn) Get() (
	string,
	
	bool,
	
) {
	return rfx.fixture.userId, rfx.fixture.ok
}

type _SessionStoreCallSpec struct {
	cookie its.Matcher[string]
	
	
}

type _SessionStoreCall struct {
	name itskit.Label
	spec _SessionStoreCallSpec
}

func NewSessionStoreCall(
	cookie its.Matcher[string],
	
) _SessionStoreCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _SessionStoreCallSpec {}
	spec.cookie = itskit.Named(
		"cookie",
		cookie,
	)
	
	
	return _SessionStoreCall{
		name: itskit.NewLabelWithLocation("func SessionStore"),
		spec: spec,
	}
}

type SessionStoreBehaviour  struct {
	name itskit.Label
	spec _SessionStoreCallSpec
	effect func(cookie string) (userId string, ok bool)
}

func (b SessionStoreBehaviour) Mock(t interface { Error(...any) }) func(cookie string) (userId string, ok bool) {
	return func (
		
		arg0 string,
		
		
	) (
		string,
		bool,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.cookie
			if matcher == nil {
				matcher = its.Never[string]()
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

func (c _SessionStoreCall) ThenReturn(

	ret0 string,

	ret1 bool,

) SessionStoreBehaviour {
	return c.ThenEffect(func(
		
		string,
		
		
	)(
		string,
		bool,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _SessionStoreCall) ThenEffect(effect func(cookie string) (userId string, ok bool)) SessionStoreBehaviour {
	return SessionStoreBehaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _UserRegistry_GetReturnFixture struct {
	ret0 testee.User
	
	ret1 error
	
}

type _UserRegistry_GetReturn struct {
	fixture _UserRegistry_GetReturnFixture
}

func (rfx _UserRegistry_GetReturn) Get() (
	testee.User,
	
	error,
	
) {
	return rfx.fixture.ret0, rfx.fixture.ret1
}

type _UserRegistry_GetCallSpec struct {
	userId its.Matcher[string]
	
	
}

type _UserRegistry_GetCall struct {
	name itskit.Label
	spec _UserRegistry_GetCallSpec
}

func NewUserRegistry_GetCall(
	userId its.Matcher[string],
	
) _UserRegistry_GetCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_GetCallSpec {}
	spec.userId = itskit.Named(
		"userId",
		userId,
	)
	
	
	return _UserRegistry_GetCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Get"),
		spec: spec,
	}
}

type UserRegistry_GetBehaviour  struct {
	name itskit.Label
	spec _UserRegistry_GetCallSpec
	effect func(userId string) ( testee.User,  error)
}

func (b UserRegistry_GetBehaviour) Mock(t interface { Error(...any) }) func(userId string) ( testee.User,  error) {
	return func (
		
		arg0 string,
		
		
	) (
		testee.User,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.userId
			if matcher == nil {
				matcher = its.Never[string]()
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

func (c _UserRegistry_GetCall) ThenReturn(

	ret0 testee.User,

	ret1 error,

) UserRegistry_GetBehaviour {
	return c.ThenEffect(func(
		
		string,
		
		
	)(
		testee.User,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _UserRegistry_GetCall) ThenEffect(effect func(userId string) ( testee.User,  error)) UserRegistry_GetBehaviour {
	return UserRegistry_GetBehaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _UserRegistry_UpdateReturnFixture struct {
	ret0 error
	
}

type _UserRegistry_UpdateReturn struct {
	fixture _UserRegistry_UpdateReturnFixture
}

func (rfx _UserRegistry_UpdateReturn) Get() (
	error,
	
) {
	return rfx.fixture.ret0
}

type _UserRegistry_UpdateCallSpec struct {
	arg0 its.Matcher[testee.User]
	
	
}

type _UserRegistry_UpdateCall struct {
	name itskit.Label
	spec _UserRegistry_UpdateCallSpec
}

func NewUserRegistry_UpdateCall(
	arg0 its.Matcher[testee.User],
	
) _UserRegistry_UpdateCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_UpdateCallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _UserRegistry_UpdateCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Update"),
		spec: spec,
	}
}

type UserRegistry_UpdateBehaviour  struct {
	name itskit.Label
	spec _UserRegistry_UpdateCallSpec
	effect func(arg0 testee.User) error
}

func (b UserRegistry_UpdateBehaviour) Mock(t interface { Error(...any) }) func(arg0 testee.User) error {
	return func (
		
		arg0 testee.User,
		
		
	) (
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
				matcher = its.Never[testee.User]()
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

func (c _UserRegistry_UpdateCall) ThenReturn(

	ret0 error,

) UserRegistry_UpdateBehaviour {
	return c.ThenEffect(func(
		
		testee.User,
		
		
	)(
		error,
		
	){
		
		return ret0
		
	})
}

func (c _UserRegistry_UpdateCall) ThenEffect(effect func(arg0 testee.User) error) UserRegistry_UpdateBehaviour {
	return UserRegistry_UpdateBehaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}



type _UserRegistry_DeleteReturnFixture struct {
	ret0 error
	
}

type _UserRegistry_DeleteReturn struct {
	fixture _UserRegistry_DeleteReturnFixture
}

func (rfx _UserRegistry_DeleteReturn) Get() (
	error,
	
) {
	return rfx.fixture.ret0
}

type _UserRegistry_DeleteCallSpec struct {
	arg0 its.Matcher[testee.User]
	
	
}

type _UserRegistry_DeleteCall struct {
	name itskit.Label
	spec _UserRegistry_DeleteCallSpec
}

func NewUserRegistry_DeleteCall(
	arg0 its.Matcher[testee.User],
	
) _UserRegistry_DeleteCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_DeleteCallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _UserRegistry_DeleteCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Delete"),
		spec: spec,
	}
}

type UserRegistry_DeleteBehaviour  struct {
	name itskit.Label
	spec _UserRegistry_DeleteCallSpec
	effect func(arg0 testee.User) error
}

func (b UserRegistry_DeleteBehaviour) Mock(t interface { Error(...any) }) func(arg0 testee.User) error {
	return func (
		
		arg0 testee.User,
		
		
	) (
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
				matcher = its.Never[testee.User]()
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

func (c _UserRegistry_DeleteCall) ThenReturn(

	ret0 error,

) UserRegistry_DeleteBehaviour {
	return c.ThenEffect(func(
		
		testee.User,
		
		
	)(
		error,
		
	){
		
		return ret0
		
	})
}

func (c _UserRegistry_DeleteCall) ThenEffect(effect func(arg0 testee.User) error) UserRegistry_DeleteBehaviour {
	return UserRegistry_DeleteBehaviour {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}




type UserRegistryImpl struct {
	
	Get func(userId string) ( testee.User,  error)
	Update func(arg0 testee.User) error
	Delete func(arg0 testee.User) error
}

func NewMockedUserRegistry(t interface { Fatal(...any) } ,impl UserRegistryImpl) testee.UserRegistry {
	return _UserRegistryMock { t: t, impl: impl }
}

type _UserRegistryMock struct {
	t interface { Fatal(...any) }
	impl UserRegistryImpl
}

func (m _UserRegistryMock) Get (
	userId string,
) (
	testee.User,
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Get == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Get is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Get(
		userId,
	)
}


func (m _UserRegistryMock) Update (
	arg0 testee.User,
) (
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Update == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Update is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Update(
		arg0,
	)
}


func (m _UserRegistryMock) Delete (
	arg0 testee.User,
) (
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Delete == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Delete is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Delete(
		arg0,
	)
}


