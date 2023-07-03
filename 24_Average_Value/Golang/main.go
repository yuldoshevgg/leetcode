package main

import (
	"fmt"
	"math"
)

func averageValue(nums []int) int {
	var (
		sum   = 0
		count = 0
	)

	for _, i := range nums {
		if i%6 == 0 {
			sum += i
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return int(math.Round(float64(sum / count)))
}

func main() {

	x := averageValue([]int{1, 3, 6, 10, 12, 15})
	fmt.Println(x)
}
