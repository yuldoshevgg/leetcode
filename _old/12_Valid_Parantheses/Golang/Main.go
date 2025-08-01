package main

import (
	"fmt"
)

func isValid(s string) bool {
	var closingToOpening = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	str := []byte(s)
	stack := make([]byte, 0)

	for _, ch := range str {
		matchingOpening, isClosing := closingToOpening[ch]

		if !isClosing {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastOpening := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if lastOpening != matchingOpening {
			return false
		}
	}

	return len(stack) == 0
}

func main() {

	x := isValid("]")
	fmt.Println(x)
}
