package main

import (
	"fmt"
	"math"
)

func solution(num int) bool {
	var d = []int{1}

	if num%2 != 0 {
		return false
	}

	for i := 2; i <= int(math.Floor(float64(num/2))); i++ {
		if num%i == 0 {
			d = append(d, i)
		}
	}

	fmt.Println(d)

	for _, i := range d {
		num -= i
	}

	fmt.Println(num)

	return num == 0
}

func main() {

	x := solution(99999999)
	fmt.Println(x)
}
