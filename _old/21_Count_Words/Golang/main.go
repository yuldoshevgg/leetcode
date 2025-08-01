package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	var result = 0

	for _, i := range words {
		if strings.HasPrefix(i, pref) {
			result++
		}
	}

	return result
}

func main() {

	x := prefixCount([]string{"leetcode", "win", "loops", "success"}, "code")
	fmt.Println(x)
}
