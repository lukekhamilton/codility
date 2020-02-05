package main

import (
	"fmt"
	"strconv"
)

// BinaryGag returns the longest binary gap. ie the amount of 0 between the 1 in he binary representation of the number given.
func BinaryGap(N int) int {
	binary := strconv.FormatInt(int64(N), 2)
	fmt.Println("binary: ", binary)
	largestGap := 0
	currentGap := 0
	startCount := false
	for _, b := range binary {
		if string(b) == "1" {
			startCount = true
			if largestGap < currentGap {
				largestGap = currentGap
			}
			currentGap = 0
			continue
		}
		if startCount {
			currentGap++
		}
	}
	return largestGap
}
