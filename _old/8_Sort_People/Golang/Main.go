package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	type people struct {
		name   string
		height int
	}

	tmp := make([]people, 0, len(names))

	for i := 0; i < len(names); i++ {
		t := people{
			name:   names[i],
			height: heights[i],
		}
		tmp = append(tmp, t)
	}

	sort.Slice(tmp, func(i, j int) bool { return tmp[i].height > tmp[j].height })

	ans := make([]string, 0, len(names))
	for _, t := range tmp {
		ans = append(ans, t.name)
	}
	return ans
}

func main() {

	x := sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	fmt.Println(x)
}
