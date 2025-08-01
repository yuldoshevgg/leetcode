package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	var (
		existWords []string
		firstRow   = "qwertyuiop"
		secondRow  = "asdfghjkl"
		thirdRow   = "zxcvbnm"
	)

	for i := 0; i < len(words); i++ {
		var ac, bc, cc = 0, 0, 0

		for j := 0; j < len(words[i]); j++ {
			letter := strings.ToLower(string(words[i][j]))

			if strings.Contains(firstRow, letter) {
				ac++
			} else if strings.Contains(secondRow, letter) {
				bc++
			} else if strings.Contains(thirdRow, letter) {
				cc++
			}
		}

		if ac == len(words[i]) || bc == len(words[i]) || cc == len(words[i]) {
			existWords = append(existWords, words[i])
		}
	}

	return existWords
}

func main() {

	x := findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	fmt.Println(x)
}
