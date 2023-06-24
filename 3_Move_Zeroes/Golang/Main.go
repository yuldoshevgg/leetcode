package main

import "fmt"

func moveZeroes(nums []int) {
	var zeroes = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroes++
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	for i := 0; i < zeroes; i++ {
		nums = append(nums, 0)
	}

	fmt.Println(nums)
}

func main() {

	moveZeroes([]int{0, 0, 1})
}
