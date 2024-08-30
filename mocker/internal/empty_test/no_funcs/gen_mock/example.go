// Code generated -- DO NOT EDIT
package gen_mock

import (
	
	
	mockkit "github.com/youta-t/its/mocker/mockkit"
	pkg1 "github.com/youta-t/its/mocker/internal/empty_test/no_funcs"
	
)



type _EmptyImpl struct {
	
}

func Empty_Build(t mockkit.TestLike, spec Empty_Spec) pkg1.Empty {
	impl := _EmptyImpl{}

	

	return _EmptyMock { t: t, impl: impl }
}

type _EmptyMock struct {
	t mockkit.TestLike
	impl _EmptyImpl
}

type Empty_Spec struct {
	
}

