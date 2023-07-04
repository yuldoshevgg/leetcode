package main

import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	i := 1

	for k > 0 {
		if len(num) >= i {
			k += num[len(num)-i]
			num[len(num)-i] = k % 10
		} else {
			num = append([]int{k % 10}, num...)
		}

		i++
		k = k / 10
	}

	return num
}

func main() {

	x := addToArrayForm([]int{1, 2, 0, 0}, 34)
	fmt.Println(x)
}
