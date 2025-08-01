package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	sentence := strings.Split(s, " ")

	return strings.Join(sentence[0:k], " ")
}

func main() {

	x := truncateSentence("What is the solution to this problem", 4)
	fmt.Println(x)
}
