package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	n := len(nums)
// 	max := nums[n-1]
// 	for i := 0; i < n-1; i++ {
// 		if nums[i] == max {
// 			return i + 1
// 		}
// 		for nums[i] == nums[i+1] {
// 			shift(nums, i+1)
// 		}
// 	}
// 	return n
// }
// func shift(nums []int, k int) {
// 	for i := k - 1; i < len(nums)-1; i++ {
// 		nums[i] = nums[i+1]
// 	}
// }

func removeDuplicates(nums []int) int {
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r := removeDuplicates(nums)
	fmt.Println("result: ", r)
	// shift(nums, 1)
	fmt.Println(nums)
}
