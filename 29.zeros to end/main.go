package main

import "fmt"

func moveZerosToEnd(arr []int) []int {
	noZerindex := 0
	for _, num := range arr {
		if num != 0 {
			arr[noZerindex] = num
			noZerindex++
		}
	}
	for noZerindex < len(arr) {
		arr[noZerindex] = 0
		noZerindex++
	}
	return arr
}

func main() {
	input := []int{1, 0, 2, 0, 3, 0, 4}
	output := moveZerosToEnd(input)
	fmt.Println(output)
}
