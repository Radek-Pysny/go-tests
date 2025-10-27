package put

import (
	"iter"
)

func EverySecond[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		b := true
		for _, x := range slice {
			b = !b
			if b {
				if !yield(x) {
					return
				}
			}
		}
	}
}

func Enumerate[T any](
	it iter.Seq[T],
) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := -1
		for x := range it {
			i++
			if !yield(i, x) {
				return
			}
		}
	}
}

func EnumerateFrom[T any](
	it iter.Seq[T], from int,
) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		from--
		for x := range it {
			from++
			if !yield(from, x) {
				return
			}
		}
	}
}

func Pairs[T any](
	seq iter.Seq[T],
) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		for {
			v1, exists1 := next()
			if !exists1 {
				return
			}
			v2, exists2 := next()
			if !yield(v1, v2) || !exists2 {
				return
			}
		}
	}
}
