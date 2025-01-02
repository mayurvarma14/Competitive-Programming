package main

import "fmt"

func reverseWords(s string) string {
	sentence := []rune(s)
	reverse(sentence, 0, len(sentence)-1)
	start := 0
	end := 0
	for i, c := range sentence {
		if c == ' ' || i == len(sentence)-1 {
			end = i - 1
			if i == len(sentence)-1 {
				end = i
			}
			reverse(sentence, start, end)
			start = i + 1
		}
	}

	result := make([]rune, 0, len(sentence))
	i := 0
	for i < len(sentence) {
		for i < len(sentence) && sentence[i] == ' ' {
			i++
		}
		start := i
		for i < len(sentence) && sentence[i] != ' ' {
			i++
		}
		if start < i {
			if len(result) > 0 {
				result = append(result, ' ')
			}
			result = append(result, sentence[start:i]...)
		}
	}
	return string(result)
}
func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	fmt.Println(reverseWords("hello world"))
}
