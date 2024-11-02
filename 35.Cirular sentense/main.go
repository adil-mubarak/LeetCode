package main

import (
	"fmt"
	"strings"
)

func isCircularSentence(sentence string) bool {
	result := strings.Split(sentence, " ")
	for i := 0; i < len(result); i++ {
		lastChar := result[i][len(result[i])-1]
		firstChar := result[(i+1)%len(result)][0]

		if lastChar != firstChar{
			return false
		}
	}
	return true
}

func main() {
	str := "MuFoevIXCZzrpXeRmTssj lYSW U jM"
	fmt.Println(isCircularSentence(str))

}
