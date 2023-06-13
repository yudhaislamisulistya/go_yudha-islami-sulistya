package main

import "fmt"

// Sample Test Case
// Input : [1, 2, 3, 4, 6], target=6
// Output : [1, 3]

func PairSum(arr []int, target int) []int {
	// solve with linear complexity O(n), not O(n^2)
	var result []int
	var mapResult = make(map[int]int)
	for index, value := range arr {
		mapResult[value] = index
	}
	for index, value := range arr {
		if _, ok := mapResult[target-value]; ok && index != mapResult[target-value] {
			result = append(result, index, mapResult[target-value])
			break
		}
	}
	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
