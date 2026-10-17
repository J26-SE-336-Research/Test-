// Command algorithms runs a short demo of every package in this module.
package main

import (
	"fmt"

	"github.com/hiruni/algorithms/graph"
	"github.com/hiruni/algorithms/searching"
	"github.com/hiruni/algorithms/sorting"
)

func main() {
	nums := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("== Sorting ==")
	fmt.Println("input:         ", nums)
	fmt.Println("BubbleSort:    ", sorting.BubbleSort(nums))
	fmt.Println("InsertionSort: ", sorting.InsertionSort(nums))
	fmt.Println("MergeSort:     ", sorting.MergeSort(nums))
	fmt.Println("QuickSort:     ", sorting.QuickSort(nums))

	sorted := sorting.MergeSort(nums)
	fmt.Println("\n== Searching ==")
	fmt.Println("LinearSearch(43): ", searching.LinearSearch(nums, 43))
	fmt.Println("BinarySearch(43): ", searching.BinarySearch(sorted, 43))
	fmt.Println("BinarySearch(5):  ", searching.BinarySearch(sorted, 5))

	g := graph.New()
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(2, 1, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 5)

	fmt.Println("\n== Graph ==")
	fmt.Println("BFS(0):      ", g.BFS(0))
	fmt.Println("DFS(0):      ", g.DFS(0))
	fmt.Println("Dijkstra(0): ", g.Dijkstra(0))
	path, cost := g.ShortestPath(0, 3)
	fmt.Printf("ShortestPath(0, 3): %v (cost %d)\n", path, cost)
}
