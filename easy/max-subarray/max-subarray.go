package main

import "fmt"

//TODO: find the index of start and end of the max subarray
func maxSubArray(nums []int) int {
	max := -9999999
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, 7, 4}
	nums := []int{-2, 1, -3, -4}
	fmt.Println(maxSubArray(nums))
}
