# sorting

Generic sorting algorithms for any `cmp.Ordered` type (ints, floats, strings).
Each function returns a new sorted slice and does not modify the input.

| Algorithm      | Best       | Average    | Worst      | Space    | Stable |
|----------------|------------|------------|------------|----------|--------|
| Bubble Sort    | O(n)       | O(n²)      | O(n²)      | O(1)*    | Yes    |
| Insertion Sort | O(n)       | O(n²)      | O(n²)      | O(1)*    | Yes    |
| Merge Sort     | O(n log n) | O(n log n) | O(n log n) | O(n)     | Yes    |
| Quick Sort     | O(n log n) | O(n log n) | O(n²)      | O(log n)*| No     |

\* Plus O(n) for the copy of the input.

## Usage

```go
import "github.com/hiruni/algorithms/sorting"

sorted := sorting.MergeSort([]int{5, 2, 9, 1})        // [1 2 5 9]
words  := sorting.QuickSort([]string{"pear", "apple"}) // [apple pear]
```

## Tests

```sh
go test ./sorting
```
