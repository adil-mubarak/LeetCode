package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s, t string) bool {
	x := strings.Split(s, "")
	y := strings.Split(t, "")
	sort.Strings(x)
	sort.Strings(y)

	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	x := make(map[rune]int)
	y := make(map[rune]int)

	for _, ch := range s {
		x[ch]++
	}
	for _, ch := range t {
		y[ch]++
	}

	for key,val := range x{
		if y[key] != val{
			return false
		}
	}
	return true
}

func main() {
	s := "car"
	t := "rat"
	fmt.Println(isAnagram(s, t))
	fmt.Println((IsAnagram(s,t)))
}
