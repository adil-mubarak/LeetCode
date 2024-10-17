// count prefix and suffix pairs
package main

import (
	"fmt"
	"strings"
)

func countPrefixSuffixPairs(word []string) int {
	count := 0
	n := len(word)
	for i := 0; i < n-1; i++ {
		for j := i+ 1; j < n; j++ {
			if isPrefixAndSuffix(word[i], word[j]) {
				count++
			}
		}
	}
	return count
}


func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}

func main() {
	wrd := []string{"a", "aba", "ababa", "aa"}
	fmt.Println(wrd)
	fmt.Println(countPrefixSuffixPairs(wrd))
}
