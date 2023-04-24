package main

import "fmt"

// O(n) time | O(1) space
func solution1(array []int, sequence []int) bool {
	index := 0
	for _, value := range array {
		if index == len(sequence) {
			break
		}
		if sequence[index] == value {
			index += 1
		}
	}
	return index == len(sequence)
}

// O(n) time | O(1) space
func solution2(array []int, sequence []int) bool {
	arrIndex := 0
	seqIndex := 0
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex += 1
		}
		arrIndex += 1
	}
	return seqIndex == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	fmt.Println("Solution 1: ", solution1(array, sequence))
	fmt.Println("Solution 2: ", solution2(array, sequence))
}
