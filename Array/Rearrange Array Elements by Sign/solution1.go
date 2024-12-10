package main

import (
	"fmt"
)

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	pos := 0
	neg := 1
	for _, num := range nums {
		if num >= 0 {
			res[pos] = num
			pos += 2
		} else {
			res[neg] = num
			neg += 2
		}
	}
	return res

}

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}
	fmt.Println(rearrangeArray(nums))
}
