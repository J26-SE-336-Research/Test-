// Package sorting provides classic sorting algorithms.
// Every function returns a new sorted slice and leaves the input unchanged.
package sorting

import "cmp"

// BubbleSort repeatedly swaps adjacent out-of-order elements.
// Time: O(n^2) worst/average, O(n) best (already sorted). Space: O(n) for the copy.
func BubbleSort[T cmp.Ordered](in []T) []T {
	a := clone(in)
	for i := 0; i < len(a)-1; i++ {
		swapped := false
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return a
}

// InsertionSort builds the sorted slice one element at a time.
// Time: O(n^2) worst, O(n) best. Good for small or nearly sorted inputs.
func InsertionSort[T cmp.Ordered](in []T) []T {
	a := clone(in)
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}

// MergeSort splits the slice in half, sorts each half, and merges them.
// Time: O(n log n) in all cases. Space: O(n). Stable.
func MergeSort[T cmp.Ordered](in []T) []T {
	if len(in) <= 1 {
		return clone(in)
	}
	mid := len(in) / 2
	return merge(MergeSort(in[:mid]), MergeSort(in[mid:]))
}

func merge[T cmp.Ordered](left, right []T) []T {
	out := make([]T, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}
	out = append(out, left[i:]...)
	return append(out, right[j:]...)
}

// QuickSort partitions around a pivot and recursively sorts each side.
// Time: O(n log n) average, O(n^2) worst. Space: O(log n) stack on average.
func QuickSort[T cmp.Ordered](in []T) []T {
	a := clone(in)
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort[T cmp.Ordered](a []T, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

// partition uses the Lomuto scheme with the middle element as pivot
// to avoid worst-case behaviour on already sorted input.
func partition[T cmp.Ordered](a []T, lo, hi int) int {
	mid := lo + (hi-lo)/2
	a[mid], a[hi] = a[hi], a[mid]
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func clone[T any](in []T) []T {
	out := make([]T, len(in))
	copy(out, in)
	return out
}
