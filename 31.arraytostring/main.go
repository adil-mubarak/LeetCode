package main

import (
	"fmt"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	n := len(nums)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if strconv.Itoa(nums[j])+strconv.Itoa(nums[j+1]) < strconv.Itoa(nums[j+1])+strconv.Itoa(nums[j]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	var result strings.Builder
	for _, num := range nums {
		result.WriteString(strconv.Itoa(num))
	}

	if result.String()[0] == '0' {
		return "0"
	}
	return result.String()
}

func main() {

	arr := []int{10, 2}
	result := largestNumber(arr)
	fmt.Println(result)

}
