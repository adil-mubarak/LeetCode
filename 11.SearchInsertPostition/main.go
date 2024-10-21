//Search insert position

package main

import "fmt"

func SearchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(SearchInsert(arr, 5))
}
