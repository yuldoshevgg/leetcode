package main

import (
	"fmt"
	"strings"
)

func solution(strs []string) string {
	var result = ""

	fmt.Println(string(strs[0][0]))

	for _, i := range strs {
		s := strings.Split(i, "")

		for _, j := range s {
			fmt.Println(j)
		}
	}

	return result
}

func main() {

	x := solution([]string{"flower", "flow", "flight"})
	fmt.Println(x)
}
