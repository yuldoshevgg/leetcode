package main

import (
	"fmt"
)

func solution(x int) bool {
	var (
		x1 = 0
		x2 = x
	)

	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	for x2 > 0 {
		reminder := x2 % 10
		x1 *= 10
		x1 += reminder
		x2 /= 10
	}

	return x1 == x
}

func main() {

	x := solution(121)
	fmt.Println(x)
}
