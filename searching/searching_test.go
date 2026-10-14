package searching

import "testing"

func TestSearches(t *testing.T) {
	sorted := []int{1, 3, 5, 7, 9, 11}
	cases := []struct {
		target int
		want   int
	}{
		{1, 0},
		{7, 3},
		{11, 5},
		{4, -1},
		{0, -1},
		{12, -1},
	}

	searches := map[string]func([]int, int) int{
		"LinearSearch":          LinearSearch[int],
		"BinarySearch":          BinarySearch[int],
		"BinarySearchRecursive": BinarySearchRecursive[int],
	}

	for name, search := range searches {
		for _, c := range cases {
			if got := search(sorted, c.target); got != c.want {
				t.Errorf("%s(%v, %d) = %d, want %d", name, sorted, c.target, got, c.want)
			}
		}
		if got := search(nil, 1); got != -1 {
			t.Errorf("%s on empty slice = %d, want -1", name, got)
		}
	}
}
