package sorting

import (
	"slices"
	"testing"
)

func TestSorts(t *testing.T) {
	sorts := map[string]func([]int) []int{
		"BubbleSort":    BubbleSort[int],
		"InsertionSort": InsertionSort[int],
		"MergeSort":     MergeSort[int],
		"QuickSort":     QuickSort[int],
	}
	cases := [][]int{
		{},
		{1},
		{2, 1},
		{5, 2, 9, 1, 5, 6},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, -1, 0, -7, 8, 3},
	}

	for name, sortFn := range sorts {
		for _, in := range cases {
			original := slices.Clone(in)
			want := slices.Clone(in)
			slices.Sort(want)

			got := sortFn(in)
			if !slices.Equal(got, want) {
				t.Errorf("%s(%v) = %v, want %v", name, in, got, want)
			}
			if !slices.Equal(in, original) {
				t.Errorf("%s modified its input: %v", name, in)
			}
		}
	}
}
