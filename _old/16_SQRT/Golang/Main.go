package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	var left = 0
	var right = x

	for left <= right {
		var mid = int(math.Floor(float64(left+right) / 2))

		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {

	x := mySqrt(4)
	fmt.Println(x)
}
