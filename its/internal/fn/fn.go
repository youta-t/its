package fn

func EqEq[T comparable](a T) func(T) bool {
	return func(t T) bool { return a == t }
}
