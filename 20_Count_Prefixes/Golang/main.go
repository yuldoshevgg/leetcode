package main

import (
	"fmt"
)

func countPrefixes(words []string, s string) int {
	var result = 0

	for _, i := range words {
		if len(s) >= len(i) && s[0:len(i)] == i {
			result++
		}
	}

	return result
}

func main() {
	x := countPrefixes([]string{"a", "a"}, "aa")
	fmt.Println(x)
}
