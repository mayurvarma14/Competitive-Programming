package main

import "fmt"

func sortColors(nums []int) {
	var count0, count1, count2 int
	for _, num := range nums {
		if num == 0 {
			count0++
		} else if num == 1 {
			count1++
		} else {
			count2++
		}
	}

	for i := 0; i < count0; i++ {
		nums[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		nums[i] = 1
	}
	for i := count0 + count1; i < len(nums); i++ {
		nums[i] = 2
	}

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums) // Output: [0 0 1 1 2 2]
}
