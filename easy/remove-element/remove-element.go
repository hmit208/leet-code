package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
func main() {
	nums := []int{3, 2, 2, 3}
	r := removeElement(nums, 3)
	fmt.Println("nums: ", nums)
	fmt.Println("result: ", r)
}
