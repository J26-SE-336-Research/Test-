package graph

import (
	"maps"
	"math"
	"slices"
	"testing"
)

// 0 --1--> 1 --2--> 3
// |        ^        ^
// 4        1        5
// v        |        |
// 2 -------+--------+
func sample() *Graph {
	g := New()
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(2, 1, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 5)
	return g
}

func TestBFS(t *testing.T) {
	got := sample().BFS(0)
	want := []int{0, 1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("BFS = %v, want %v", got, want)
	}
}

func TestDFS(t *testing.T) {
	got := sample().DFS(0)
	want := []int{0, 1, 3, 2}
	if !slices.Equal(got, want) {
		t.Errorf("DFS = %v, want %v", got, want)
	}
}

func TestDijkstra(t *testing.T) {
	got := sample().Dijkstra(0)
	want := map[int]int{0: 0, 1: 1, 2: 4, 3: 3}
	if !maps.Equal(got, want) {
		t.Errorf("Dijkstra = %v, want %v", got, want)
	}
}

func TestShortestPath(t *testing.T) {
	g := sample()

	path, cost := g.ShortestPath(0, 3)
	if !slices.Equal(path, []int{0, 1, 3}) || cost != 3 {
		t.Errorf("ShortestPath(0,3) = %v, %d; want [0 1 3], 3", path, cost)
	}

	path, cost = g.ShortestPath(3, 0)
	if path != nil || cost != math.MaxInt {
		t.Errorf("ShortestPath(3,0) = %v, %d; want nil, MaxInt", path, cost)
	}

	path, cost = g.ShortestPath(2, 2)
	if !slices.Equal(path, []int{2}) || cost != 0 {
		t.Errorf("ShortestPath(2,2) = %v, %d; want [2], 0", path, cost)
	}
}
