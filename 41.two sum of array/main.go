package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}

func main() {
	arr := []int{2, 3, 4}
	fmt.Println(twoSum(arr,6))

}