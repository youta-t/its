package fn

func Map[A, B any](as []A, mapper func(A) B) []B {
	ret := make([]B, len(as))
	for nth, a := range as {
		ret[nth] = mapper(a)
	}
	return ret
}
