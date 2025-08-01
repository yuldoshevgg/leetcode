package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 0

	for i := 1; i < n; i++ {
		n2, n1 = n1, n2+n1
	}

	return n1
}

func main() {

	x := fib(12)
	fmt.Println(x)
}
