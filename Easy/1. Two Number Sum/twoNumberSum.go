package main

import (
	"fmt"
	"sort"
)

// O(n^2) time | O(1) space
func solution1(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

// O(n) time | O(n) space
func solution2(array []int, target int) []int {
	nums := map[int]bool{}
	for _, x := range array {
		y := target - x
		if _, found := nums[y]; found {
			return []int{y, x}
		}
		nums[x] = true
	}
	return []int{}
}

// O(nlog(n)) time | O(1) space
func solution3(array []int, target int) []int {
	sort.Ints(array)
	l, r := 0, len(array)-1
	for l < r {
		tar := array[l] + array[r]
		if tar == target {
			return []int{array[l], array[r]}
		} else if tar < target {
			l += 1
		} else {
			r -= 1
		}
	}
	return []int{}
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	var targetSum int = 10

	fmt.Println("Solution 1: ", solution1(array, targetSum))
	fmt.Println("Solution 2: ", solution2(array, targetSum))
	fmt.Println("Solution 3: ", solution3(array, targetSum))
}
