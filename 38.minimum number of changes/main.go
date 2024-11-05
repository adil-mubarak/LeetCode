package main

import "fmt"

func minChanges(str string) int {
	count := 0
	for i := 0; i < len(str); i+=2 {
		if str[i] != str[i+1] {
			count++
		}
	}
	return count
}

func main() {
	str := "1010"
	fmt.Println(minChanges(str))
}
