package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	var max = 0

	for _, i := range sentences {
		s := strings.Split(i, " ")
		if len(s) > max {
			max = len(s)
		}
	}

	return max
}

func main() {

	x := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})
	fmt.Println(x)
}
