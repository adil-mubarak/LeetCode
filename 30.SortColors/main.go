package main

import "fmt"

func SortColor(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	num := []int{2, 0, 2, 1, 1, 0}
	SortColor(num)
	fmt.Println(num)
}