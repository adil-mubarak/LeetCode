package main

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	runes := []rune(word)
	var result strings.Builder
	for i := 0; i < len(runes); {
		count := 1
		for i+1 < len(runes) && runes[i] == runes[i+1] && count < 9 {

			count++
			i++
		}
		result.WriteString(fmt.Sprintf("%d%c", count, runes[i]))
		i++
	}
	return result.String()
}

func main() {
	str := "aaabbcc"

	fmt.Println(compressedString(str))
}
