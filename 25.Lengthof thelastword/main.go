package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
   words := strings.Fields(s)

   if len(words) == 0{
	return 0
   }

   return len(words[len(words)-1])

}

func main() {
	str := "hello worlds"
	fmt.Println(lengthOfLastWord(str))
	

}
