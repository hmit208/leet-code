package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	if m > len(nums1) {
// 		return
// 	}
// 	i, j := 0, 0
// 	res := []int{}
// 	for i < m && j < n {
// 		if nums1[i] < nums2[j] {
// 			res = append(res, nums1[i])
// 			i++
// 		} else {
// 			res = append(res, nums2[j])
// 			j++
// 		}
// 	}
// 	if i < m {
// 		res = append(res, nums1[i:m]...)
// 	}
// 	if j < n {
// 		res = append(res, nums2[j:]...)
// 	}
// 	copy(nums1, res)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	fmt.Println(i)
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
	// if i >= 0 {
	// 	copy(nums1[:i+1], nums1[:i+1])
	// }
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, len(nums2))
	fmt.Println(nums1)
}
