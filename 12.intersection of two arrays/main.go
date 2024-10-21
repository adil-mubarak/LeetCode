// intersection of two arrays
package main

import "fmt"

func Intersect(nums1, nums2 []int) []int {
	freq := make(map[int]int)
	var result []int

	for _, num := range nums1 {
		freq[num]++
	}

	for _, num := range nums2 {
		if count, exist := freq[num]; exist && count > 0 {
			result = append(result, num)
			freq[num]--
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(Intersect(nums1, nums2))
}
