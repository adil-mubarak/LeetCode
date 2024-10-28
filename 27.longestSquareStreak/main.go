package main

import (
	"fmt"
	"sort"
)

func LongestSquareStreak(nums []int) int {
	sort.Ints(nums)
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	longStreak := 0

	for _, num := range nums {
		streak := 0
		for numMap[num]{
			streak++
			num *= num
		}
		if longStreak < streak{
			longStreak = streak
		}
	}
	if longStreak <= 1 {
		return -1
	}
	return longStreak

}

func main() {
	nums := []int{4, 3, 6, 16, 8, 2}
	fmt.Println(LongestSquareStreak(nums))
}
