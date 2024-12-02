package main

import "fmt"

func majorityElement(arr []int) int {
	candidate := 0
	count := 0

	for _, num := range arr {
		if count == 0 {
			candidate = num
			count = 1
		} else if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	arr := []int{3, 2, 3}
	fmt.Println(majorityElement(arr))
}
