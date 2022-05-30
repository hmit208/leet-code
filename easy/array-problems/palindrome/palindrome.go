package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 10 {
        return x >= 0
    }
	return x == getReverse(x)
}

func getReverse(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}

func main() {
	fmt.Println("12321: ", isPalindrome(12321))
	fmt.Println("123", isPalindrome(123))
	fmt.Println("123321", isPalindrome(123321))
	fmt.Println("12531", isPalindrome(12531))
	fmt.Println("14741", isPalindrome(14741))
	fmt.Println("1441", isPalindrome(1441))
	fmt.Println("290092", isPalindrome(290092))
	fmt.Println("28082", isPalindrome(28082))
}
