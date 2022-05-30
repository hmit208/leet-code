package main

import "fmt"

/* Brute force O(n2) solution

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	profit := 0
// 	for i := len(prices) - 1; i > 0; i-- {
// 		for j := i - 1; j >= 0; j-- {
// 			profit = max(profit, prices[i]-prices[j])
// 		}
// 	}
// 	return profit
// }

*/

/* Kadane's algorithm O(n) solution
func maxProfit(prices []int) int {
	maxCur := 0
	maxSoFar := 0
	for i := 1; i < len(prices); i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(0, maxCur)
		maxSoFar = max(maxCur, maxSoFar)
	}
	return maxSoFar
}
*/

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		} else {
			minPrice = min(minPrice, prices[i])
		}
	}
	return maxProfit
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
