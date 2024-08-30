// Code generated -- DO NOT EDIT
package gen_mock

import (
	
	itskit "github.com/youta-t/its/itskit"
	mockkit "github.com/youta-t/its/mocker/mockkit"
	
)

type _FCallSpec struct {
	
}

type _FCall struct {
	name itskit.Label
	spec _FCallSpec
}

func F_Expects(
) _FCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _FCallSpec {}
	
	return _FCall{
		name: itskit.NewLabelWithLocation("func F"),
		spec: spec,
	}
}

type _FBehavior  struct {
	name itskit.Label
	spec _FCallSpec
	effect func()
}

func (b *_FBehavior) Fn(t mockkit.TestLike) func() {
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

func (c _FCall) ThenReturn(

) mockkit.FuncBehavior[ func ()  ] {
	return c.ThenEffect(func(
		
		
	){
		
	})
}

func (c _FCall) ThenEffect(effect func()) mockkit.FuncBehavior[ func () ] {
	return &_FBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}




