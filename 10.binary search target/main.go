package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{-1, 0, 2, 3, 9, 10}
	fmt.Println(BinarySearch(arr,9))
}