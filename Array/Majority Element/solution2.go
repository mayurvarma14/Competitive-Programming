package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	count := 0
	majorityElement := nums[0]
	for _, num := range nums {
		if count == 0 {
			majorityElement = num
			count++
		} else if majorityElement == num {
			count++
		} else {
			count--
		}

	}
	return majorityElement

}

func main() {
	nums := []int{6, 5, 5}
	fmt.Println(majorityElement(nums))
}
