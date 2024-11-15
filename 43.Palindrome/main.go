package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	y := strconv.Itoa(x)
	runes := []rune(y)
	for i,j := 0,len(runes)-1;i < j;i,j = i+1,j-1{
		if runes[i] != runes[j]{
			return false
		}
	}
	return true
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))

}
