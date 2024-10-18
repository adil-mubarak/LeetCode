// longest word in dictionary
package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string)string{
	sort.Strings(words)

	var result string

	buildWord := make(map[string]bool)
	buildWord[""]=true

	for _,word := range words{
		if buildWord[word[:len(word)-1]]{
			buildWord[word] = true

			if len(word) > len(result){
				result = word
			}
		}
	}
	return result

}

func main() {
	word := []string{"w","wo","wor","worl","world"}
	word2 := []string{"a","ap","app","appl","banana","apple"}

	fmt.Println(longestWord(word))
	fmt.Println(longestWord(word2))
}