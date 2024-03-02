package its

import "github.com/youta-t/its/itskit"

// ForItems broadcasts matcherFactory for each elements, and wrap with s.
//
// This provides shorthands of Matchers using slice of Matchers.
//
// # Example
//
//	ForItems(Some, EqEq, []int{1, 2, 3, 4, 5})
//
// is, equiverent to
//
//	Some(
//		EqEq(1), EqEq(2), EqEq(3), EqEq(4), EqEq(5),
//	)
//
// # Args
//
// - s func(matchers []Matcher[T]) Matcher[X]: wraps matchers comes from matcherFactory
//
// - matcherFactory func(U) Matcher[T]: it applied for each items in wants to generate matchers for them
//
// - wants []U: values.
func ForItems[T, U, X any, F func(U) Matcher[T], S func(...Matcher[T]) Matcher[X]](
	s S, matcherFactory F, wants []U,
) Matcher[X] {
	cancel := itskit.SkipStack()
	defer cancel()

	ret := make([]Matcher[T], len(wants))
	for i := range wants {
		ret[i] = matcherFactory(wants[i])
	}
	return s(ret...)
}

// ForEntries broadcasts matcherFactory for each values in map, and wrap with s.
//
// This provides shordhands of Matchers using map of Matchers.
//
// # Example
//
//	ForEntries(Map, EqEq, map[string]int{"a": 1, "b": 2})
//
// is, equiverent to
//
//	Map(map[string]Matcher[int]{
//		"a": its.EqEq(1), "b": its.EqEq(2)
//	})
//
// # Args
//
// - s func(map[K]Matcher[T] Matcher[X]): wraps a map of matchers comes from matcherFactory
//
// - matcherFactory func(U)Matcher[T]: it applied for each values in enteies to generate map of Matchers.
//
// - wants: map entries.
func ForEntries[
	K comparable, T, U, X any, F func(U) Matcher[T], S func(map[K]Matcher[T]) Matcher[X],
](
	s S,
	matcherFactory F,
	wants map[K]U,
) Matcher[X] {
	cancel := itskit.SkipStack()
	defer cancel()

	m := map[K]Matcher[T]{}
	for k := range wants {
		m[k] = matcherFactory(wants[k])
	}

	return s(m)
}
