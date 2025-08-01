package main

import (
	"fmt"
	"strings"
)

func solution(address string) string {

	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {

	x := solution("1.1.1.1")
	fmt.Println(x)
}
