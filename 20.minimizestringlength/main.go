package main

import "fmt"

func minimizestringlenght(s string) int {
	seen := make(map[rune]bool)

	result := []rune{}

	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}
	ans := string(result)
	return len(ans)
}

func main() {
	str := "aaabc"
	fmt.Println(minimizestringlenght(str))
}