package main

import "fmt"

func IsValid(s string) bool {
	var result []rune

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			result = append(result, ch)
		} else if ch == ')' || ch == '}' || ch == ']' {
			if len(result) == 0 {
				return false
			}
			top := result[len(result)-1]

			if (ch == ')' && top != '(') || (ch == '}' && top != '{') || (ch == ']' && top != '[') {
				return false
			}
			result = result[:len(result)-1]
		}
	}
	return len(result) == 0
}

func main() {
	str := "([])"
	str2 := "(][)"
	fmt.Println(IsValid(str))
	fmt.Println(IsValid(str2))
}