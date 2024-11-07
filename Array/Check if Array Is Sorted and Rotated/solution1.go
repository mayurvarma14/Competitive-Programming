package main

import "fmt"

func check(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	pivot := false
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i-1] <= nums[i] {
			continue
		} else if pivot == false {
			pivot = true
		} else {
			return false
		}
	}
	if pivot == true && !(nums[0] >= nums[len(nums)-1]) {
		return false
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{2, 1, 3, 4}
	fmt.Println(check(nums))
}
