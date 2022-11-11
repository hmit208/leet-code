package main

import (
	"fmt"
)

// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	for _, v := range nums {
// 		m[v]++
// 	}
// 	for k, v := range m {
// 		if v > len(nums)/2 {
// 			return k
// 		}
// 	}
// 	return 0
// }

func majorityElement(nums []int) int {

	res, times := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			times++
		} else {
			times--
			if times < 0 {
				res = nums[i]
				times = 1
			}
		}
	}
	return res
}

func main() {
	// fmt.Println(majorityElement([]int{1, 3, 5, 6, 6, 6, 6}))
	// fmt.Println(selfImplement([]int{1, 3, 5, 6, 6, 6, 6}))
	m := make(map[string]interface{})
	m["a"] = 1
	fmt.Println(m)
}

func selfImplement(nums []int) int {
	can := nums[0]
	count := 0
	for _, v := range nums {
		if v == can {
			count++
		} else if count == 0 {
			can = v
			count++
		} else {
			count--
		}
	}
	return can
}

func m(nums []int) int {
	count := 1
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count++
		} else if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return 0
}
