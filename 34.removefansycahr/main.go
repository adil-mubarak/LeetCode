package main

import "fmt"

func removefancy(s string) string {
	runes := []rune(s)
	result := []rune{runes[0]}

	count := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			result = append(result, runes[i])
		}
	}
	return string(result)
}

func main() {
	str := "leeetcode"
	fmt.Println(removefancy(str))
}
