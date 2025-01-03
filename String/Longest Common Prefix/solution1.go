package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix_length := -1
	prefix_compare := strs[0]
	for i, c := range prefix_compare {
		for _, str := range strs {
			if i == len(str) || rune(str[i]) != c {
				if prefix_length == -1 {
					return ""
				}
				return prefix_compare[:prefix_length+1]
			}
		}
		prefix_length = i
	}
	return prefix_compare[:prefix_length+1]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	// strs := []string{"aflow", "bflow"}
	fmt.Println(longestCommonPrefix(strs))
}
