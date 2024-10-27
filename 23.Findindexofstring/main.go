package main

import "fmt"

func strStr(haystack, needlse string) int {
	for i := 0; i <= len(haystack)-len(needlse); i++ {
		if haystack[i:len(needlse)+i] == needlse {
			return i
		}
	}
	return -1
}

func main() {
	str1 := "sadbutsas"
	str2 := "sad"

	fmt.Println(strStr(str1,str2))
}