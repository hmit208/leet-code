package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
   if len(nums) == 0 {
	   return nil
   } 
   return &TreeNode{
	   Val: nums[len(nums)/2],
	   Left: sortedArrayToBST(nums[:len(nums)/2]),
	   Right: sortedArrayToBST(nums[len(nums)/2+1:]),
   }
}
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	a  := sortedArrayToBST(nums)
	fmt.Println(a)
}