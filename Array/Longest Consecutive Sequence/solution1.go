package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	longest := 0
	count := 0
	for num := range m {
		if _, ok := m[num-1]; !ok {
			start := num
			for {
				if _, ok := m[start]; ok {
					start++
					count++
				} else {
					longest = max(longest, count)
					count = 0
					break
				}

			}
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	// nums := []int{0, -1}
	fmt.Println(longestConsecutive(nums))
}
