//Perfect square

package main

import "fmt"

func PerfectSquare(num int)bool{
	if num == 1{
	 return	true
	}
	low,high := 1,num

	for low <= high{
		mid := low + (high - low)/2
		square := mid*mid
		if square == num{
			return true
		}else if square < num{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return false
}

func main() {
	num1 := 16
	num2 := 14

	fmt.Println("checking the number 16 is a perfect square: ",PerfectSquare(num1))
	fmt.Println("checking the number 14 is a perfect square: ",PerfectSquare(num2))
}