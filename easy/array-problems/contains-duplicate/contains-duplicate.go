package main

import "fmt"

func containsDuplicate(nums []int) bool {
    m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = 1
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,2,2,3,4}))
}