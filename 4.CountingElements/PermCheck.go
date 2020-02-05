package CountingElements

func PermCheck(A []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	occurred := make(map[int]bool)
	for _, e := range A {
		if e < 1 || e > n {
			return 0
		}
		if occurred[e] {
			return 0
		}
		occurred[e] = true
	}
	return 1
}
