package main

import (
	"fmt"
	"strings"
)

func LowerCase(s string) string {
	lower := strings.ToLower(s)
	return lower
}

func main() {
	str := "Hello World"
	fmt.Println(LowerCase(str))
}