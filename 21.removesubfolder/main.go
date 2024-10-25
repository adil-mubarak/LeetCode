package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeFolderSub(folder []string) []string {
	sort.Strings(folder)
	result := []string{}
	
	prev := folder[0]
	
	result = append(result, prev)
	for i := 1; i < len(folder); i++{
		if !strings.HasPrefix(folder[i], prev+"/"){
			prev = folder[i]
			result = append(result, prev)
		}

	}
	return result
}

func main() {
	folder := []string{"/a" ,"/a/b" ,"/c/d" ,"/c/d/e" ,"/c/f"}
	result := removeFolderSub(folder)
	fmt.Println(result)
}
