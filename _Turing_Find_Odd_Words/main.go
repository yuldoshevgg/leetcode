package main

import (
	"fmt"
	"sort"
	"strings"
)

func findOddWords(set1, set2 string) []string {
	var (
		result []string
		setX   = set1 + " " + set2
		arr    = strings.Split(setX, " ")
	)

	sort.Slice(arr, func(i, j int) bool {
		return arr[j] > arr[i]
	})

	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
		for j := i + 1; j < len(arr); j++ {
			if arr[i] != arr[j] {
				result = append(result, arr[i])
				break
			} else if arr[i] == arr[j] {
				fmt.Println(arr[i], arr[j])
				break
			}
		}
	}

	return result
}

func main() {

	x := findOddWords("turing community is the best", "turing community is the greatest")
	fmt.Println(x)
}
