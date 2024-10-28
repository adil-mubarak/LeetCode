package main

import "fmt"

func Merge(num1 []int, m int, num2 []int, n int) {
	sorted := Sort(num1[:m], num2[:n])
	for i := 0; i < len(sorted); i++ {
		num1[i] = sorted[i]
	}
}
func Sort(num1 []int, num2 []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(num1) && j < len(num2) {
		if num1[i] < num2[j] {
			result = append(result, num1[i])
			i++
		} else {
			result = append(result, num2[j])
			j++
		}
	}
	result = append(result, num1[i:]...)
	result = append(result, num2[j:]...)
	return result
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3

	Merge(nums1, m, nums2, n)
	fmt.Println("Merged array:", nums1)
}