package CountingElements

import (
	"fmt"
	"sort"
)

func MissingInteger(A []int) int {
	sort.Ints(A)
	fmt.Println(A)
	min := 1
	for _, e := range A {
		if e == min {
			min++
		}
	}
	return min
}
