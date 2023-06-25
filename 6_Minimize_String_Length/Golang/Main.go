package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimizedStringLength(s string) int {
	var (
		length  = 1
		splited = strings.Split(s, "")
	)

	sort.Slice(splited, func(i, j int) bool {
		return splited[i] < splited[j]
	})

	var currentChar = splited[0]

	for i := 0; i < len(splited); i++ {
		if splited[i] != currentChar {
			currentChar = splited[i]
			length++
		}
	}

	return length
}

func main() {

	x := minimizedStringLength("cbbd")
	fmt.Println(x)
}
