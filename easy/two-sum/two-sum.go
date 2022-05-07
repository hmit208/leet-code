package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{i, idx}
		}
		m[nums[i]] = i

	}
	return []int{0, 0}
}

func main() {
	// nums := []int{3, 2, 4}
	// fmt.Println(twoSum(nums, 6))
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(a))
	fmt.Println(len(a))
	a = append(a, 9)
	var b [2]int
	copy(b[0:], a)
	fmt.Println(b)
}

func test(s []int) {
	fmt.Println("input: ", &s[0])
}
