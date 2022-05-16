package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	prev := generateRow(2, nil)
	for i := 3; i <= numRows; i++ {
		prev = generateRow(i, prev)
		res = append(res, prev)
	}
	return res
}
func generateRow(numElements int, prev []int) []int {
	if numElements == 1 {
		return []int{1}
	}
	if numElements == 2 {
		return []int{1, 1}
	}
	res := []int{1}
	for i := 1; i < numElements-1; i++ {
		res = append(res, prev[i-1]+prev[i])
	}
	res = append(res, 1)
	return res
}

func main() {
	fmt.Println(generate(5))
}
