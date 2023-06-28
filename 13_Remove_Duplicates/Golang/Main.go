package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var (
		element = nums[0]
		l       = 1
	)

	for i := 1; i < len(nums); i++ {
		if nums[i] != element {
			nums[l] = nums[i]
			l++
		}

		element = nums[i]
	}

	return l
}

func main() {

	x := removeDuplicates([]int{1, 1, 2})
	fmt.Println(x)
}
