package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	var result []string

	for _, w := range words {
		for _, j := range words {
			if w == j {
				continue
			}

			if strings.Contains(j, w) {
				result = append(result, w)
				break
			}
		}
	}

	return result
}

func main() {

	x := stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"})
	fmt.Println(x)
}
