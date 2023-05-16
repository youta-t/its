package try

type Fataler interface {
	Fatal(...any)
}

type Try[T any] interface {
	Get() (T, error)
	OrFatal(Fataler) T
}

func To[T any](value T, err error) Try[T] {
	if err != nil {
		return &ng[T]{err: err}
	}
	return &ok[T]{value: value}
}

type ok[T any] struct {
	value T
}

func (o ok[T]) Get() (T, error) {
	return o.value, nil
}

func (o ok[T]) OrFatal(Fataler) T {
	return o.value
}

type ng[T any] struct {
	err error
}

func (n ng[T]) Get() (T, error) {
	return *new(T), n.err
}

func (n ng[T]) OrFatal(f Fataler) T {
	f.Fatal(n.err)
	return *new(T)
}
