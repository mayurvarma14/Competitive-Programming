package main

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	pivot := length - k
	reverse(nums, 0, pivot-1)
	reverse(nums, pivot, length-1)
	reverse(nums, 0, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
