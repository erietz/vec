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

// Calls the callback function, fn, on each element of the vector and returns a
// new vector that contains the results
//
// {fn}: The callback function
//
// {index}: The index of the current element being processed in the vector
func (vec Vec[T]) Map(fn func(t T, index int) T) Vec[T] {
	newVec := Vec[T]{}
	for i, v := range vec {
		newVec = append(newVec, fn(v, i))
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
