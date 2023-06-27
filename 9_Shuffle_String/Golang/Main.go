package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	cap := len(s)

	result := make([]byte, cap, cap)
	for i := 0; i < len(indices); i++ {
		result[indices[i]] = s[i]
	}

	return string(result)
}

func main() {

	x := restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})
	fmt.Println(x)
}
