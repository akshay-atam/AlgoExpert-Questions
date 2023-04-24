package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(n) space
func solution1(array []int) []int {
	sortedArray := make([]int, len(array))

	for i, val := range array {
		sortedArray[i] = val * val
	}

	sort.Ints(sortedArray)
	return sortedArray
}

// O(n) time | O(n) space
func solution2(array []int) []int {
	sortedArray := make([]int, len(array))

	smaller := 0
	larger := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		smallerValue := array[smaller]
		largerValue := array[larger]

		if abs(smallerValue) > abs(largerValue) {
			sortedArray[i] = smallerValue * smallerValue
			smaller++
		} else {
			sortedArray[i] = largerValue * largerValue
			larger--
		}
	}
	return sortedArray
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	array := []int{1, 2, 3, 5, 6, 8, 9}
	fmt.Println("Solution 1: ", solution1(array))
	fmt.Println("Solution 2: ", solution2(array))

}
