package main

import "fmt"

func rotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s == goal {
			return true
		}
		s = s[1:] + string(s[0])
	}
	return false
}

func main() {
	s := "abcde"
	goal := "bcdea"

	fmt.Println(rotateString(s, goal))
}
