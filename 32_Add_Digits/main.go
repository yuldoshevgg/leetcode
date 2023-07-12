package main

import (
	"fmt"
	"math"
)

func addDigits(num int) int {
	var result = 0

	if num < 10 {
		return num
	}

	for num != 0 {
		result = result + (num % 10)
		num = int(math.Trunc(float64(num / 10)))
	}

	return addDigits(result)
}

func main() {

	x := addDigits(38)
	fmt.Println(x)
}
