package main

import "fmt"

func maxDepth(s string) int {
	stack := make([]rune, 0)
	var max int
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > max {
			max = len(stack)
		}
	}
	return max
}
func main() {
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
}
