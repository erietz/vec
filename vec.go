package vec

import "fmt"

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

func main() {
	v1 := Vec[int]{1, 32, 23, 2432, 4324, 4, 34, 2, 3, 4, 2, 34, 234}
	v2 := v1.Filter(func(i int) bool { return i == 2 })
	v3 := v1.Map(func(i int) int { return i * 2 })
	v4 := Map(v1, func(i int) string { return fmt.Sprintf("foo%d", i) })
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}
