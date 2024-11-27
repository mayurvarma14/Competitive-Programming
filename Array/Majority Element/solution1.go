package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)

	maxCount := 0
	currentCount := 0
	majorityElement := nums[0]
	currentElement := nums[0]

	for _, num := range nums {
		if num == currentElement {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
				majorityElement = currentElement
			}
			currentElement = num
			currentCount = 1
		}
	}

	if currentCount > maxCount {
		majorityElement = currentElement
	}

	return majorityElement
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
