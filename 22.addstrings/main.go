package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	result := ""
	carry := 0

	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		var x, y int

		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}

		sum := x + y + carry

		carry = sum / 10
		result = strconv.Itoa(sum%10) + result
	}
	return result

}

func main() {
	num1 := "123"
	num2 := "123"
	result := addStrings(num1, num2)
	fmt.Println(result)
}
