package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	var res int64
	for i := start; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		res = res*10 + int64(s[i]-'0')

		if int64(sign)*res > math.MaxInt32 {
			return math.MaxInt32
		}
		if int64(sign)*res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return int(int64(sign) * res)
}

func main() {
	// fmt.Println(myAtoi(" -04w2"))
	// fmt.Println(myAtoi("0-1"))
	// fmt.Println(myAtoi("   -042"))
	// fmt.Println(myAtoi("-91283472332"))
	// fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("20000000000000000000"))

}
