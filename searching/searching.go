// Package searching provides basic search algorithms.
package searching

import "cmp"

// LinearSearch scans the slice from left to right.
// It returns the index of target, or -1 if not found. Time: O(n).
func LinearSearch[T comparable](a []T, target T) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}

// BinarySearch finds target in a slice sorted in ascending order.
// It returns the index of target, or -1 if not found. Time: O(log n).
func BinarySearch[T cmp.Ordered](a []T, target T) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // avoids overflow of (lo+hi)
		switch {
		case a[mid] == target:
			return mid
		case a[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive is the recursive form of BinarySearch.
func BinarySearchRecursive[T cmp.Ordered](a []T, target T) int {
	return binarySearch(a, target, 0, len(a)-1)
}

func binarySearch[T cmp.Ordered](a []T, target T, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	switch {
	case a[mid] == target:
		return mid
	case a[mid] < target:
		return binarySearch(a, target, mid+1, hi)
	default:
		return binarySearch(a, target, lo, mid-1)
	}
}
