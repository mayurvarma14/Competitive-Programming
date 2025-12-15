package main

import (
	"fmt"
)

func WaveArray(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := range arr {
		if i%2 == 0 {
			continue
		}
		if arr[i] > arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}

		if i+1 < l && arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

}

func main() {
	nums := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray(nums)
	fmt.Println(nums)
}
