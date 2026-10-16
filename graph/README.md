# graph

A weighted, directed graph (adjacency list) with traversal and shortest-path algorithms.

| Method          | What it does                                        | Time             |
|-----------------|-----------------------------------------------------|------------------|
| `BFS`           | Breadth-first traversal order from a start node     | O(V + E)         |
| `DFS`           | Depth-first (pre-order) traversal from a start node | O(V + E)         |
| `Dijkstra`      | Shortest distance from start to every reachable node| O((V + E) log V) |
| `ShortestPath`  | Shortest path and its cost between two nodes        | O((V + E) log V) |

Dijkstra's algorithm requires non-negative edge weights.

## Usage

```go
import "github.com/hiruni/algorithms/graph"

g := graph.New()
g.AddEdge(0, 1, 1)
g.AddEdge(0, 2, 4)
g.AddEdge(2, 1, 1)
g.AddEdge(1, 3, 2)

g.BFS(0)                     // [0 1 2 3]
g.Dijkstra(0)                // map[0:0 1:1 2:4 3:3]
path, cost := g.ShortestPath(0, 3) // [0 1 3], 3
```

Use `AddUndirectedEdge(a, b, w)` to add an edge in both directions.

## Tests

```sh
go test ./graph
```
