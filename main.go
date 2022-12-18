package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func qsort() {
	arr := rand.Perm(10000000)

	seqQsort(arr)

	fmt.Println(sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] }))

	arr = rand.Perm(10000000)

	parQsort(arr)

	fmt.Println(sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] }))
}

func bfs() {
	l := 100
	graph := cubicGraph(l)

	var seq, par []int32
	seqBfs(graph, 0)
	for _, n := range graph.nodes {
		seq = append(seq, n.distance)
	}
	parBfs(graph, 0)
	for _, n := range graph.nodes {
		par = append(par, n.distance)
	}

	for i := 0; i < l; i++ {
		if seq[i] != par[i] {
			fmt.Println("not equal")
			return
		}
	}
	fmt.Println("equal")
}

func main() {
	qsort()
	bfs()
}
