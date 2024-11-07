package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
