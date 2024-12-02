package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	odd := []int{}
	even := []int{}

	for i := range nums {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return append(even, odd...)
}

func main() {
	arr := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(arr))
}
