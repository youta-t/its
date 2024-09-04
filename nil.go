package its

import (
	"fmt"
	"reflect"

	"github.com/youta-t/its/itskit"
	"github.com/youta-t/its/itskit/itsio"
)

type nilMatcher[T any] struct {
	label itskit.Label
}

func (n nilMatcher[T]) Match(got T) itskit.Match {
	rg := reflect.ValueOf(got)
	typ := reflect.ValueOf(*new(T))
	switch k := typ.Kind(); k {
	case reflect.Pointer:
		isnil := rg.IsNil()
		if isnil {
			return itskit.OK(n.label.Fill("<nil>"))
		}
		return itskit.NG(n.label.Fill(rg.Elem().Interface()))
	case reflect.Chan, reflect.Func:
		return itskit.NewMatch(
			rg.IsNil(),
			n.label.Fill(fmt.Sprintf("%T", rg.Interface())),
		)
	case reflect.Interface, reflect.Map, reflect.Slice:
		return itskit.NewMatch(
			rg.IsNil(),
			n.label.Fill(rg.Interface()),
		)
	default:
		isnil := any(got) == nil
		if isnil {
			return itskit.OK(n.label.Fill("<nil>"))
		}
		return itskit.NG(n.label.Fill(rg.Interface()))
	}
}

func (n nilMatcher[T]) Write(w itsio.Writer) error {
	return n.label.Write(w)
}

func (n nilMatcher[T]) String() string {
	return itskit.MatcherToString(n)
}

func Nil[T any]() Matcher[T] {
	cancel := itskit.SkipStack()
	defer cancel()

	return nilMatcher[T]{
		label: itskit.NewLabelWithLocation("(%+v) is nil", itskit.Got),
	}
}
