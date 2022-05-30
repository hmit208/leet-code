package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{i, idx}
		}
		m[nums[i]] = i

	}
	return []int{0, 0}
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}
