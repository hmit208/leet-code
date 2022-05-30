package main

import "fmt"

// func plusOne(digits []int) []int {
// 	n := len(digits)
// 	if n == 0 {
// 		return digits
// 	}
// 	result := []int{}
// 	var originalNumber int
// 	for _, v := range digits {
// 		originalNumber = originalNumber*10 + v
// 	}
// 	originalNumber++
// 	for originalNumber > 0 {
// 		result = append([]int{originalNumber % 10}, result...)
// 		originalNumber /= 10
// 	}
// 	return result
// }

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}
func main() {
	fmt.Println(plusOne([]int{1, 0, 2, 3}))
}
