// Package graph provides a simple weighted, directed graph and
// traversal / shortest-path algorithms over it.
package graph

import (
	"container/heap"
	"math"
)

// Graph is a directed graph stored as an adjacency list.
// Nodes are identified by int IDs.
type Graph struct {
	adj map[int][]Edge
}

// Edge is a directed, weighted edge to node To.
type Edge struct {
	To     int
	Weight int
}

// New returns an empty graph.
func New() *Graph {
	return &Graph{adj: make(map[int][]Edge)}
}

// AddEdge adds a directed edge from -> to with the given weight.
func (g *Graph) AddEdge(from, to, weight int) {
	g.adj[from] = append(g.adj[from], Edge{To: to, Weight: weight})
	if _, ok := g.adj[to]; !ok {
		g.adj[to] = nil
	}
}

// AddUndirectedEdge adds edges in both directions.
func (g *Graph) AddUndirectedEdge(a, b, weight int) {
	g.AddEdge(a, b, weight)
	g.AddEdge(b, a, weight)
}

// BFS returns nodes in breadth-first order starting from start.
// Time: O(V + E).
func (g *Graph) BFS(start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	var order []int

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		order = append(order, n)
		for _, e := range g.adj[n] {
			if !visited[e.To] {
				visited[e.To] = true
				queue = append(queue, e.To)
			}
		}
	}
	return order
}

// DFS returns nodes in depth-first (pre-order) order starting from start.
// Time: O(V + E).
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	var order []int
	var visit func(int)
	visit = func(n int) {
		visited[n] = true
		order = append(order, n)
		for _, e := range g.adj[n] {
			if !visited[e.To] {
				visit(e.To)
			}
		}
	}
	visit(start)
	return order
}

// Dijkstra returns the shortest distance from start to every reachable node.
// Unreachable nodes are absent from the result. Edge weights must be non-negative.
// Time: O((V + E) log V).
func (g *Graph) Dijkstra(start int) map[int]int {
	dist := map[int]int{start: 0}
	pq := &priorityQueue{{node: start, dist: 0}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.dist > dist[cur.node] {
			continue // stale entry
		}
		for _, e := range g.adj[cur.node] {
			nd := cur.dist + e.Weight
			if d, ok := dist[e.To]; !ok || nd < d {
				dist[e.To] = nd
				heap.Push(pq, item{node: e.To, dist: nd})
			}
		}
	}
	return dist
}

// ShortestPath returns the node sequence of the shortest path from start
// to end and its total weight. It returns (nil, math.MaxInt) if end is unreachable.
func (g *Graph) ShortestPath(start, end int) ([]int, int) {
	dist := map[int]int{start: 0}
	prev := make(map[int]int)
	pq := &priorityQueue{{node: start, dist: 0}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.node == end {
			break
		}
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range g.adj[cur.node] {
			nd := cur.dist + e.Weight
			if d, ok := dist[e.To]; !ok || nd < d {
				dist[e.To] = nd
				prev[e.To] = cur.node
				heap.Push(pq, item{node: e.To, dist: nd})
			}
		}
	}

	total, ok := dist[end]
	if !ok {
		return nil, math.MaxInt
	}
	path := []int{end}
	for n := end; n != start; {
		n = prev[n]
		path = append([]int{n}, path...)
	}
	return path, total
}

type item struct {
	node, dist int
}

// priorityQueue is a min-heap of items ordered by dist.
type priorityQueue []item

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any)        { *pq = append(*pq, x.(item)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}
