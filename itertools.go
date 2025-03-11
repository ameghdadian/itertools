// Package itertools provides utility functions to work with iterators.
package itertools

import (
	"iter"
	"math/rand/v2"
	"slices"
)

// Concat receives an arbitrary number of slices and returns an iterator to
// iterate on them successively, in the order they are given.
func Concat[Slice []T, T any](seqs ...Slice) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ConcatIter receives an arbitrary number of slices and returns an iterator to
// iterate over the elements, in the order they are given.
func ConcatIter[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Reverse receives an arbitrary number of slices and returns an iterator to
// iterate over the elements, in the reverse order they are given.
//
// For example, given the input:
//
//	Reverse([]int{10, 20, 30}, []int{40, 50, 60})
//
// The iterator following the output:
//
//	{60, 50, 40, 30, 20, 10}
func Reverse[Slice []T, T any](seqs ...Slice) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(seqs) - 1; i >= 0; i-- {
			for j := len(seqs[i]) - 1; j >= 0; j-- {
				if !yield(seqs[i][j]) {
					return
				}
			}
		}
	}
}

// Shuffle receives a slice and returns an iterator over its shuffled elements.
//
// Note: In order not to allocate extra memory, it shuffles given slice(array) in place.
// If you need the original slice(array) untouched, you can clone your slice using [slices.Clone]
// and then pass it as the argument.
func Shuffle[Slice []T, T any](seq Slice) iter.Seq[T] {
	return func(yield func(T) bool) {
		swapFunc := func(i int, j int) {
			seq[i], seq[j] = seq[j], seq[i]
		}
		rand.Shuffle(len(seq), swapFunc)
		for v := range slices.Values(seq) {
			if !yield(v) {
				return
			}
		}
	}
}

// Filter returns an iterator that contains the elements of seq for which filterFn
// returns true.
func Filter[Slice []T, T comparable](seq Slice, filterFn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range seq {
			if filterFn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map executes a user-supplied function on each element of the slice and returns
// an iterator over modified elements.
func Map[Slice []T, T any](seq Slice, mapFn func(int, T) T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range seq {
			if !yield(i, mapFn(i, v)) {
				return
			}
		}
	}
}

// ForEach executes a user-supplied function on each element of the sequence. It does not
// return any output.
func ForEach[T any](seq iter.Seq[T], fn func(int, T)) {
	var indx int
	for v := range seq {
		fn(indx, v)
		indx += 1
	}
}

// Reduce executes a user-supplied "reducer" function on each element of the array,
// in order, passing in the return value from the calculation on the preceding element.
func Reduce[T any](seq iter.Seq[T], reducer func(acc T, cur T) T, init T) T {
	for v := range seq {
		init = reducer(init, v)
	}

	return init
}
