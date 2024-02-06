package list

type List[T any] interface {
	// Len returns length of this list
	Len() int

	// Nth returns N-th value in this list
	//
	// If n < 0, n is assumed as "negative offset",
	// and it returns l.Nth(l.Len() - abs(n)).
	//
	// # Return
	//
	// - T: Nth value
	//
	// - bool: if N is out of bound, it is false. Otherwise true.
	Nth(int) (T, bool)

	// Find returns the first value which the given function returns true.
	//
	// # Return
	//
	// - T : found value
	//
	// - bool: true if found.
	Find(func(T) bool) (T, bool)

	// PopLeft pops leftmost value(= index 0).
	//
	// After calling PopLeft(), 0-th value is deleted from the list.
	//
	// # Return
	//
	// - T: popped value.
	//
	// - bool: if 0 < Len(), true. Otherwise false.
	PopLeft() (T, bool)

	// PopLeft pops rightmost value(= index Len() - 1).
	//
	// After calling PopRight(), (Len()-1)-th value is deleted from the list.
	//
	// # Return
	//
	// - T: popped value.
	//
	// - bool: if 0 < Len(), true. Otherwise false.
	PopRight() (T, bool)

	// Append append new elements into the list
	Append(...T)

	// Slice retreives elements in the list as a slice.
	Slice() []T
}

func New[T any](initial []T) List[T] {
	l := make([]T, len(initial))
	copy(l, initial)
	return (*list[T])(&l)
}

type list[T any] []T

func (l *list[T]) Len() int {
	return len(*l)
}

func (l *list[T]) Nth(n int) (T, bool) {
	if l.Len() == 0 {
		return *new(T), false
	}
	if n < 0 {
		return l.Nth(l.Len() + n)
	}

	if l.Len()-1 < n {
		return *new(T), false
	}
	return (*l)[n], true
}

func (l *list[T]) Last() (T, bool) {
	return l.Nth(l.Len() - 1)
}

func (l *list[T]) PopLeft() (T, bool) {
	if l.Len() <= 0 {
		return *new(T), false
	}
	v, rest := (*l)[0], (*l)[1:]
	*l = rest
	return v, true
}

func (l *list[T]) PopRight() (T, bool) {
	if l.Len() <= 0 {
		return *new(T), false
	}
	rest, v := (*l)[:l.Len()-1], (*l)[l.Len()-1]
	*l = rest
	return v, true
}

func (l *list[T]) Append(v ...T) {
	*l = append(*l, v...)
}

func (l *list[T]) Slice() []T {
	return *l
}

func (l *list[T]) Find(p func(T) bool) (T, bool) {
	for _, v := range *l {
		if p(v) {
			return v, true
		}
	}
	return *new(T), false
}
