package main

import (
	"fmt"
	"strconv"
)

func baseballGame(ops []string) int {
	var record = []int{}
	var total = 0

	for i := 0; i < len(ops); i++ {
		if _, err := strconv.Atoi(ops[i]); err == nil {
			num, _ := strconv.Atoi(ops[i])
			record = append(record, num)
		} else if ops[i] == "C" {
			record = record[0 : len(record)-1]
		} else if ops[i] == "D" {
			record = append(record, record[len(record)-1]*2)
		} else if ops[i] == "+" {
			record = append(record, record[len(record)-1]+record[len(record)-2])
		}
	}

	if len(record) > 0 {
		for i := 0; i < len(record); i++ {
			total += record[i]
		}
	}

	return total
}

func main() {

	x := baseballGame([]string{"1", "C"})
	fmt.Println(x)
}
