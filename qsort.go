package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
)

const (
	minQsort    = 50
	minParallel = 1000

	maxGoroutines = 8
)

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		h := arr[i]
		j := i - 1

		for j >= 0 && h < arr[j] {
			arr[j+1] = arr[j]

			j--
		}

		arr[j+1] = h
	}
}

func partition(arr []int) ([]int, []int) {
	arr[len(arr)/2], arr[0] = arr[0], arr[len(arr)/2]

	pivot := arr[0]
	mid := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			mid += 1
			arr[i], arr[mid] = arr[mid], arr[i]
		}
	}

	arr[0], arr[mid] = arr[mid], arr[0]

	if mid < len(arr)/2 {
		return arr[:mid], arr[mid+1:]
	} else {
		return arr[mid+1:], arr[:mid]
	}
}

func seq(arr []int) {
	var prefix []int

	for len(arr) >= minQsort {
		prefix, arr = partition(arr)

		seq(prefix)
	}

	insertSort(arr)
}

func parHelper(arr []int, wg *sync.WaitGroup) {
	var prefix []int

	for len(arr) >= minQsort {
		prefix, arr = partition(arr)

		if len(prefix) > minParallel && runtime.NumGoroutine() < maxGoroutines {
			wg.Add(1)

			go func(d []int) {
				parHelper(d, wg)

				wg.Done()
			}(prefix)
		} else {
			parHelper(prefix, wg)
		}
	}

	insertSort(arr)
}

func par(arr []int) {
	var wg sync.WaitGroup

	parHelper(arr, &wg)

	wg.Wait()
}

func main() {
	arr := rand.Perm(10000000)

	seq(arr)

	fmt.Println(sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] }))

	arr = rand.Perm(10000000)

	par(arr)

	fmt.Println(sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] }))
}
