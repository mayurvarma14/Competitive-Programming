package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const threshold = 5000 // Switch to sequential for small subarrays

// Sequential Merge Sort
func sequentialMergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	sequentialMergeSort(arr, start, mid)
	sequentialMergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

// Parallel Merge Sort
func parallelMergeSort(arr []int, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	if start >= end {
		return
	}

	// Switch to sequential for small subarrays
	if end-start <= threshold {
		sequentialMergeSort(arr, start, end)
		return
	}

	mid := (start + end) / 2

	var childWg sync.WaitGroup
	childWg.Add(2)

	go parallelMergeSort(arr, start, mid, &childWg)
	go parallelMergeSort(arr, mid+1, end, &childWg)

	childWg.Wait()
	merge(arr, start, mid, end)
}

// Merge function (common for both sequential and parallel)
func merge(arr []int, start, mid, end int) {
	temp := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= end {
		temp[k] = arr[j]
		j++
		k++
	}

	for i := 0; i < len(temp); i++ {
		arr[start+i] = temp[i]
	}
}

// Benchmarking function
func benchmarkMergeSort(arr []int, sortFunc func([]int, int, int), name string) {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	start := time.Now()
	sortFunc(arrCopy, 0, len(arrCopy)-1)
	elapsed := time.Since(start)

	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	n := 1000000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	// Benchmark Sequential Merge Sort
	benchmarkMergeSort(arr, sequentialMergeSort, "Sequential Merge Sort")

	// Benchmark Parallel Merge Sort
	var wg sync.WaitGroup
	wg.Add(1)
	benchmarkMergeSort(arr, func(arr []int, start, end int) {
		parallelMergeSort(arr, start, end, &wg)
		wg.Wait()
	}, "Parallel Merge Sort")
}
