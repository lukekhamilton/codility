package TimeComplexity

func PermMissingElem(A []int) int {
	n := len(A) + 1
	total := n * (n + 1) / 2
	for _, e := range A {
		total -= e
	}
	return total
}
