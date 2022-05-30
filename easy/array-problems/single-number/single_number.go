package main

import "fmt"

// func singleNumber(nums []int) int {
// 	numsMap := make(map[int]int)
// 	for _, num := range nums {
// 		numsMap[num]++
// 	}
// 	for key, value := range numsMap {
// 		if value == 1 {
// 			return key
// 		}
// 	}
// 	return 0
// }

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}
func main() {
	// fmt.Println(singleNumber([]int{2, 2, 1, 4, 4}))
	// 5 = 0101, 2 = 0010, 5 ^ 2 = 0101 ^ 0010 = 0111
	//
	// fmt.Println(5 ^ 1)
	fmt.Println(1 ^ 4 ^ 4)
	fmt.Println(4 ^ 4 ^ 1)
	fmt.Println(4 ^ 1 ^ 4)

}
