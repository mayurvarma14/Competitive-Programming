package main

import "fmt"

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	total_sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		total_sum += nums[i]
		count += m[total_sum-k]
		m[total_sum] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0))
}
