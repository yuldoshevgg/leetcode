package main

import (
	"fmt"
	"strings"
)

func vowalString(words []string, left, right int) int {
	var vowals = "aeiou"
	var count = 0

	for i := left; i <= right; i++ {
		spString := strings.Split(words[i], "")

		if strings.Contains(vowals, spString[0]) && strings.Contains(vowals, spString[len(spString)-1]) {
			count++
		}
	}

	return count
}

func main() {

	x := vowalString([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4)
	fmt.Println(x)
}
