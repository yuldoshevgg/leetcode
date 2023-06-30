package main

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	var result int64 = 0

	for i := 0; i < len(nums); i++ {
		if len(nums) > 1 {
			value, _ := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
			result += int64(value)
			nums = nums[1:]
			nums = nums[:len(nums)-1]
			i--
		} else {
			fmt.Println(nums[i])
			result += int64(nums[i])
		}
	}

	return result
}

func main() {

	x := findTheArrayConcVal([]int{5, 14, 13, 8, 12})
	fmt.Println(x)
}
