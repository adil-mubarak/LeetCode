package main

import "fmt"

func romanToInt(s string) int {
    romanToIntMap := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }

    total := 0
    n := len(s)

    for i := 0;i < n;i++{
        currentValue := romanToIntMap[s[i]]

        if i+1 < n && currentValue < romanToIntMap[s[i+1]]{
            total -= currentValue
        }else{
            total += currentValue
        }
    }
    return total
}


func main() {
	fmt.Println(romanToInt("III")) // 3
}