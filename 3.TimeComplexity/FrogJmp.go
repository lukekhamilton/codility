package TimeComplexity

func FrogJmp(X int, Y int, D int) int {
	d := Y - X
	s := d / D
	if d%D == 0 {
		return s
	} else {
		return s + 1
	}
}
