package main

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	if target == nil && arr == nil {
		return true
	}

	if target == nil  ||  arr == nil {
		return false
	}

	if len(target)!= len(arr){
		return false
	}

	sort.Ints(target)
	sort.Ints(arr)

	for i := range target{
		if target[i] != arr[i]{
			return false
		}
	}

	return true
}

func main() {
	target := []int{3, 7, 9}
	arr := []int{3,7,11}

	fmt.Println(canBeEqual(target, arr))
}
