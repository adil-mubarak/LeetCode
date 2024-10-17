//longest commen prefix
package main

import (
	"fmt"
	"sort"
)

func LongestPrefix(str []string)string{
	if len(str) == 0{
		return ""
	}

	sort.Strings(str)
	fmt.Println(str)

	l := len(str)
	first := str[0]

	for i := range first{
		if first[i] != str[l-1][i]{
			return first[:i]
		}
	}
	return first

}

func main() {
	str := []string{"flower", "flow", "flight"}
	str1 := []string{"race","dog","bat"}
	fmt.Println(str)
	fmt.Println(LongestPrefix(str))
	fmt.Println(str1)
	fmt.Println(LongestPrefix(str1))
}