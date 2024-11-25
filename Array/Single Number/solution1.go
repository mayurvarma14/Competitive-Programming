package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i = i + 2 {

		if nums[i] == nums[i+1] {
			continue
		} else {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func main() {

	nums := []int{4, 1, 2, 1, 2, 5, 5}

	fmt.Println(singleNumber(nums))

}
