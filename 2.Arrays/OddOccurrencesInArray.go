package Arrays

// Credit: https://github.com/zengfenfei/codility/blob/master/OddOccurrencesInArray.go
func OddOccurrencesInArray0(A []int) int {
	odd := 0
	for _, e := range A {
		odd ^= e
	}
	return odd
}

func OddOccurrencesInArray(A []int) int {
	counters := make(map[int]int)

	for _, e := range A {
		counters[e]++
	}

	for k, v := range counters {
		if v == 1 {
			return k
		}
	}
	return 0
}
