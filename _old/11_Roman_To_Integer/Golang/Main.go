package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	var (
		result = 0
		romans = map[string]int{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		}
	)

	if strings.Contains(s, "IV") {
		result += 4
		s = strings.Replace(s, "IV", "", -1)
	}

	if strings.Contains(s, "IX") {
		result += 9
		s = strings.Replace(s, "IX", "", -1)
	}

	if strings.Contains(s, "XL") {
		result += 40
		s = strings.Replace(s, "XL", "", -1)
	}

	if strings.Contains(s, "XC") {
		result += 90
		s = strings.Replace(s, "XC", "", -1)
	}

	if strings.Contains(s, "CD") {
		result += 400
		s = strings.Replace(s, "CD", "", -1)
	}

	if strings.Contains(s, "CM") {
		result += 900
		s = strings.Replace(s, "CM", "", -1)
	}

	for i, r := range romans {
		if strings.Contains(s, i) {
			result += strings.Count(s, i) * r
		}
	}

	return result
}

func main() {

	x := romanToInt("XXXIV")
	fmt.Println(x)
}
