package main

import (
	"fmt"
	"math"
)

func searchInsert(nums []int, target int) int {
	var (
		start = 0
		end   = len(nums) - 1
	)

	for start <= end {
		var mid = int(math.Floor(float64((start + end) / 2)))

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		}

		if nums[mid] < target {
			start = mid + 1
		}
	}

	return start
}

func main() {

	x := searchInsert([]int{1}, 1)
	fmt.Println(x)
}
