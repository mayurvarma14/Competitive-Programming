package main

import "fmt"

func printInt(number int) {

	hexMap := "0123456789ABCDEF"
	quotient := number / 16
	remainder := number % 16
	if number < 16 {
		fmt.Printf("%c", hexMap[remainder])
		return
	}
	printInt(quotient)
	fmt.Printf("%c", hexMap[remainder])

}
func main() {
	printInt(123456)
}
