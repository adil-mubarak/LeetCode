package main

import (
	"fmt"
)

func findTheDifference(s, t string) byte {
	var result byte

	for i := 0; i < len(s) ; i++{
		result ^= s[i]
	}

	for i := 0; i < len(t); i++{
		result ^= t[i]
	}
	return result
}

func main() {
	s := ""
	t := "y"
	result := findTheDifference(s, t)
	ans := string(result)
	fmt.Println(ans)
}
