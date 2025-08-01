package main

import "fmt"

func missingNumber(nums []int) int {
	total := len(nums) * (len(nums) + 1) / 2
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return total - sum
}

func main() {

	r := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	fmt.Println(r)
}
