# searching

Search algorithms that return the index of the target, or `-1` if it is not present.

| Algorithm                | Time     | Requires sorted input |
|--------------------------|----------|-----------------------|
| `LinearSearch`           | O(n)     | No                    |
| `BinarySearch`           | O(log n) | Yes (ascending)       |
| `BinarySearchRecursive`  | O(log n) | Yes (ascending)       |

## Usage

```go
import "github.com/hiruni/algorithms/searching"

nums := []int{1, 3, 5, 7, 9}
searching.BinarySearch(nums, 7)  // 3
searching.LinearSearch(nums, 4)  // -1
```

## Tests

```sh
go test ./searching
```
