# Test-

A small collection of classic algorithms implemented in Go, with tests.

## Packages

| Package                    | Algorithms                                              |
|----------------------------|---------------------------------------------------------|
| [`sorting`](sorting/)      | Bubble Sort, Insertion Sort, Merge Sort, Quick Sort     |
| [`searching`](searching/)  | Linear Search, Binary Search (iterative and recursive)  |
| [`graph`](graph/)          | BFS, DFS, Dijkstra's shortest path                      |

Each package has its own README with complexity tables and usage examples.

## Requirements

Go 1.22 or newer.

## Run the demo

```sh
go run .
```

## Run the tests

```sh
go test ./...
```

## Project layout

```text
.
├── main.go           # demo that uses every package
├── go.mod
├── sorting/
├── searching/
└── graph/
```
