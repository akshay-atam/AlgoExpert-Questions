package main

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	currentMaxChange := 0

	for _, coin := range coins {
		if coin > currentMaxChange+1 {
			return currentMaxChange + 1
		}
		currentMaxChange += coin
	}
	return currentMaxChange + 1
}

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println("Minimum amount of change that cannot be created: ", NonConstructibleChange(coins))
}
