package main

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var (
		result []bool
		max    = 0
	)

	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	for _, c := range candies {
		if c+extraCandies >= max {
			result = append(result, true)
			continue
		}

		result = append(result, false)
	}

	return result
}

func main() {

	x := kidsWithCandies([]int{12, 1, 12}, 10)
	fmt.Println(x)
}
