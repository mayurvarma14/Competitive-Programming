package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	single := nums[0]
	if len(nums) == 1 {
		return single
	}
	for i := 1; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}

func main() {

	nums := []int{1}

	fmt.Println(singleNumber(nums))

}
