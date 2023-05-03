package vec

type Vec[T any] []T

func (vec Vec[T]) Filter(fn func(t T) bool) Vec[T] {
	newVec := Vec[T]{}
	for _, v := range vec {
		if fn(v) {
			newVec = append(newVec, v)
		}
	}
	return newVec
}

func (vec Vec[T]) Map(fn func(t T) T) Vec[T] {
	newVec := Vec[T]{}
	for _, v := range vec {
		newVec = append(newVec, fn(v))
	}
	return newVec
}

func Map[X, Y any](vec Vec[X], fn func(x X) Y) Vec[Y] {
	newVec := Vec[Y]{}
	for _, v := range vec {
		newVec = append(newVec, fn(v))
	}
	return newVec
}
